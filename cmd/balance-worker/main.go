package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"syscall"
	"time"

	health "github.com/AppsFlyer/go-sundheit"
	healthhttp "github.com/AppsFlyer/go-sundheit/http"
	"github.com/IBM/sarama"
	"github.com/ThreeDotsLabs/watermill"
	wmkafka "github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-slog/otelslog"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	slogmulti "github.com/samber/slog-multi"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"

	"github.com/openmeterio/openmeter/config"
	"github.com/openmeterio/openmeter/internal/entitlement/balanceworker"
	"github.com/openmeterio/openmeter/internal/event/publisher"
	"github.com/openmeterio/openmeter/internal/ingest/kafkaingest"
	omwatermillkafka "github.com/openmeterio/openmeter/internal/watermill/driver/kafka"
	"github.com/openmeterio/openmeter/pkg/contextx"
	"github.com/openmeterio/openmeter/pkg/framework/operation"
	"github.com/openmeterio/openmeter/pkg/gosundheit"
	pkgkafka "github.com/openmeterio/openmeter/pkg/kafka"
	kafkametrics "github.com/openmeterio/openmeter/pkg/kafka/metrics"
)

const (
	defaultShutdownTimeout = 5 * time.Second
	otelName               = "openmeter.io/backend"
)

func main() {
	v, flags := viper.NewWithOptions(viper.WithDecodeHook(config.DecodeHook())), pflag.NewFlagSet("OpenMeter", pflag.ExitOnError)
	ctx := context.Background()

	config.SetViperDefaults(v, flags)

	flags.String("config", "", "Configuration file")
	flags.Bool("version", false, "Show version information")

	_ = flags.Parse(os.Args[1:])

	if v, _ := flags.GetBool("version"); v {
		fmt.Printf("%s version %s (%s) built on %s\n", "Open Meter", version, revision, revisionDate)

		os.Exit(0)
	}

	if c, _ := flags.GetString("config"); c != "" {
		v.SetConfigFile(c)
	}

	err := v.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		panic(err)
	}

	var conf config.Configuration
	err = v.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	err = conf.Validate()
	if err != nil {
		panic(err)
	}

	extraResources, _ := resource.New(
		context.Background(),
		resource.WithContainer(),
		resource.WithAttributes(
			semconv.ServiceName("openmeter"),
			semconv.ServiceVersion(version),
			semconv.DeploymentEnvironment(conf.Environment),
		),
	)
	res, _ := resource.Merge(
		resource.Default(),
		extraResources,
	)

	logger := slog.New(slogmulti.Pipe(
		otelslog.NewHandler,
		contextx.NewLogHandler,
		operation.NewLogHandler,
	).Handler(conf.Telemetry.Log.NewHandler(os.Stdout)))
	logger = otelslog.WithResource(logger, res)

	slog.SetDefault(logger)

	telemetryRouter := chi.NewRouter()
	telemetryRouter.Mount("/debug", middleware.Profiler())

	// Initialize OTel Metrics
	otelMeterProvider, err := conf.Telemetry.Metrics.NewMeterProvider(ctx, res)
	if err != nil {
		logger.Error("failed to initialize OpenTelemetry Metrics provider", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer func() {
		// Use dedicated context with timeout for shutdown as parent context might be canceled
		// by the time the execution reaches this stage.
		ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeout)
		defer cancel()

		if err := otelMeterProvider.Shutdown(ctx); err != nil {
			logger.Error("shutting down meter provider", slog.String("error", err.Error()))
		}
	}()
	otel.SetMeterProvider(otelMeterProvider)
	metricMeter := otelMeterProvider.Meter(otelName)

	if conf.Telemetry.Metrics.Exporters.Prometheus.Enabled {
		telemetryRouter.Handle("/metrics", promhttp.Handler())
	}

	// Initialize OTel Tracer
	otelTracerProvider, err := conf.Telemetry.Trace.NewTracerProvider(ctx, res)
	if err != nil {
		logger.Error("failed to initialize OpenTelemetry Trace provider", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer func() {
		// Use dedicated context with timeout for shutdown as parent context might be canceled
		// by the time the execution reaches this stage.
		ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeout)
		defer cancel()

		if err := otelTracerProvider.Shutdown(ctx); err != nil {
			logger.Error("shutting down tracer provider", slog.String("error", err.Error()))
		}
	}()

	otel.SetTracerProvider(otelTracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Validate service prerequisites

	if !conf.Events.Enabled {
		logger.Error("events are disabled, exiting")
		os.Exit(1)
	}

	// Configure health checker
	healthChecker := health.New(health.WithCheckListeners(gosundheit.NewLogger(logger.With(slog.String("component", "healthcheck")))))
	{
		handler := healthhttp.HandleHealthJSON(healthChecker)
		telemetryRouter.Handle("/healthz", handler)

		// Kubernetes style health checks
		telemetryRouter.HandleFunc("/healthz/live", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte("ok"))
		})
		telemetryRouter.Handle("/healthz/ready", handler)
	}

	var group run.Group

	// Create  subscriber
	wmSubscriber, err := initKafkaSubscriber(conf, logger)
	if err != nil {
		logger.Error("failed to initialize Kafka subscriber", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// Create publisher
	kafkaPublisher, err := initKafkaProducer(ctx, conf, logger, metricMeter, &group)
	if err != nil {
		logger.Error("failed to initialize Kafka producer", slog.String("error", err.Error()))
		os.Exit(1)
	}

	wmPublisher, err := initEventPublisher(ctx, logger, conf, kafkaPublisher)
	if err != nil {
		logger.Error("failed to initialize event publisher", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// Initialize worker
	workerOptions := balanceworker.WorkerOptions{
		SystemEventsTopic: conf.Events.SystemEvents.Topic,
		// TODO: IngestEventsTopic
		Subscriber: wmSubscriber,

		TargetTopic: conf.Events.SystemEvents.Topic,
		Publisher:   wmPublisher.publisher,

		Marshaler: wmPublisher.marshaler,

		Logger: logger,
	}

	if conf.BalanceWorker.PoisionQueue.Enabled {
		workerOptions.PoisonQueue = &balanceworker.WorkerPoisonQueueOptions{
			Topic:            conf.BalanceWorker.PoisionQueue.Topic,
			Throttle:         conf.BalanceWorker.PoisionQueue.Throttle.Enabled,
			ThrottleDuration: conf.BalanceWorker.PoisionQueue.Throttle.Duration,
			ThrottleCount:    conf.BalanceWorker.PoisionQueue.Throttle.Count,
		}
	}

	worker, err := balanceworker.New(workerOptions)
	if err != nil {
		logger.Error("failed to initialize worker", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// Run worker components

	// Telemetry server
	server := &http.Server{
		Addr:    conf.Telemetry.Address,
		Handler: telemetryRouter,
	}
	defer server.Close()

	group.Add(
		func() error { return server.ListenAndServe() },
		func(err error) { _ = server.Shutdown(ctx) },
	)

	// Balance worker
	group.Add(
		func() error { return worker.Run(ctx) },
		func(err error) { _ = worker.Close() },
	)

	// Handle shutdown
	group.Add(run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM))

	// Run the group
	err = group.Run()
	if e := (run.SignalError{}); errors.As(err, &e) {
		slog.Info("received signal; shutting down", slog.String("signal", e.Signal.String()))
	} else if !errors.Is(err, http.ErrServerClosed) {
		logger.Error("application stopped due to error", slog.String("error", err.Error()))
	}
}

func initKafkaSubscriber(conf config.Configuration, logger *slog.Logger) (message.Subscriber, error) {
	wmConfig := wmkafka.SubscriberConfig{
		Brokers:               []string{conf.Ingest.Kafka.Broker},
		OverwriteSaramaConfig: sarama.NewConfig(),
		ConsumerGroup:         conf.BalanceWorker.ConsumerGroupName,
		ReconnectRetrySleep:   100 * time.Millisecond,
		Unmarshaler:           wmkafka.DefaultMarshaler{},
	}

	wmConfig.OverwriteSaramaConfig.Metadata.RefreshFrequency = conf.Ingest.Kafka.TopicMetadataRefreshInterval.Duration()
	wmConfig.OverwriteSaramaConfig.ClientID = "openmeter/balance-worker"

	switch conf.Ingest.Kafka.SecurityProtocol {
	case "SASL_SSL":
		wmConfig.OverwriteSaramaConfig.Net.SASL.Enable = true
		wmConfig.OverwriteSaramaConfig.Net.SASL.User = conf.Ingest.Kafka.SaslUsername
		wmConfig.OverwriteSaramaConfig.Net.SASL.Password = conf.Ingest.Kafka.SaslPassword
		wmConfig.OverwriteSaramaConfig.Net.SASL.Mechanism = sarama.SASLMechanism(conf.Ingest.Kafka.SecurityProtocol)
		wmConfig.OverwriteSaramaConfig.Net.TLS.Enable = true
		wmConfig.OverwriteSaramaConfig.Net.TLS.Config = &tls.Config{}
	default:
	}

	if err := wmConfig.Validate(); err != nil {
		logger.Error("failed to validate Kafka subscriber configuration", slog.String("error", err.Error()))
		return nil, err
	}

	// Initialize Kafka subscriber
	subscriber, err := wmkafka.NewSubscriber(wmConfig, watermill.NewSlogLogger(logger))
	if err != nil {
		logger.Error("failed to initialize Kafka subscriber", slog.String("error", err.Error()))
		return nil, err
	}

	return subscriber, nil
}

type eventPublisher struct {
	publisher message.Publisher
	marshaler publisher.CloudEventMarshaler
}

func initEventPublisher(ctx context.Context, logger *slog.Logger, conf config.Configuration, kafkaProducer *kafka.Producer) (*eventPublisher, error) {
	eventDriver := omwatermillkafka.NewPublisher(kafkaProducer)

	if conf.BalanceWorker.PoisionQueue.AutoProvision.Enabled {
		adminClient, err := kafka.NewAdminClientFromProducer(kafkaProducer)
		if err != nil {
			return nil, fmt.Errorf("failed to create Kafka admin client: %w", err)
		}

		defer adminClient.Close()

		if err := pkgkafka.ProvisionTopic(ctx,
			adminClient,
			logger,
			conf.BalanceWorker.PoisionQueue.Topic,
			conf.BalanceWorker.PoisionQueue.AutoProvision.Partitions); err != nil {
			return nil, fmt.Errorf("failed to auto provision topic: %w", err)
		}
	}

	return &eventPublisher{
		publisher: eventDriver,
		marshaler: publisher.NewCloudEventMarshaler(omwatermillkafka.AddPartitionKeyFromSubject),
	}, nil
}

func initKafkaProducer(ctx context.Context, config config.Configuration, logger *slog.Logger, metricMeter metric.Meter, group *run.Group) (*kafka.Producer, error) {
	// Initialize Kafka Admin Client
	kafkaConfig := config.Ingest.Kafka.CreateKafkaConfig()

	// Initialize Kafka Producer
	producer, err := kafka.NewProducer(&kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("init kafka ingest: %w", err)
	}

	// Initialize Kafka Client Statistics reporter
	kafkaMetrics, err := kafkametrics.New(metricMeter)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka client metrics: %w", err)
	}

	group.Add(kafkaingest.KafkaProducerGroup(ctx, producer, logger, kafkaMetrics))

	go pkgkafka.ConsumeLogChannel(producer, logger.WithGroup("kafka").WithGroup("producer"))

	slog.Debug("connected to Kafka")
	return producer, nil
}