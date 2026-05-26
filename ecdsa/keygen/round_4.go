// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package keygen

import (
	"github.com/bnb-chain/tss-lib/v3/tss"
)

func (round *round4) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// 1-3. (concurrent)
// r3 messages are assumed to be available and != nil in this function

// consume unbuffered channels (end the goroutines)

// who caused the error(s)

func (round *round4) CanAccept(msg tss.ParsedMessage) bool {
	_ = "STUB: not implemented"
	// not expecting any incoming messages in this round
	return false
}

func (round *round4) Update() (bool, *tss.Error) {
	_ = "STUB: not implemented"
	// not expecting any incoming messages in this round
	return false, nil
}

func (round *round4) NextRound() tss.Round {
	_ = "STUB: not implemented"
	// finished!
	return *new(tss.Round)
}
