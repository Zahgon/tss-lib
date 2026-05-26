// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package keygen

import (
	"github.com/bnb-chain/tss-lib/v3/tss"
)

const (
	paillierBitsLen = 2048
)

func (round *round2) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// 6. verify dln proofs, store r1 message pieces, ensure uniqueness of h1j, h2j

// save NTilde_j, h1_j, h2_j, ...

// used in round 4

// 5. p2p send share ij to Pj

// do not send to this Pj, but store for round 3

// 7. BROADCAST de-commitments of Shamir poly*G

func (round *round2) CanAccept(msg tss.ParsedMessage) bool { _ = "STUB: not implemented"; return false }

func (round *round2) Update() (bool, *tss.Error) {
	_ = "STUB: not implemented"
	// guard - VERIFY de-commit for all Pj
	return false, nil
}

func (round *round2) NextRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }
