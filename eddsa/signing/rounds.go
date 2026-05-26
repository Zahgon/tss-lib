// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"github.com/bnb-chain/tss-lib/v3/common"
	"github.com/bnb-chain/tss-lib/v3/eddsa/keygen"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

const (
	TaskName = "eddsa-signing"
)

type (
	base struct {
		*tss.Parameters
		key     *keygen.LocalPartySaveData
		data    *common.SignatureData
		temp    *localTempData
		out     chan<- tss.Message
		end     chan<- *common.SignatureData
		ok      []bool // `ok` tracks parties which have been verified by Update()
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
	finalization struct {
		*round3
	}
)

var (
	_ tss.Round = (*round1)(nil)
	_ tss.Round = (*round2)(nil)
	_ tss.Round = (*round3)(nil)
	_ tss.Round = (*finalization)(nil)
)

// ----- //

func (round *base) Params() *tss.Parameters { _ = "STUB: not implemented"; return nil }

func (round *base) RoundNumber() int { _ = "STUB: not implemented"; return 0 }

// CanProceed is inherited by other rounds
func (round *base) CanProceed() bool { _ = "STUB: not implemented"; return false }

// WaitingFor is called by a Party for reporting back to the caller
func (round *base) WaitingFor() []*tss.PartyID { _ = "STUB: not implemented"; return nil }

func (round *base) WrapError(err error, culprits ...*tss.PartyID) *tss.Error {
	_ = "STUB: not implemented"
	return nil
}

// ----- //

// `ok` tracks parties which have been verified by Update()
func (round *base) resetOK() { _ = "STUB: not implemented"; return }

// get ssid from local params
func (round *base) getSSID() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ec curve
// parties

// BigXj
// round number
