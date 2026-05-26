// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package keygen

import (
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/tss"
)

var zero = big.NewInt(0)

// round 1 represents round 1 of the keygen part of the EDDSA TSS spec
func newRound1(params *tss.Parameters, save *LocalPartySaveData, temp *localTempData, out chan<- tss.Message, end chan<- *LocalPartySaveData) tss.Round {
	_ = "STUB: not implemented"
	return *new(tss.Round)
}

func (round *round1) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// GG20 session binding: use caller-provided session nonce if available.
// For keygen, all parties must agree on the nonce via external coordination
// (e.g., coordinator-assigned session ID) since no shared session-unique
// value is available within the protocol itself.

// 1. calculate "partial" key share ui

// 2. compute the vss shares

// security: the original u_i may be discarded
// clears the secret data from memory
// silences a linter warning

// 3. make commitment -> (C, D)

// for this P: SAVE
// - shareID
// and keep in temporary storage:
// - VSS Vs
// - our set of Shamir shares

// BROADCAST commitments

func (round *round1) CanAccept(msg tss.ParsedMessage) bool { _ = "STUB: not implemented"; return false }

func (round *round1) Update() (bool, *tss.Error) { _ = "STUB: not implemented"; return false, nil }

// vss check is in round 2

func (round *round1) NextRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }
