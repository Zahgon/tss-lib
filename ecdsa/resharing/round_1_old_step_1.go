// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package resharing

import (
	"github.com/bnb-chain/tss-lib/v3/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

// round 1 represents round 1 of the keygen part of the GG18 ECDSA TSS spec (Gennaro, Goldfeder; 2018)
func newRound1(params *tss.ReSharingParameters, input, save *keygen.LocalPartySaveData, temp *localTempData, out chan<- tss.Message, end chan<- *keygen.LocalPartySaveData) tss.Round {
	_ = "STUB: not implemented"
	return *new(tss.Round)
}

func (round *round1) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// resets both round.oldOK and round.newOK

// GG20 session binding: use caller-provided session nonce if available.
// For resharing, all parties must agree on the nonce via external coordination
// (e.g., coordinator-assigned session ID) since no shared session-unique
// value is available within the protocol itself.

// 1. PrepareForSigning() -> w_i

// 2.

// 3.

// 4. populate temp data

// 5. "broadcast" C_i to members of the NEW committee

func (round *round1) CanAccept(msg tss.ParsedMessage) bool {
	_ = "STUB: not implemented"
	// accept messages from old -> new committee
	return false
}

func (round *round1) Update() (bool, *tss.Error) {
	_ = "STUB: not implemented"
	// only the new committee receive in this round
	return false, nil
}

// accept messages from old -> new committee

// save the ecdsa pub received from the old committee

// uh oh - anomaly!

func (round *round1) NextRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }
