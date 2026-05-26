// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package resharing

import (
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/tss"
)

var zero = big.NewInt(0)

func (round *round2) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// resets both round.oldOK and round.newOK

// check consistency of SSID

// 2. "broadcast" "ACK" members of the OLD committee

// 1.
// generate Paillier public key E_i, private key and proof
// generate safe primes for ZKPs later on
// compute ntilde, h1, h2 (uses safe primes)
// use the pre-params if they were provided to the LocalParty constructor

// generate the dlnproofs for resharing

// for this P: SAVE de-commitments, paillier keys for round 2

func (round *round2) CanAccept(msg tss.ParsedMessage) bool { _ = "STUB: not implemented"; return false }

func (round *round2) Update() (bool, *tss.Error) { _ = "STUB: not implemented"; return false, nil }

// accept messages from new -> old committee

// accept message from new -> committee

// accept messages from new -> old committee

// accept messages from new -> new committee

func (round *round2) NextRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }
