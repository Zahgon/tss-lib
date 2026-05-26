// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package resharing

import (
	"github.com/bnb-chain/tss-lib/v3/tss"
)

func (round *round3) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// resets both round.oldOK and round.newOK

// 1-2. send share to Pj from the new committee

// 3. broadcast de-commitment to new committees

func (round *round3) CanAccept(msg tss.ParsedMessage) bool { _ = "STUB: not implemented"; return false }

func (round *round3) Update() (bool, *tss.Error) {
	_ = "STUB: not implemented"
	// only the new committee receive in this round
	return false, nil
}

// accept messages from old -> new committee

func (round *round3) NextRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }
