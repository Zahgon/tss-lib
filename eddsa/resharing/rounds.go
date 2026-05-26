// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package resharing

import (
	"github.com/bnb-chain/tss-lib/v3/eddsa/keygen"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

const (
	TaskName = "eddsa-resharing"
)

type (
	base struct {
		*tss.ReSharingParameters
		temp        *localTempData
		input, save *keygen.LocalPartySaveData
		out         chan<- tss.Message
		end         chan<- *keygen.LocalPartySaveData
		oldOK,      // old committee "ok" tracker
		newOK []bool // `ok` tracks parties which have been verified by Update(); this one is for the new committee
		started bool
		number  int
	}
	round1 struct {
		*base
	}
	round2 struct {
		*round1
	}
	round3 struct {
		*round2
	}
	round4 struct {
		*round3
	}
	round5 struct {
		*round4
	}
)

var (
	_ tss.Round = (*round1)(nil)
	_ tss.Round = (*round2)(nil)
	_ tss.Round = (*round3)(nil)
	_ tss.Round = (*round4)(nil)
	_ tss.Round = (*round5)(nil)
)

// ----- //

func (round *base) Params() *tss.Parameters { _ = "STUB: not implemented"; return nil }

func (round *base) ReSharingParams() *tss.ReSharingParameters {
	_ = "STUB: not implemented"
	return nil
}

func (round *base) RoundNumber() int { _ = "STUB: not implemented"; return 0 }

// CanProceed is inherited by other rounds
func (round *base) CanProceed() bool { _ = "STUB: not implemented"; return false }

// WaitingFor is called by a Party for reporting back to the caller
func (round *base) WaitingFor() []*tss.PartyID { _ = "STUB: not implemented"; return nil }

// consolidate into the list

func (round *base) WrapError(err error, culprits ...*tss.PartyID) *tss.Error {
	_ = "STUB: not implemented"
	return nil
}

// ----- //

// `oldOK` tracks parties which have been verified by Update()
func (round *base) resetOK() { _ = "STUB: not implemented"; return }

// sets all pairings in `oldOK` to true
func (round *base) allOldOK() { _ = "STUB: not implemented"; return }

// sets all pairings in `newOK` to true
func (round *base) allNewOK() { _ = "STUB: not implemented"; return }
