// Code generated by ent, DO NOT EDIT.

package billinginvoiceline

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/alpacahq/alpacadecimal"
	billingentity "github.com/openmeterio/openmeter/openmeter/billing/entity"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
	"github.com/openmeterio/openmeter/pkg/currencyx"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldContainsFold(FieldID, id))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldNamespace, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldDescription, v))
}

// InvoiceID applies equality check predicate on the "invoice_id" field. It's identical to InvoiceIDEQ.
func InvoiceID(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldInvoiceID, v))
}

// PeriodStart applies equality check predicate on the "period_start" field. It's identical to PeriodStartEQ.
func PeriodStart(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldPeriodStart, v))
}

// PeriodEnd applies equality check predicate on the "period_end" field. It's identical to PeriodEndEQ.
func PeriodEnd(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldPeriodEnd, v))
}

// InvoiceAt applies equality check predicate on the "invoice_at" field. It's identical to InvoiceAtEQ.
func InvoiceAt(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldInvoiceAt, v))
}

// Currency applies equality check predicate on the "currency" field. It's identical to CurrencyEQ.
func Currency(v currencyx.Code) predicate.BillingInvoiceLine {
	vc := string(v)
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldCurrency, vc))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v alpacadecimal.Decimal) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldQuantity, v))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldContainsFold(FieldNamespace, v))
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIsNull(FieldMetadata))
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotNull(FieldMetadata))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldContainsFold(FieldDescription, v))
}

// InvoiceIDEQ applies the EQ predicate on the "invoice_id" field.
func InvoiceIDEQ(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldInvoiceID, v))
}

// InvoiceIDNEQ applies the NEQ predicate on the "invoice_id" field.
func InvoiceIDNEQ(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldInvoiceID, v))
}

// InvoiceIDIn applies the In predicate on the "invoice_id" field.
func InvoiceIDIn(vs ...string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldInvoiceID, vs...))
}

// InvoiceIDNotIn applies the NotIn predicate on the "invoice_id" field.
func InvoiceIDNotIn(vs ...string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldInvoiceID, vs...))
}

// InvoiceIDGT applies the GT predicate on the "invoice_id" field.
func InvoiceIDGT(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldInvoiceID, v))
}

// InvoiceIDGTE applies the GTE predicate on the "invoice_id" field.
func InvoiceIDGTE(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldInvoiceID, v))
}

// InvoiceIDLT applies the LT predicate on the "invoice_id" field.
func InvoiceIDLT(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldInvoiceID, v))
}

// InvoiceIDLTE applies the LTE predicate on the "invoice_id" field.
func InvoiceIDLTE(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldInvoiceID, v))
}

// InvoiceIDContains applies the Contains predicate on the "invoice_id" field.
func InvoiceIDContains(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldContains(FieldInvoiceID, v))
}

// InvoiceIDHasPrefix applies the HasPrefix predicate on the "invoice_id" field.
func InvoiceIDHasPrefix(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldHasPrefix(FieldInvoiceID, v))
}

// InvoiceIDHasSuffix applies the HasSuffix predicate on the "invoice_id" field.
func InvoiceIDHasSuffix(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldHasSuffix(FieldInvoiceID, v))
}

// InvoiceIDEqualFold applies the EqualFold predicate on the "invoice_id" field.
func InvoiceIDEqualFold(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEqualFold(FieldInvoiceID, v))
}

// InvoiceIDContainsFold applies the ContainsFold predicate on the "invoice_id" field.
func InvoiceIDContainsFold(v string) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldContainsFold(FieldInvoiceID, v))
}

// PeriodStartEQ applies the EQ predicate on the "period_start" field.
func PeriodStartEQ(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldPeriodStart, v))
}

// PeriodStartNEQ applies the NEQ predicate on the "period_start" field.
func PeriodStartNEQ(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldPeriodStart, v))
}

// PeriodStartIn applies the In predicate on the "period_start" field.
func PeriodStartIn(vs ...time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldPeriodStart, vs...))
}

// PeriodStartNotIn applies the NotIn predicate on the "period_start" field.
func PeriodStartNotIn(vs ...time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldPeriodStart, vs...))
}

// PeriodStartGT applies the GT predicate on the "period_start" field.
func PeriodStartGT(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldPeriodStart, v))
}

// PeriodStartGTE applies the GTE predicate on the "period_start" field.
func PeriodStartGTE(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldPeriodStart, v))
}

