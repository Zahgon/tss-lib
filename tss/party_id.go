// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package tss

import (
	"math/big"
)

type (
	// PartyID represents a participant in the TSS protocol rounds.
	// Note: The `id` and `moniker` are provided for convenience to allow you to track participants easier.
	// The `id` is intended to be a unique string representation of `key` and `moniker` can be anything (even left blank).
	PartyID struct {
		*MessageWrapper_PartyID
		Index int `json:"index"`
	}

	UnSortedPartyIDs []*PartyID
	SortedPartyIDs   []*PartyID
)

func (pid *PartyID) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

// --- ProtoBuf Extensions

func (mpid *MessageWrapper_PartyID) KeyInt() *big.Int { _ = "STUB: not implemented"; return nil }

// ----- //

// NewPartyID constructs a new PartyID
// Exported, used in `tss` client. `key` should remain consistent between runs for each party.
func NewPartyID(id, moniker string, key *big.Int) *PartyID { _ = "STUB: not implemented"; return nil }

// not known until sorted

func (pid PartyID) String() string { _ = "STUB: not implemented"; return "" }

// ----- //

// SortPartyIDs sorts a list of []*PartyID by their keys in ascending order
// Exported, used in `tss` client
func SortPartyIDs(ids UnSortedPartyIDs, startAt ...int) SortedPartyIDs {
	_ = "STUB: not implemented"
	return *new(SortedPartyIDs)
}

// assign party indexes

// GenerateTestPartyIDs generates a list of mock PartyIDs for tests
func GenerateTestPartyIDs(count int, startAt ...int) SortedPartyIDs {
	_ = "STUB: not implemented"
	return *new(SortedPartyIDs)
}

// default `i`

// this key makes tests more deterministic

func (spids SortedPartyIDs) Keys() []*big.Int { _ = "STUB: not implemented"; return nil }

func (spids SortedPartyIDs) ToUnSorted() UnSortedPartyIDs {
	_ = "STUB: not implemented"
	return *new(UnSortedPartyIDs)
}

func (spids SortedPartyIDs) FindByKey(key *big.Int) *PartyID { _ = "STUB: not implemented"; return nil }

func (spids SortedPartyIDs) Exclude(exclude *PartyID) SortedPartyIDs {
	_ = "STUB: not implemented"
	return *new(SortedPartyIDs)
}

// exclude

// Sortable

func (spids SortedPartyIDs) Len() int { _ = "STUB: not implemented"; return 0 }

func (spids SortedPartyIDs) Less(a, b int) bool { _ = "STUB: not implemented"; return false }

func (spids SortedPartyIDs) Swap(a, b int) { _ = "STUB: not implemented"; return }
