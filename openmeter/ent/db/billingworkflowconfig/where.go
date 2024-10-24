// Code generated by ent, DO NOT EDIT.

package billingworkflowconfig

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	billingentity "github.com/openmeterio/openmeter/openmeter/billing/entity"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
	"github.com/openmeterio/openmeter/pkg/datex"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldContainsFold(FieldID, id))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldNamespace, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldDeletedAt, v))
}

// LineCollectionPeriod applies equality check predicate on the "line_collection_period" field. It's identical to LineCollectionPeriodEQ.
func LineCollectionPeriod(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldLineCollectionPeriod, vc))
}

// InvoiceAutoAdvance applies equality check predicate on the "invoice_auto_advance" field. It's identical to InvoiceAutoAdvanceEQ.
func InvoiceAutoAdvance(v bool) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceAutoAdvance, v))
}

// InvoiceDraftPeriod applies equality check predicate on the "invoice_draft_period" field. It's identical to InvoiceDraftPeriodEQ.
func InvoiceDraftPeriod(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceDraftPeriod, vc))
}

// InvoiceDueAfter applies equality check predicate on the "invoice_due_after" field. It's identical to InvoiceDueAfterEQ.
func InvoiceDueAfter(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceDueAfter, vc))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldContainsFold(FieldNamespace, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotNull(FieldDeletedAt))
}

// CollectionAlignmentEQ applies the EQ predicate on the "collection_alignment" field.
func CollectionAlignmentEQ(v billingentity.AlignmentKind) predicate.BillingWorkflowConfig {
	vc := v
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldCollectionAlignment, vc))
}

// CollectionAlignmentNEQ applies the NEQ predicate on the "collection_alignment" field.
func CollectionAlignmentNEQ(v billingentity.AlignmentKind) predicate.BillingWorkflowConfig {
	vc := v
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldCollectionAlignment, vc))
}

// CollectionAlignmentIn applies the In predicate on the "collection_alignment" field.
func CollectionAlignmentIn(vs ...billingentity.AlignmentKind) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldCollectionAlignment, v...))
}

// CollectionAlignmentNotIn applies the NotIn predicate on the "collection_alignment" field.
func CollectionAlignmentNotIn(vs ...billingentity.AlignmentKind) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldCollectionAlignment, v...))
}

// LineCollectionPeriodEQ applies the EQ predicate on the "line_collection_period" field.
func LineCollectionPeriodEQ(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldLineCollectionPeriod, vc))
}

// LineCollectionPeriodNEQ applies the NEQ predicate on the "line_collection_period" field.
func LineCollectionPeriodNEQ(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldLineCollectionPeriod, vc))
}

// LineCollectionPeriodIn applies the In predicate on the "line_collection_period" field.
func LineCollectionPeriodIn(vs ...datex.ISOString) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldLineCollectionPeriod, v...))
}

// LineCollectionPeriodNotIn applies the NotIn predicate on the "line_collection_period" field.
func LineCollectionPeriodNotIn(vs ...datex.ISOString) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldLineCollectionPeriod, v...))
}

// LineCollectionPeriodGT applies the GT predicate on the "line_collection_period" field.
func LineCollectionPeriodGT(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldLineCollectionPeriod, vc))
}

// LineCollectionPeriodGTE applies the GTE predicate on the "line_collection_period" field.
func LineCollectionPeriodGTE(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldLineCollectionPeriod, vc))
}

// LineCollectionPeriodLT applies the LT predicate on the "line_collection_period" field.
func LineCollectionPeriodLT(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldLineCollectionPeriod, vc))
}

// LineCollectionPeriodLTE applies the LTE predicate on the "line_collection_period" field.
func LineCollectionPeriodLTE(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldLineCollectionPeriod, vc))
}

// LineCollectionPeriodContains applies the Contains predicate on the "line_collection_period" field.
func LineCollectionPeriodContains(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldContains(FieldLineCollectionPeriod, vc))
}

// LineCollectionPeriodHasPrefix applies the HasPrefix predicate on the "line_collection_period" field.
func LineCollectionPeriodHasPrefix(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldHasPrefix(FieldLineCollectionPeriod, vc))
}

// LineCollectionPeriodHasSuffix applies the HasSuffix predicate on the "line_collection_period" field.
func LineCollectionPeriodHasSuffix(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldHasSuffix(FieldLineCollectionPeriod, vc))
}

// LineCollectionPeriodEqualFold applies the EqualFold predicate on the "line_collection_period" field.
func LineCollectionPeriodEqualFold(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldEqualFold(FieldLineCollectionPeriod, vc))
}

// LineCollectionPeriodContainsFold applies the ContainsFold predicate on the "line_collection_period" field.
func LineCollectionPeriodContainsFold(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldContainsFold(FieldLineCollectionPeriod, vc))
}

// InvoiceAutoAdvanceEQ applies the EQ predicate on the "invoice_auto_advance" field.
func InvoiceAutoAdvanceEQ(v bool) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceAutoAdvance, v))
}

// InvoiceAutoAdvanceNEQ applies the NEQ predicate on the "invoice_auto_advance" field.
func InvoiceAutoAdvanceNEQ(v bool) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldInvoiceAutoAdvance, v))
}

// InvoiceDraftPeriodEQ applies the EQ predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodEQ(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceDraftPeriod, vc))
}

// InvoiceDraftPeriodNEQ applies the NEQ predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodNEQ(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldInvoiceDraftPeriod, vc))
}

// InvoiceDraftPeriodIn applies the In predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodIn(vs ...datex.ISOString) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldInvoiceDraftPeriod, v...))
}

// InvoiceDraftPeriodNotIn applies the NotIn predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodNotIn(vs ...datex.ISOString) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldInvoiceDraftPeriod, v...))
}

// InvoiceDraftPeriodGT applies the GT predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodGT(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldInvoiceDraftPeriod, vc))
}

// InvoiceDraftPeriodGTE applies the GTE predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodGTE(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldInvoiceDraftPeriod, vc))
}

// InvoiceDraftPeriodLT applies the LT predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodLT(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldInvoiceDraftPeriod, vc))
}

// InvoiceDraftPeriodLTE applies the LTE predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodLTE(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldInvoiceDraftPeriod, vc))
}

// InvoiceDraftPeriodContains applies the Contains predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodContains(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldContains(FieldInvoiceDraftPeriod, vc))
}

// InvoiceDraftPeriodHasPrefix applies the HasPrefix predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodHasPrefix(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldHasPrefix(FieldInvoiceDraftPeriod, vc))
}

// InvoiceDraftPeriodHasSuffix applies the HasSuffix predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodHasSuffix(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldHasSuffix(FieldInvoiceDraftPeriod, vc))
}

// InvoiceDraftPeriodEqualFold applies the EqualFold predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodEqualFold(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldEqualFold(FieldInvoiceDraftPeriod, vc))
}

// InvoiceDraftPeriodContainsFold applies the ContainsFold predicate on the "invoice_draft_period" field.
func InvoiceDraftPeriodContainsFold(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldContainsFold(FieldInvoiceDraftPeriod, vc))
}

// InvoiceDueAfterEQ applies the EQ predicate on the "invoice_due_after" field.
func InvoiceDueAfterEQ(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceDueAfter, vc))
}

// InvoiceDueAfterNEQ applies the NEQ predicate on the "invoice_due_after" field.
func InvoiceDueAfterNEQ(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldInvoiceDueAfter, vc))
}

// InvoiceDueAfterIn applies the In predicate on the "invoice_due_after" field.
func InvoiceDueAfterIn(vs ...datex.ISOString) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldInvoiceDueAfter, v...))
}

