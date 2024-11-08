package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/alpacahq/alpacadecimal"

	billingentity "github.com/openmeterio/openmeter/openmeter/billing/entity"
	"github.com/openmeterio/openmeter/pkg/currencyx"
	"github.com/openmeterio/openmeter/pkg/datex"
	"github.com/openmeterio/openmeter/pkg/framework/entutils"
	"github.com/openmeterio/openmeter/pkg/timezone"
)

type BillingProfile struct {
	ent.Schema
}

func (BillingProfile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entutils.ResourceMixin{},
		entutils.CustomerAddressMixin{
			FieldPrefix: "supplier",
		},
	}
}

func (BillingProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("tax_app_id").
			Immutable().
			SchemaType(map[string]string{
				dialect.Postgres: "char(26)",
			}),
		field.String("invoicing_app_id").
			Immutable().
			SchemaType(map[string]string{
				dialect.Postgres: "char(26)",
			}),
		field.String("payment_app_id").
			Immutable().
			SchemaType(map[string]string{
				dialect.Postgres: "char(26)",
			}),
		field.String("workflow_config_id").
			NotEmpty(),
		field.Bool("default").
			Default(false),
		field.String("supplier_name").
			NotEmpty(),
		field.String("supplier_tax_code").
			Optional().
			Nillable(),
	}
}

func (BillingProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("billing_invoices", BillingInvoice.Type),
		edge.To("billing_customer_override", BillingCustomerOverride.Type),
		edge.From("workflow_config", BillingWorkflowConfig.Type).
			Ref("billing_profile").
			Field("workflow_config_id").
			Unique().
			Required(),
		edge.From("tax_app", App.Type).
			Ref("billing_profile_tax_app").
			Field("tax_app_id").
			Unique().
			Immutable().
			Required(),
		edge.From("invoicing_app", App.Type).
			Ref("billing_profile_invoicing_app").
			Field("invoicing_app_id").
			Unique().
			Immutable().
			Required(),
		edge.From("payment_app", App.Type).
			Ref("billing_profile_payment_app").
			Field("payment_app_id").
			Unique().
			Immutable().
			Required(),
	}
}

func (BillingProfile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("namespace", "default", "deleted_at").
			Unique(),
	}
}

type BillingWorkflowConfig struct {
	ent.Schema
}

func (BillingWorkflowConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entutils.IDMixin{},
		entutils.NamespaceMixin{},
		entutils.TimeMixin{},
	}
}

func (BillingWorkflowConfig) Fields() []ent.Field {
	return []ent.Field{
		// TODO: later we will add more alignment details here (e.g. monthly, yearly, etc.)
		field.Enum("collection_alignment").
			GoType(billingentity.AlignmentKind("")),

		field.String("line_collection_period").GoType(datex.ISOString("")),

		field.Bool("invoice_auto_advance"),

		field.String("invoice_draft_period").GoType(datex.ISOString("")),

		field.String("invoice_due_after").GoType(datex.ISOString("")),

		field.Enum("invoice_collection_method").
			GoType(billingentity.CollectionMethod("")),
	}
}

func (BillingWorkflowConfig) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("namespace", "id"),
	}
}

func (BillingWorkflowConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("billing_invoices", BillingInvoice.Type).
			Unique(),
		edge.To("billing_profile", BillingProfile.Type).
			Unique(),
	}
}

type BillingCustomerOverride struct {
	ent.Schema
}

func (BillingCustomerOverride) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entutils.IDMixin{},
		entutils.NamespaceMixin{},
		entutils.TimeMixin{},
	}
}

func (BillingCustomerOverride) Fields() []ent.Field {
	return []ent.Field{
		field.String("customer_id").
			Unique().
			Immutable().
			SchemaType(map[string]string{
				dialect.Postgres: "char(26)",
			}),

		// For now we are not allowing for provider type overrides (that should be a separate billing provider entity).
		//
		// When we have the provider configs ready, we will add the field overrides for those specific fields.
		field.String("billing_profile_id").
			Optional().
			Nillable().
			SchemaType(map[string]string{
				dialect.Postgres: "char(26)",
			}),

		// Workflow config overrides
		// TODO: later we will add more alignment details here (e.g. monthly, yearly, etc.)
		field.Enum("collection_alignment").
			GoType(billingentity.AlignmentKind("")).
			Optional().
			Nillable(),

		field.String("line_collection_period").
			GoType(datex.ISOString("")).
			Optional().
			Nillable(),

		field.Bool("invoice_auto_advance").
			Optional().
			Nillable(),

		field.String("invoice_draft_period").
			GoType(datex.ISOString("")).
			Optional().
			Nillable(),

		field.String("invoice_due_after").
			GoType(datex.ISOString("")).
			Optional().
			Nillable(),

		field.Enum("invoice_collection_method").
			GoType(billingentity.CollectionMethod("")).
			Optional().
			Nillable(),
	}
}

func (BillingCustomerOverride) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("namespace", "id").Unique(),
		index.Fields("namespace", "customer_id").Unique(),
	}
}

func (BillingCustomerOverride) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("billing_customer_override").
			Field("customer_id").
			Unique().
			Required().
			Immutable(),
		edge.From("billing_profile", BillingProfile.Type).
			Ref("billing_customer_override").
			Field("billing_profile_id").
			Unique(),
	}
}

type BillingInvoiceLine struct {
	ent.Schema
}

func (BillingInvoiceLine) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entutils.ResourceMixin{},
	}
}

func (BillingInvoiceLine) Fields() []ent.Field {
	return []ent.Field{
		field.String("invoice_id").
			SchemaType(map[string]string{
				"postgres": "char(26)",
			}),

		field.Time("period_start"),
		field.Time("period_end"),
		field.Time("invoice_at"),

		// TODO[dependency]: overrides (as soon as plan override entities are ready)

		field.Enum("type").
			GoType(billingentity.InvoiceLineType("")),

		field.Enum("status").
			GoType(billingentity.InvoiceLineStatus("")),

		field.String("currency").
			GoType(currencyx.Code("")).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				"postgres": "varchar(3)",
			}),

		// Quantity is optional as for usage-based billing we can only persist this value,
		// when the invoice is issued
		field.Other("quantity", alpacadecimal.Decimal{}).
			Optional().
			Nillable().
			SchemaType(map[string]string{
				"postgres": "numeric",
			}),

		field.JSON("tax_overrides", &billingentity.TaxOverrides{}).
			SchemaType(map[string]string{
				"postgres": "jsonb",
			}).
			Optional(),
	}
}

func (BillingInvoiceLine) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("namespace", "invoice_id"),
	}
}

func (BillingInvoiceLine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("billing_invoice", BillingInvoice.Type).
			Ref("billing_invoice_lines").
			Field("invoice_id").
			Unique().
			Required(),
		edge.To("billing_invoice_manual_lines", BillingInvoiceManualLineConfig.Type).
			StorageKey(edge.Column("manual_line_config_id")).
			Unique(),
	}
}

type BillingInvoiceManualLineConfig struct {
	ent.Schema
}

func (BillingInvoiceManualLineConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entutils.NamespaceMixin{},
		entutils.IDMixin{},
	}
}

func (BillingInvoiceManualLineConfig) Fields() []ent.Field {
	return []ent.Field{
		// Either coming from the
		field.Other("unit_price", alpacadecimal.Decimal{}).
			SchemaType(map[string]string{
				"postgres": "numeric",
			}),
	}
}

func (BillingInvoiceManualLineConfig) Edgexs() []ent.Edge {
	return []ent.Edge{
		edge.From("billing_invoice_line", BillingInvoiceLine.Type).
			Ref("billing_invoice_manual_lines").
			Field("id").
			Unique().
			Required(),
	}
}

type BillingInvoice struct {
	ent.Schema
}

func (BillingInvoice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// This cannot be a resource mixin as the invoice doesn't have a name field
		entutils.IDMixin{},
		entutils.NamespaceMixin{},
		entutils.MetadataAnnotationsMixin{},
		entutils.TimeMixin{},
		entutils.CustomerAddressMixin{
			FieldPrefix: "supplier",
		},
		entutils.CustomerAddressMixin{
			FieldPrefix: "customer",
		},
	}
}

