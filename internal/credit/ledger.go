package credit

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/shopspring/decimal"
)

type Ledger struct {
	// ID is the ID of the ledger instance
	ID ulid.ULID `json:"id,omitempty"`

	// Subject specifies which metering subject this ledger is referring to
	Subject string `json:"subject"`

	Metadata map[string]string `json:"metadata,omitempty"`
}

// Render implements the chi renderer interface.
func (c Ledger) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type LedgerEntryType string

// Used to sort ledger entries by type.
var entryTypeWeight = map[LedgerEntryType]int{
	LedgerEntryTypeGrantUsage: 1,
	LedgerEntryTypeReset:      2,
	LedgerEntryTypeGrant:      3,
	LedgerEntryTypeVoid:       4,
}

// Defines values for LedgerEntryType.
const (
	LedgerEntryTypeGrant      LedgerEntryType = "GRANT"
	LedgerEntryTypeVoid       LedgerEntryType = "VOID"
	LedgerEntryTypeReset      LedgerEntryType = "RESET"
	LedgerEntryTypeGrantUsage LedgerEntryType = "GRANT_USAGE"
)

func (LedgerEntryType) Values() (kinds []string) {
	for _, s := range []LedgerEntryType{
		LedgerEntryTypeGrant,
		LedgerEntryTypeVoid,
		LedgerEntryTypeReset,
		LedgerEntryTypeGrantUsage,
	} {
		kinds = append(kinds, string(s))
	}
	return
}

// LedgerEntry is a credit ledger entry.
type LedgerEntry struct {
	ID        *ulid.ULID       `json:"id,omitempty"`
	Type      LedgerEntryType  `json:"type"`
	Time      time.Time        `json:"time"`
	FeatureID *ulid.ULID       `json:"featureId,omitempty"`
	Amount    *decimal.Decimal `json:"amount,omitempty"`
	Period    *Period          `json:"period,omitempty"`
}

type Period struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

// Render implements the chi renderer interface.
func (c LedgerEntry) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewLedgerEntryList() LedgerEntryList {
	return LedgerEntryList{
		list: []LedgerEntry{},
	}
}

type LedgerEntryList struct {
	list []LedgerEntry
}

func (c LedgerEntryList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (f LedgerEntryList) GetEntries() []LedgerEntry {
	list := make([]LedgerEntry, len(f.list))
	_ = copy(list, f.list)

	// Sort ledger entries by time
	sort.Slice(list, func(i, j int) bool {
		// Sort ledger entries by type if time is equal
		if list[i].Time.Equal(list[j].Time) {
			return entryTypeWeight[list[i].Type] < entryTypeWeight[list[j].Type]
		}

		return list[i].Time.Before(list[j].Time)
	})

	return list
}

func (f LedgerEntryList) MarshalJSON() ([]byte, error) {
	list := f.GetEntries()
	return json.Marshal(&list)
}

func (c *LedgerEntryList) Append(other LedgerEntryList) {
	c.list = append(c.list, other.list...)
}

func (c *LedgerEntryList) AddGrant(grant Grant) {
	c.list = append(c.list, LedgerEntry{
		ID:        grant.ID,
		Type:      LedgerEntryTypeGrant,
		Time:      grant.EffectiveAt,
		FeatureID: grant.FeatureID,
		Amount:    &grant.Amount,
	})
}

func (c *LedgerEntryList) AddVoidGrant(grant Grant) {
	c.list = append(c.list, LedgerEntry{
		ID:        grant.ParentID,
		Type:      LedgerEntryTypeVoid,
		Time:      grant.EffectiveAt,
		FeatureID: grant.FeatureID,
		Amount:    &grant.Amount,
	})
}

func (c *LedgerEntryList) AddReset(reset Reset) {
	c.list = append(c.list, LedgerEntry{
		ID:   reset.ID,
		Type: LedgerEntryTypeReset,
		Time: reset.EffectiveAt,
	})
}

func (c *LedgerEntryList) AddGrantUsage(grantBalance GrantBalance, from time.Time, to time.Time, amount decimal.Decimal) {
	now := time.Now()
	if to.After(now) {
		to = now
	}

	c.list = append(c.list, LedgerEntry{
		ID:        grantBalance.ID,
		Type:      LedgerEntryTypeGrantUsage,
		Time:      to,
		FeatureID: grantBalance.FeatureID,
		Amount:    &amount,
		Period: &Period{
			From: from,
			To:   to,
		},
	})
}
