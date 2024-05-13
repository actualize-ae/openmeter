package credit

import (
	"context"
	"time"

	"github.com/oklog/ulid/v2"
)

type ListGrantsParams struct {
	LedgerIDs         []ulid.ULID
	From              *time.Time
	To                *time.Time
	FromHighWatermark bool
	IncludeVoid       bool
	Limit             int
}

type ListFeaturesParams struct {
	IncludeArchived bool
}

type ListLedgersParams struct {
	Subjects []string
	Offset   int
	Limit    int
}

type Connector interface {
	// Ledger
	CreateLedger(ctx context.Context, namespace string, ledger Ledger) (Ledger, error)
	ListLedgers(ctx context.Context, namespace string, params ListLedgersParams) ([]Ledger, error)
	UpsertLedger(ctx context.Context, namespace string, ledgerIn UpsertLedger) (Ledger, error)

	// Grant
	CreateGrant(ctx context.Context, namespace string, grant Grant) (Grant, error)
	VoidGrant(ctx context.Context, namespace string, grant Grant) (Grant, error)
	ListGrants(ctx context.Context, namespace string, params ListGrantsParams) ([]Grant, error)
	GetGrant(ctx context.Context, namespace string, id ulid.ULID) (Grant, error)

	// Credit
	GetBalance(ctx context.Context, namespace string, ledgerID ulid.ULID, cutline time.Time) (Balance, error)
	GetHistory(ctx context.Context, namespace string, ledgerID ulid.ULID, from time.Time, to time.Time, limit int) (LedgerEntryList, error)
	GetHighWatermark(ctx context.Context, namespace string, ledgerID ulid.ULID) (HighWatermark, error)
	Reset(ctx context.Context, namespace string, reset Reset) (Reset, []Grant, error)

	// Feature
	CreateFeature(ctx context.Context, namespace string, feature Feature) (Feature, error)
	DeleteFeature(ctx context.Context, namespace string, ledgerID ulid.ULID) error
	ListFeatures(ctx context.Context, namespace string, params ListFeaturesParams) ([]Feature, error)
	GetFeature(ctx context.Context, namespace string, ledgerID ulid.ULID) (Feature, error)
}
