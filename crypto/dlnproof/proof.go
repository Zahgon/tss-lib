// Copyright © 2019-2020 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

// Zero-knowledge proof of knowledge of the discrete logarithm over safe prime product

// A proof of knowledge of the discrete log of an element h2 = hx1 with respect to h1.
// In our protocol, we will run two of these in parallel to prove that two elements h1,h2 generate the same group modN.

package dlnproof

import (
	"io"
	"math/big"
)

const Iterations = 128

type (
	Proof struct {
		Alpha,
		T [Iterations]*big.Int
	}
)

var one = big.NewInt(1)

func NewDLNProof(Session []byte, h1, h2, x, p, q, N *big.Int, rand io.Reader) *Proof {
	_ = "STUB: not implemented"
	return nil
}

// SECURITY: Use constant-time exponentiation

// SECURITY: Use constant-time multiplication for secret x

func (p *Proof) Verify(Session []byte, h1, h2, N *big.Int) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Proof) Serialize() ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func UnmarshalDLNProof(bzs [][]byte) (*Proof, error) { _ = "STUB: not implemented"; return nil, nil }