// PeriodStartLT applies the LT predicate on the "period_start" field.
func PeriodStartLT(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldPeriodStart, v))
}

// PeriodStartLTE applies the LTE predicate on the "period_start" field.
func PeriodStartLTE(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldPeriodStart, v))
}

// PeriodEndEQ applies the EQ predicate on the "period_end" field.
func PeriodEndEQ(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldPeriodEnd, v))
}

// PeriodEndNEQ applies the NEQ predicate on the "period_end" field.
func PeriodEndNEQ(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldPeriodEnd, v))
}

// PeriodEndIn applies the In predicate on the "period_end" field.
func PeriodEndIn(vs ...time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldPeriodEnd, vs...))
}

// PeriodEndNotIn applies the NotIn predicate on the "period_end" field.
func PeriodEndNotIn(vs ...time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldPeriodEnd, vs...))
}

// PeriodEndGT applies the GT predicate on the "period_end" field.
func PeriodEndGT(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldPeriodEnd, v))
}

// PeriodEndGTE applies the GTE predicate on the "period_end" field.
func PeriodEndGTE(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldPeriodEnd, v))
}

// PeriodEndLT applies the LT predicate on the "period_end" field.
func PeriodEndLT(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldPeriodEnd, v))
}

// PeriodEndLTE applies the LTE predicate on the "period_end" field.
func PeriodEndLTE(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldPeriodEnd, v))
}

// InvoiceAtEQ applies the EQ predicate on the "invoice_at" field.
func InvoiceAtEQ(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldInvoiceAt, v))
}

// InvoiceAtNEQ applies the NEQ predicate on the "invoice_at" field.
func InvoiceAtNEQ(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldInvoiceAt, v))
}

// InvoiceAtIn applies the In predicate on the "invoice_at" field.
func InvoiceAtIn(vs ...time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldInvoiceAt, vs...))
}

// InvoiceAtNotIn applies the NotIn predicate on the "invoice_at" field.
func InvoiceAtNotIn(vs ...time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldInvoiceAt, vs...))
}

// InvoiceAtGT applies the GT predicate on the "invoice_at" field.
func InvoiceAtGT(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldInvoiceAt, v))
}

// InvoiceAtGTE applies the GTE predicate on the "invoice_at" field.
func InvoiceAtGTE(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldInvoiceAt, v))
}

// InvoiceAtLT applies the LT predicate on the "invoice_at" field.
func InvoiceAtLT(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldInvoiceAt, v))
}

// InvoiceAtLTE applies the LTE predicate on the "invoice_at" field.
func InvoiceAtLTE(v time.Time) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldInvoiceAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v billingentity.InvoiceLineType) predicate.BillingInvoiceLine {
	vc := v
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v billingentity.InvoiceLineType) predicate.BillingInvoiceLine {
	vc := v
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...billingentity.InvoiceLineType) predicate.BillingInvoiceLine {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...billingentity.InvoiceLineType) predicate.BillingInvoiceLine {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldType, v...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v billingentity.InvoiceLineStatus) predicate.BillingInvoiceLine {
	vc := v
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v billingentity.InvoiceLineStatus) predicate.BillingInvoiceLine {
	vc := v
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...billingentity.InvoiceLineStatus) predicate.BillingInvoiceLine {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...billingentity.InvoiceLineStatus) predicate.BillingInvoiceLine {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldStatus, v...))
}

// CurrencyEQ applies the EQ predicate on the "currency" field.
func CurrencyEQ(v currencyx.Code) predicate.BillingInvoiceLine {
	vc := string(v)
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldCurrency, vc))
}

// CurrencyNEQ applies the NEQ predicate on the "currency" field.
func CurrencyNEQ(v currencyx.Code) predicate.BillingInvoiceLine {
	vc := string(v)
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldCurrency, vc))
}

// CurrencyIn applies the In predicate on the "currency" field.
func CurrencyIn(vs ...currencyx.Code) predicate.BillingInvoiceLine {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldCurrency, v...))
}

// CurrencyNotIn applies the NotIn predicate on the "currency" field.
func CurrencyNotIn(vs ...currencyx.Code) predicate.BillingInvoiceLine {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldCurrency, v...))
}

// CurrencyGT applies the GT predicate on the "currency" field.
func CurrencyGT(v currencyx.Code) predicate.BillingInvoiceLine {
	vc := string(v)
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldCurrency, vc))
}

