// Code generated by ent, DO NOT EDIT.

package notificationeventdeliverystatus

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/openmeterio/openmeter/internal/ent/db/predicate"
	"github.com/openmeterio/openmeter/internal/notification"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldContainsFold(FieldID, id))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldEQ(FieldNamespace, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldEQ(FieldUpdatedAt, v))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldContainsFold(FieldNamespace, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.FieldLTE(FieldUpdatedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v notification.EventType) predicate.NotificationEventDeliveryStatus {
	vc := v
	return predicate.NotificationEventDeliveryStatus(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v notification.EventType) predicate.NotificationEventDeliveryStatus {
	vc := v
	return predicate.NotificationEventDeliveryStatus(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...notification.EventType) predicate.NotificationEventDeliveryStatus {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotificationEventDeliveryStatus(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...notification.EventType) predicate.NotificationEventDeliveryStatus {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotificationEventDeliveryStatus(sql.FieldNotIn(FieldType, v...))
}

// StateEQ applies the EQ predicate on the "State" field.
func StateEQ(v notification.EventDeliveryStatusState) predicate.NotificationEventDeliveryStatus {
	vc := v
	return predicate.NotificationEventDeliveryStatus(sql.FieldEQ(FieldState, vc))
}

// StateNEQ applies the NEQ predicate on the "State" field.
func StateNEQ(v notification.EventDeliveryStatusState) predicate.NotificationEventDeliveryStatus {
	vc := v
	return predicate.NotificationEventDeliveryStatus(sql.FieldNEQ(FieldState, vc))
}

// StateIn applies the In predicate on the "State" field.
func StateIn(vs ...notification.EventDeliveryStatusState) predicate.NotificationEventDeliveryStatus {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotificationEventDeliveryStatus(sql.FieldIn(FieldState, v...))
}

// StateNotIn applies the NotIn predicate on the "State" field.
func StateNotIn(vs ...notification.EventDeliveryStatusState) predicate.NotificationEventDeliveryStatus {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotificationEventDeliveryStatus(sql.FieldNotIn(FieldState, v...))
}

// HasEvents applies the HasEdge predicate on the "events" edge.
func HasEvents() predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, EventsTable, EventsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventsWith applies the HasEdge predicate on the "events" edge with a given conditions (other predicates).
func HasEventsWith(preds ...predicate.NotificationEvent) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(func(s *sql.Selector) {
		step := newEventsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NotificationEventDeliveryStatus) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NotificationEventDeliveryStatus) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NotificationEventDeliveryStatus) predicate.NotificationEventDeliveryStatus {
	return predicate.NotificationEventDeliveryStatus(sql.NotPredicates(p))
}