// InvoiceDueAfterNotIn applies the NotIn predicate on the "invoice_due_after" field.
func InvoiceDueAfterNotIn(vs ...datex.ISOString) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldInvoiceDueAfter, v...))
}

// InvoiceDueAfterGT applies the GT predicate on the "invoice_due_after" field.
func InvoiceDueAfterGT(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldInvoiceDueAfter, vc))
}

// InvoiceDueAfterGTE applies the GTE predicate on the "invoice_due_after" field.
func InvoiceDueAfterGTE(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldInvoiceDueAfter, vc))
}

// InvoiceDueAfterLT applies the LT predicate on the "invoice_due_after" field.
func InvoiceDueAfterLT(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldInvoiceDueAfter, vc))
}

// InvoiceDueAfterLTE applies the LTE predicate on the "invoice_due_after" field.
func InvoiceDueAfterLTE(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldInvoiceDueAfter, vc))
}

// InvoiceDueAfterContains applies the Contains predicate on the "invoice_due_after" field.
func InvoiceDueAfterContains(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldContains(FieldInvoiceDueAfter, vc))
}

// InvoiceDueAfterHasPrefix applies the HasPrefix predicate on the "invoice_due_after" field.
func InvoiceDueAfterHasPrefix(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldHasPrefix(FieldInvoiceDueAfter, vc))
}

// InvoiceDueAfterHasSuffix applies the HasSuffix predicate on the "invoice_due_after" field.
func InvoiceDueAfterHasSuffix(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldHasSuffix(FieldInvoiceDueAfter, vc))
}

// InvoiceDueAfterEqualFold applies the EqualFold predicate on the "invoice_due_after" field.
func InvoiceDueAfterEqualFold(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldEqualFold(FieldInvoiceDueAfter, vc))
}

// InvoiceDueAfterContainsFold applies the ContainsFold predicate on the "invoice_due_after" field.
func InvoiceDueAfterContainsFold(v datex.ISOString) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldContainsFold(FieldInvoiceDueAfter, vc))
}

// InvoiceCollectionMethodEQ applies the EQ predicate on the "invoice_collection_method" field.
func InvoiceCollectionMethodEQ(v billingentity.CollectionMethod) predicate.BillingWorkflowConfig {
	vc := v
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceCollectionMethod, vc))
}

// InvoiceCollectionMethodNEQ applies the NEQ predicate on the "invoice_collection_method" field.
func InvoiceCollectionMethodNEQ(v billingentity.CollectionMethod) predicate.BillingWorkflowConfig {
	vc := v
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldInvoiceCollectionMethod, vc))
}

// InvoiceCollectionMethodIn applies the In predicate on the "invoice_collection_method" field.
func InvoiceCollectionMethodIn(vs ...billingentity.CollectionMethod) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldInvoiceCollectionMethod, v...))
}

// InvoiceCollectionMethodNotIn applies the NotIn predicate on the "invoice_collection_method" field.
func InvoiceCollectionMethodNotIn(vs ...billingentity.CollectionMethod) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldInvoiceCollectionMethod, v...))
}

// HasBillingInvoices applies the HasEdge predicate on the "billing_invoices" edge.
func HasBillingInvoices() predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BillingInvoicesTable, BillingInvoicesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingInvoicesWith applies the HasEdge predicate on the "billing_invoices" edge with a given conditions (other predicates).
func HasBillingInvoicesWith(preds ...predicate.BillingInvoice) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(func(s *sql.Selector) {
		step := newBillingInvoicesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBillingProfile applies the HasEdge predicate on the "billing_profile" edge.
func HasBillingProfile() predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BillingProfileTable, BillingProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingProfileWith applies the HasEdge predicate on the "billing_profile" edge with a given conditions (other predicates).
func HasBillingProfileWith(preds ...predicate.BillingProfile) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(func(s *sql.Selector) {
		step := newBillingProfileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BillingWorkflowConfig) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BillingWorkflowConfig) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BillingWorkflowConfig) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.NotPredicates(p))
}
