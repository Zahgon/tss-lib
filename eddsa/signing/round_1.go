// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"github.com/bnb-chain/tss-lib/v3/common"
	"github.com/bnb-chain/tss-lib/v3/eddsa/keygen"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

// round 1 represents round 1 of the signing part of the EDDSA TSS spec
func newRound1(params *tss.Parameters, key *keygen.LocalPartySaveData, data *common.SignatureData, temp *localTempData, out chan<- tss.Message, end chan<- *common.SignatureData) tss.Round {
	_ = "STUB: not implemented"
	return *new(tss.Round)
}

func (round *round1) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// GG20 session binding: use caller-provided session nonce if available,
// otherwise fall back to the message hash for per-session SSID uniqueness.

// 1. select ri

// 2. make commitment

// 3. store r1 message pieces

// 4. broadcast commitment

func (round *round1) Update() (bool, *tss.Error) { _ = "STUB: not implemented"; return false, nil }

func (round *round1) CanAccept(msg tss.ParsedMessage) bool { _ = "STUB: not implemented"; return false }

func (round *round1) NextRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }

// ----- //

// helper to call into PrepareForSigning()
func (round *round1) prepare() error { _ = "STUB: not implemented"; return nil }
