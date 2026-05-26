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

// 1,9. calculate xi

// 2-3.

// ours

// 4-11.

// 6-8.

// 4-9.

// For old parties, the modProof could be not exist
// Not return error for compatibility reason

// For old parties, the facProof could be not exist
// Not return error for compatibility reason

// (9) handled above

// consume unbuffered channels (end the goroutines)

// who caused the error(s)

// collect culprits to error out with

// who caused the error(s)

// 10-11.

// 12-16. compute Xj for each Pj

// who caused the error(s)

// 17. compute and SAVE the ECDSA public key `y`

// PRINT public key & private share

// BROADCAST paillier proof for Pi

func (round *round3) CanAccept(msg tss.ParsedMessage) bool { _ = "STUB: not implemented"; return false }

func (round *round3) Update() (bool, *tss.Error) { _ = "STUB: not implemented"; return false, nil }

// proof check is in round 4

func (round *round3) NextRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }
