// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package resharing

import (
	"github.com/bnb-chain/tss-lib/v3/tss"
)

func (round *round4) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// resets both round.oldOK and round.newOK

// both committees proceed to round 5 after receiving "ACK" messages from the new committee

// 1-3. verify paillier & dln proofs, store message pieces, ensure uniqueness of h1j, h2j

// who caused the error(s)

// save NTilde_j, h1_j, h2_j received in NewCommitteeStep1 here

// 4.

// 5-9.

// P1..P_t+1. Ps are indexed from 0 here
// 6-7.

// 6. unpack flat "v" commitment content

// they're points so * 2
// TODO collect culprits and return a list of them as per convention

// 8.

// TODO collect culprits and return a list of them as per convention

// 9.

// 10-13.

// 14.

// 15-19.

// who caused the error(s)

// Send facProof to new parties

// Send an "ACK" message to both committees to signal that we're ready to save our data

func (round *round4) CanAccept(msg tss.ParsedMessage) bool { _ = "STUB: not implemented"; return false }

func (round *round4) Update() (bool, *tss.Error) {
	_ = "STUB: not implemented"
	// accept messages from new -> old&new committees
	return false, nil
}

func (round *round4) NextRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }
