// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/openmeterio/openmeter/internal/ent/db/notificationevent"
	"github.com/openmeterio/openmeter/internal/ent/db/notificationrule"
	"github.com/openmeterio/openmeter/internal/notification"
)

// NotificationEvent is the model entity for the NotificationEvent schema.
type NotificationEvent struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Namespace holds the value of the "namespace" field.
	Namespace string `json:"namespace,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The event type the rule associated with
	Type notification.EventType `json:"type,omitempty"`
	// RuleID holds the value of the "rule_id" field.
	RuleID string `json:"rule_id,omitempty"`
	// Payload holds the value of the "payload" field.
	Payload string `json:"payload,omitempty"`
	// Annotations holds the value of the "annotations" field.
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NotificationEventQuery when eager-loading is set.
	Edges        NotificationEventEdges `json:"edges"`
	selectValues sql.SelectValues
}

// NotificationEventEdges holds the relations/edges for other nodes in the graph.
type NotificationEventEdges struct {
	// DeliveryStatuses holds the value of the delivery_statuses edge.
	DeliveryStatuses []*NotificationEventDeliveryStatus `json:"delivery_statuses,omitempty"`
	// Rules holds the value of the rules edge.
	Rules *NotificationRule `json:"rules,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DeliveryStatusesOrErr returns the DeliveryStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e NotificationEventEdges) DeliveryStatusesOrErr() ([]*NotificationEventDeliveryStatus, error) {
	if e.loadedTypes[0] {
		return e.DeliveryStatuses, nil
	}
	return nil, &NotLoadedError{edge: "delivery_statuses"}
}

// RulesOrErr returns the Rules value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NotificationEventEdges) RulesOrErr() (*NotificationRule, error) {
	if e.Rules != nil {
		return e.Rules, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: notificationrule.Label}
	}
	return nil, &NotLoadedError{edge: "rules"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NotificationEvent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notificationevent.FieldID, notificationevent.FieldNamespace, notificationevent.FieldType, notificationevent.FieldRuleID, notificationevent.FieldPayload:
			values[i] = new(sql.NullString)
		case notificationevent.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case notificationevent.FieldAnnotations:
			values[i] = notificationevent.ValueScanner.Annotations.ScanValue()
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NotificationEvent fields.
func (ne *NotificationEvent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notificationevent.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ne.ID = value.String
			}
		case notificationevent.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				ne.Namespace = value.String
			}
		case notificationevent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ne.CreatedAt = value.Time
			}
		case notificationevent.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ne.Type = notification.EventType(value.String)
			}
		case notificationevent.FieldRuleID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rule_id", values[i])
			} else if value.Valid {
				ne.RuleID = value.String
			}
		case notificationevent.FieldPayload:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payload", values[i])
			} else if value.Valid {
				ne.Payload = value.String
			}
		case notificationevent.FieldAnnotations:
			if value, err := notificationevent.ValueScanner.Annotations.FromValue(values[i]); err != nil {
				return err
			} else {
				ne.Annotations = value
			}
		default:
			ne.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the NotificationEvent.
// This includes values selected through modifiers, order, etc.
func (ne *NotificationEvent) Value(name string) (ent.Value, error) {
	return ne.selectValues.Get(name)
}

// QueryDeliveryStatuses queries the "delivery_statuses" edge of the NotificationEvent entity.
func (ne *NotificationEvent) QueryDeliveryStatuses() *NotificationEventDeliveryStatusQuery {
	return NewNotificationEventClient(ne.config).QueryDeliveryStatuses(ne)
}

// QueryRules queries the "rules" edge of the NotificationEvent entity.
func (ne *NotificationEvent) QueryRules() *NotificationRuleQuery {
	return NewNotificationEventClient(ne.config).QueryRules(ne)
}

// Update returns a builder for updating this NotificationEvent.
// Note that you need to call NotificationEvent.Unwrap() before calling this method if this NotificationEvent
// was returned from a transaction, and the transaction was committed or rolled back.
func (ne *NotificationEvent) Update() *NotificationEventUpdateOne {
	return NewNotificationEventClient(ne.config).UpdateOne(ne)
}

// Unwrap unwraps the NotificationEvent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ne *NotificationEvent) Unwrap() *NotificationEvent {
	_tx, ok := ne.config.driver.(*txDriver)
	if !ok {
		panic("db: NotificationEvent is not a transactional entity")
	}
	ne.config.driver = _tx.drv
	return ne
}

// String implements the fmt.Stringer.
func (ne *NotificationEvent) String() string {
	var builder strings.Builder
	builder.WriteString("NotificationEvent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ne.ID))
	builder.WriteString("namespace=")
	builder.WriteString(ne.Namespace)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ne.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", ne.Type))
	builder.WriteString(", ")
	builder.WriteString("rule_id=")
	builder.WriteString(ne.RuleID)
	builder.WriteString(", ")
	builder.WriteString("payload=")
	builder.WriteString(ne.Payload)
	builder.WriteString(", ")
	builder.WriteString("annotations=")
	builder.WriteString(fmt.Sprintf("%v", ne.Annotations))
	builder.WriteByte(')')
	return builder.String()
}

// NotificationEvents is a parsable slice of NotificationEvent.
type NotificationEvents []*NotificationEvent
