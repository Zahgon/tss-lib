// Copyright © 2019-2023 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package modproof

import (
	"io"
	"math/big"
)

const (
	Iterations         = 80
	ProofModBytesParts = Iterations*2 + 3
)

var one = big.NewInt(1)

type (
	ProofMod struct {
		W *big.Int
		X [Iterations]*big.Int
		A *big.Int
		B *big.Int
		Z [Iterations]*big.Int
	}
)

// isQuadraticResidue checks Euler criterion
func isQuadraticResidue(X, N *big.Int) bool { _ = "STUB: not implemented"; return false }

func NewProof(Session []byte, N, P, Q *big.Int, rand io.Reader) (*ProofMod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fig 16.1

// Fig 16.2

// Fig 16.3

// N^(-1) mod Phi: Phi is even so bigmod (which requires odd modulus) cannot be used.
// This is a prover-side computation where we already hold P, Q — timing leakage from
// ModInverse here doesn't expose secrets to external observers.

// Fix bitLen of A and B

// for fourth-root: expo = ((Phi + 4) / 8)^2 mod Phi

// Use constant-time exponentiation with secret-derived exponents

func NewProofFromBytes(bzs [][]byte) (*ProofMod, error) { _ = "STUB: not implemented"; return nil, nil }

func (pf *ProofMod) Verify(Session []byte, N *big.Int) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: add basic properties checker

// Fig 16. Verification

func (pf *ProofMod) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (pf *ProofMod) Bytes() [ProofModBytesParts][]byte { _ = "STUB: not implemented"; return nil }
