// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/common"
	"github.com/bnb-chain/tss-lib/v3/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

var zero = big.NewInt(0)

// round 1 represents round 1 of the signing part of the GG18 ECDSA TSS spec (Gennaro, Goldfeder; 2018)
func newRound1(params *tss.Parameters, key *keygen.LocalPartySaveData, data *common.SignatureData, temp *localTempData, out chan<- tss.Message, end chan<- *common.SignatureData) tss.Round {
	_ = "STUB: not implemented"
	return *new(tss.Round)
}

func (round *round1) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// Spec requires calculate H(M) here,
// but considered different blockchain use different hash function we accept the converted big.Int
// if this big.Int is not belongs to Zq, the client might not comply with common rule (for ECDSA):
// https://github.com/btcsuite/btcd/blob/c26ffa870fd817666a857af1bf6498fabba1ffe3/btcec/signature.go#L263

// GG20 session binding: use caller-provided session nonce if available,
// otherwise fall back to the message hash for per-session SSID uniqueness.

// Use ssid + j (receiver's index) as Session context so that the verifier (party j)
// can reconstruct the same challenge using their ContextI = ssid + j in round 2.

func (round *round1) Update() (bool, *tss.Error) { _ = "STUB: not implemented"; return false, nil }

func (round *round1) CanAccept(msg tss.ParsedMessage) bool { _ = "STUB: not implemented"; return false }

func (round *round1) NextRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }

// ----- //

// helper to call into PrepareForSigning()
func (round *round1) prepare() error { _ = "STUB: not implemented"; return nil }

// adding the key derivation delta to the xi's
// Suppose x has shamir shares x_0,     x_1,     ..., x_n
// So x + D has shamir shares  x_0 + D, x_1 + D, ..., x_n + D
