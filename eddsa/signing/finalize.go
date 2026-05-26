// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"github.com/bnb-chain/tss-lib/v3/tss"
)

func (round *finalization) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// save the signature for final output

func (round *finalization) CanAccept(msg tss.ParsedMessage) bool {
	_ = "STUB: not implemented"
	// not expecting any incoming messages in this round
	return false
}

func (round *finalization) Update() (bool, *tss.Error) {
	_ = "STUB: not implemented"
	// not expecting any incoming messages in this round
	return false, nil
}

func (round *finalization) NextRound() tss.Round {
	_ = "STUB: not implemented"
	// finished!
	return *new(tss.Round)
}
