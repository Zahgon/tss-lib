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

// Alice_end

// Alice_end_wc

// consume error channels; wait for goroutines

// SECURITY: Use constant-time multiplication for secret values k, gamma, w

func (round *round3) Update() (bool, *tss.Error) { _ = "STUB: not implemented"; return false, nil }

func (round *round3) CanAccept(msg tss.ParsedMessage) bool { _ = "STUB: not implemented"; return false }

func (round *round3) NextRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }
