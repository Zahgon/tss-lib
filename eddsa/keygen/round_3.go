// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package keygen

import (
	"github.com/bnb-chain/tss-lib/v3/tss"
)

func (round *round3) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// 1,10. calculate xi

// 2-3.

// ours

// 4-12.

// 6-9.

// 4-10.

// (9) handled above

// consume unbuffered channels (end the goroutines)

// who caused the error(s)

// collect culprits to error out with

// who caused the error(s)

// 11-12.

// 13-17. compute Xj for each Pj

// who caused the error(s)

// 18. compute and SAVE the EDDSA public key `y`

// PRINT public key & private share

func (round *round3) CanAccept(msg tss.ParsedMessage) bool {
	_ = "STUB: not implemented"
	// not expecting any incoming messages in this round
	return false
}

func (round *round3) Update() (bool, *tss.Error) {
	_ = "STUB: not implemented"
	// not expecting any incoming messages in this round
	return false, nil
}

func (round *round3) NextRound() tss.Round {
	_ = "STUB: not implemented"
	// finished!
	return *new(tss.Round)
}
