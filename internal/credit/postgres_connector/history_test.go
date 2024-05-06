package postgres_connector

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"

	"github.com/openmeterio/openmeter/internal/credit"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db"
	meter_model "github.com/openmeterio/openmeter/internal/meter"
	"github.com/openmeterio/openmeter/pkg/models"
)

func TestPostgresConnectorLedger(t *testing.T) {
	namespace := "default"

	meter := models.Meter{
		Namespace:   namespace,
		ID:          "meter-1",
		Slug:        "meter-1",
		Aggregation: models.MeterAggregationSum,
	}
	meterRepository := meter_model.NewInMemoryRepository([]models.Meter{meter})
	featureIn := credit.Feature{
		Namespace: namespace,
		MeterSlug: meter.Slug,
		Name:      "feature-1",
	}

	tt := []struct {
		name        string
		description string
		test        func(t *testing.T, connector credit.Connector, streamingConnector *mockStreamingConnector, db_client *db.Client, ledger credit.Ledger)
	}{
		{
			name:        "GetHistory",
			description: "Should return ledger entries",
			test: func(t *testing.T, connector credit.Connector, streamingConnector *mockStreamingConnector, db_client *db.Client, ledger credit.Ledger) {
				ctx := context.Background()
				feature := createFeature(t, connector, namespace, featureIn)
				// We need to truncate the time to workaround pgx driver timezone issue
				// We also move it to the past to avoid timezone issues
				t1 := time.Now().Truncate(time.Hour * 24).Add(-time.Hour * 24).In(time.UTC)
				t2 := t1.Add(time.Hour).Truncate(0).In(time.UTC)
				t3 := t2.Add(time.Hour).Truncate(0).In(time.UTC)
				t4 := t3.Add(time.Hour).Truncate(0).In(time.UTC)

				grant1, err := connector.CreateGrant(ctx, namespace, credit.Grant{
					LedgerID:    ledger.ID,
					FeatureID:   feature.ID,
					Type:        credit.GrantTypeUsage,
					Amount:      100,
					Priority:    1,
					EffectiveAt: t1,
					Expiration: credit.ExpirationPeriod{
						Duration: credit.ExpirationPeriodDurationMonth,
						Count:    1,
					},
					Rollover: &credit.GrantRollover{
						Type: credit.GrantRolloverTypeRemainingAmount,
					},
				})
				assert.NoError(t, err)

				grant2, err := connector.CreateGrant(ctx, namespace, credit.Grant{
					LedgerID:    ledger.ID,
					FeatureID:   feature.ID,
					Type:        credit.GrantTypeUsage,
					Amount:      100,
					Priority:    1,
					EffectiveAt: t2,
					Expiration: credit.ExpirationPeriod{
						Duration: credit.ExpirationPeriodDurationMonth,
						Count:    1,
					},
					Rollover: &credit.GrantRollover{
						Type: credit.GrantRolloverTypeOriginalAmount,
					},
				})
				assert.NoError(t, err)

				// Void grant2
				_, err = connector.VoidGrant(ctx, namespace, grant2)
				assert.NoError(t, err)

				usage := 1.0
				streamingConnector.addRow(meter.Slug, models.MeterQueryRow{
					Value: usage,
					// Grant 1's effective time is t1, so usage starts from t1
					WindowStart: t1,
					// Reset time is t3, so usage ends at t3
					WindowEnd: t3,
					GroupBy:   map[string]*string{},
				})

				reset, rolloverGrants, err := connector.Reset(ctx, namespace, credit.Reset{
					LedgerID:    ledger.ID,
					EffectiveAt: t3,
				})
				assert.NoError(t, err)

				// Get ledger
				ledgerList, err := connector.GetHistory(ctx, namespace, ledger.ID, t1, t4, 0)
				assert.NoError(t, err)

				// Expected
				ledgerEntries := ledgerList.GetEntries()
				ledgerUsage := -1 * usage
				reamingAmount := grant1.Amount - usage

				// Assert balance
				assert.Equal(t, []credit.LedgerEntry{
					// Original grant
					{
						ID:        grant1.ID,
						Type:      credit.LedgerEntryTypeGrant,
						Time:      t1,
						FeatureID: feature.ID,
						Amount:    &grant1.Amount,
					},
					// Void
					{
						ID:        grant2.ID,
						Type:      credit.LedgerEntryTypeVoid,
						Time:      t2,
						FeatureID: feature.ID,
						Amount:    &grant2.Amount,
					},
					// Usage
					{
						ID:        grant1.ID,
						Type:      credit.LedgerEntryTypeGrantUsage,
						Time:      t3,
						FeatureID: feature.ID,
						Amount:    &ledgerUsage,
						Period: &credit.Period{
							From: t1,
							To:   t3,
						},
					},
					// Reset
					{
						ID:   reset.ID,
						Type: credit.LedgerEntryTypeReset,
						Time: t3,
					},
					// Rolled over grant
					{
						ID:        rolloverGrants[0].ID,
						Type:      credit.LedgerEntryTypeGrant,
						Time:      t3,
						FeatureID: feature.ID,
						Amount:    &reamingAmount,
					},
				}, ledgerEntries)
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tc.description)
			driver := initDB(t)
			databaseClient := db.NewClient(db.Driver(driver))
			defer databaseClient.Close()

			// Note: lock manager cannot be shared between tests as these parallel tests write the same ledger
			streamingConnector := newMockStreamingConnector()
			connector := NewPostgresConnector(slog.Default(), databaseClient, streamingConnector, meterRepository)

			// let's provision a ledger
			ledger, err := connector.CreateLedger(context.Background(), namespace, credit.Ledger{
				Subject: ulid.Make().String(),
			}, false)

			assert.NoError(t, err)

			tc.test(t, connector, streamingConnector, databaseClient, ledger)
		})
	}
}