func (BillingInvoice) Fields() []ent.Field {
	return []ent.Field{
		// Customer/supplier
		field.String("supplier_name").
			NotEmpty(),

		field.String("supplier_tax_code").
			Optional().
			Nillable(),

		field.String("customer_name").
			NotEmpty(),

		field.String("customer_timezone").
			GoType(timezone.Timezone("")).
			Optional().
			Nillable(),

		// Invoice number
		field.String("number").
			Optional().
			Nillable(),

		field.Enum("type").
			GoType(billingentity.InvoiceType("")),

		field.String("description").
			Optional().
			Nillable(),

		field.String("customer_id").
			NotEmpty().
			SchemaType(map[string]string{
				"postgres": "char(26)",
			}).
			Immutable(),

		field.String("source_billing_profile_id").
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				"postgres": "char(26)",
			}),
		field.Time("voided_at").
			Optional().
			Nillable(),

		// issued_at can be in the future in case of pre-issuing invoices
		field.Time("issued_at").
			Optional().
			Nillable(),

		field.Time("draft_until").
			Optional().
			Nillable(),

		field.String("currency").
			GoType(currencyx.Code("")).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				"postgres": "varchar(3)",
			}),

		field.Time("due_at").
			Optional().
			Nillable(),

		field.Enum("status").
			GoType(billingentity.InvoiceStatus("")),

		// Cloned profile settings
		field.String("workflow_config_id").
			SchemaType(map[string]string{
				"postgres": "char(26)",
			}),

		field.String("tax_app_id").
			Immutable().
			SchemaType(map[string]string{
				dialect.Postgres: "char(26)",
			}),
		field.String("invoicing_app_id").
			Immutable().
			SchemaType(map[string]string{
				dialect.Postgres: "char(26)",
			}),
		field.String("payment_app_id").
			Immutable().
			SchemaType(map[string]string{
				dialect.Postgres: "char(26)",
			}),

		// These fields are optional as they are calculated from the invoice lines, which might not
		// be present on an invoice.
		field.Time("period_start").
			Optional().
			Nillable(),

		field.Time("period_end").
			Optional().
			Nillable(),
	}
}

func (BillingInvoice) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("namespace", "id"),
		index.Fields("namespace", "customer_id"),
	}
}

func (BillingInvoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("source_billing_profile", BillingProfile.Type).
			Ref("billing_invoices").
			Field("source_billing_profile_id").
			Required().
			Unique().
			Immutable(), // Billing profile changes are forbidden => invoice must be voided in this case
		edge.From("billing_workflow_config", BillingWorkflowConfig.Type).
			Ref("billing_invoices").
			Field("workflow_config_id").
			Unique().
			Required(),
		edge.To("billing_invoice_lines", BillingInvoiceLine.Type),
		edge.To("billing_invoice_validation_issues", BillingInvoiceValidationIssue.Type),
		edge.From("billing_invoice_customer", Customer.Type).
			Ref("billing_invoice").
			Field("customer_id").
			Unique().
			Required().
			Immutable(),
		edge.From("tax_app", App.Type).
			Ref("billing_invoice_tax_app").
			Field("tax_app_id").
			Unique().
			Immutable().
			Required(),
		edge.From("invoicing_app", App.Type).
			Ref("billing_invoice_invoicing_app").
			Field("invoicing_app_id").
			Unique().
			Immutable().
			Required(),
		edge.From("payment_app", App.Type).
			Ref("billing_invoice_payment_app").
			Field("payment_app_id").
			Unique().
			Immutable().
			Required(),
	}
}

type BillingInvoiceValidationIssue struct {
	ent.Schema
}

func (BillingInvoiceValidationIssue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entutils.IDMixin{},
		entutils.NamespaceMixin{},
		entutils.TimeMixin{},
	}
}

func (BillingInvoiceValidationIssue) Fields() []ent.Field {
	return []ent.Field{
		field.String("invoice_id").
			NotEmpty().
			SchemaType(map[string]string{
				"postgres": "char(26)",
			}),

		field.Enum("severity").
			GoType(billingentity.ValidationIssueSeverity("")),

		field.String("code").
			Nillable().
			Optional(),

		field.String("message").
			NotEmpty(),

		// Note: field is conflicting with the ent builtin methods, so we use path instead
		field.String("path").
			Optional().
			Nillable(),

		field.String("component"),

		field.Bytes("dedupe_hash").
			MinLen(32).
			MaxLen(32),
	}
}

func (BillingInvoiceValidationIssue) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("namespace", "invoice_id", "dedupe_hash").Unique(),
	}
}

func (BillingInvoiceValidationIssue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("billing_invoice", BillingInvoice.Type).
			Ref("billing_invoice_validation_issues").
			Field("invoice_id").
			Unique().
			Required(),
	}
}
