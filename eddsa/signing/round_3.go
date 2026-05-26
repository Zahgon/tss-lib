// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"github.com/bnb-chain/tss-lib/v3/tss"
)

func (round *round3) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// 1. init R

// 2-6. compute R

// 7. compute lambda

// h = hash512(k || A || M)

// 8. compute si

// 9. store r3 message pieces

// 10. broadcast si to other parties

func (round *round3) Update() (bool, *tss.Error) { _ = "STUB: not implemented"; return false, nil }

func (round *round3) CanAccept(msg tss.ParsedMessage) bool { _ = "STUB: not implemented"; return false }

func (round *round3) NextRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }
