package credit

import (
	"fmt"
	"time"
)

// Represents a point in time balance of grants
type GrantBalanceMap map[string]float64

func (g GrantBalanceMap) Copy() GrantBalanceMap {
	r := make(GrantBalanceMap, len(g))
	for k, v := range g {
		r[k] = v
	}
	return r
}

func (g GrantBalanceMap) Burn(grantID string, amount float64) {
	balance := g[grantID]
	g[grantID] = balance - amount
}

func (g GrantBalanceMap) Set(grantID string, amount float64) {
	g[grantID] = amount
}

// returns the combined balance of all grants
func (g GrantBalanceMap) Balance() float64 {
	var balance float64
	for _, v := range g {
		balance += v
	}
	return balance
}

// Whether the contents of the GrantBalanceMap exactly matches
// the list of provided grants.
// Return false if it has additional grants or if it misses any grants
func (g GrantBalanceMap) ExactlyForGrants(grants []Grant) bool {
	gmap := map[string]struct{}{}
	for _, grant := range grants {
		gmap[grant.ID] = struct{}{}
	}

	if len(gmap) != len(g) {
		return false
	}

	for k := range gmap {
		if _, ok := g[k]; !ok {
			return false
		}
	}
	return true
}

func (g GrantBalanceMap) ActiveBalances(grantMap map[string]Grant, at time.Time) (GrantBalanceMap, error) {
	filtered := GrantBalanceMap{}
	for grantID, grantBalance := range g {
		grant, ok := grantMap[grantID]
		// inconsistency check, shouldn't happen
		if !ok {
			return nil, fmt.Errorf("received balance for unknown grant: %s", grantID)
		}

		if !grant.ActiveAt(at) {
			continue
		}

		filtered.Set(grantID, grantBalance)
	}
	return filtered, nil
}

type GrantBalanceSnapshot struct {
	Balances GrantBalanceMap
	Overage  float64
	At       time.Time
}

func (g GrantBalanceSnapshot) Balance() float64 {
	return g.Balances.Balance()
}

// No balance has been saved since start of measurement for the owner
type GrantBalanceNoSavedBalanceForOwnerError struct {
	Owner NamespacedGrantOwner
	Time  time.Time
}

func (e GrantBalanceNoSavedBalanceForOwnerError) Error() string {
	return fmt.Sprintf("no saved balance for owner %s in namespace %s before %s", e.Owner.ID, e.Owner.Namespace, e.Time)
}
