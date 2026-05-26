// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package resharing

import (
	"github.com/bnb-chain/tss-lib/v3/tss"
)

func (round *round5) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// 21.
// for this P: SAVE data

// misc: build list of paillier public keys to save

func (round *round5) CanAccept(msg tss.ParsedMessage) bool { _ = "STUB: not implemented"; return false }

func (round *round5) Update() (bool, *tss.Error) { _ = "STUB: not implemented"; return false, nil }

func (round *round5) NextRound() tss.Round {
	_ = "STUB: not implemented"
	// both committees are finished!
	return *new(tss.Round)
}