// CurrencyGTE applies the GTE predicate on the "currency" field.
func CurrencyGTE(v currencyx.Code) predicate.BillingInvoiceLine {
	vc := string(v)
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldCurrency, vc))
}

// CurrencyLT applies the LT predicate on the "currency" field.
func CurrencyLT(v currencyx.Code) predicate.BillingInvoiceLine {
	vc := string(v)
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldCurrency, vc))
}

// CurrencyLTE applies the LTE predicate on the "currency" field.
func CurrencyLTE(v currencyx.Code) predicate.BillingInvoiceLine {
	vc := string(v)
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldCurrency, vc))
}

// CurrencyContains applies the Contains predicate on the "currency" field.
func CurrencyContains(v currencyx.Code) predicate.BillingInvoiceLine {
	vc := string(v)
	return predicate.BillingInvoiceLine(sql.FieldContains(FieldCurrency, vc))
}

// CurrencyHasPrefix applies the HasPrefix predicate on the "currency" field.
func CurrencyHasPrefix(v currencyx.Code) predicate.BillingInvoiceLine {
	vc := string(v)
	return predicate.BillingInvoiceLine(sql.FieldHasPrefix(FieldCurrency, vc))
}

// CurrencyHasSuffix applies the HasSuffix predicate on the "currency" field.
func CurrencyHasSuffix(v currencyx.Code) predicate.BillingInvoiceLine {
	vc := string(v)
	return predicate.BillingInvoiceLine(sql.FieldHasSuffix(FieldCurrency, vc))
}

// CurrencyEqualFold applies the EqualFold predicate on the "currency" field.
func CurrencyEqualFold(v currencyx.Code) predicate.BillingInvoiceLine {
	vc := string(v)
	return predicate.BillingInvoiceLine(sql.FieldEqualFold(FieldCurrency, vc))
}

// CurrencyContainsFold applies the ContainsFold predicate on the "currency" field.
func CurrencyContainsFold(v currencyx.Code) predicate.BillingInvoiceLine {
	vc := string(v)
	return predicate.BillingInvoiceLine(sql.FieldContainsFold(FieldCurrency, vc))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v alpacadecimal.Decimal) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v alpacadecimal.Decimal) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...alpacadecimal.Decimal) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...alpacadecimal.Decimal) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v alpacadecimal.Decimal) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v alpacadecimal.Decimal) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v alpacadecimal.Decimal) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v alpacadecimal.Decimal) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldLTE(FieldQuantity, v))
}

// QuantityIsNil applies the IsNil predicate on the "quantity" field.
func QuantityIsNil() predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIsNull(FieldQuantity))
}

// QuantityNotNil applies the NotNil predicate on the "quantity" field.
func QuantityNotNil() predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotNull(FieldQuantity))
}

// TaxOverridesIsNil applies the IsNil predicate on the "tax_overrides" field.
func TaxOverridesIsNil() predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldIsNull(FieldTaxOverrides))
}

// TaxOverridesNotNil applies the NotNil predicate on the "tax_overrides" field.
func TaxOverridesNotNil() predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.FieldNotNull(FieldTaxOverrides))
}

// HasBillingInvoice applies the HasEdge predicate on the "billing_invoice" edge.
func HasBillingInvoice() predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BillingInvoiceTable, BillingInvoiceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingInvoiceWith applies the HasEdge predicate on the "billing_invoice" edge with a given conditions (other predicates).
func HasBillingInvoiceWith(preds ...predicate.BillingInvoice) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(func(s *sql.Selector) {
		step := newBillingInvoiceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBillingInvoiceManualLines applies the HasEdge predicate on the "billing_invoice_manual_lines" edge.
func HasBillingInvoiceManualLines() predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BillingInvoiceManualLinesTable, BillingInvoiceManualLinesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingInvoiceManualLinesWith applies the HasEdge predicate on the "billing_invoice_manual_lines" edge with a given conditions (other predicates).
func HasBillingInvoiceManualLinesWith(preds ...predicate.BillingInvoiceManualLineConfig) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(func(s *sql.Selector) {
		step := newBillingInvoiceManualLinesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BillingInvoiceLine) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BillingInvoiceLine) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BillingInvoiceLine) predicate.BillingInvoiceLine {
	return predicate.BillingInvoiceLine(sql.NotPredicates(p))
}
