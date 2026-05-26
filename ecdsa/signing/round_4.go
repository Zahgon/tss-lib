// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"github.com/bnb-chain/tss-lib/v3/tss"
)

func (round *round4) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// compute the multiplicative inverse thelta mod q

// SECURITY: Use constant-time modular inverse for secret theta
// See: https://github.com/golang/go/issues/20654

func (round *round4) Update() (bool, *tss.Error) { _ = "STUB: not implemented"; return false, nil }

func (round *round4) CanAccept(msg tss.ParsedMessage) bool { _ = "STUB: not implemented"; return false }

func (round *round4) NextRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }
