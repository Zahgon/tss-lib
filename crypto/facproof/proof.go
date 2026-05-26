// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package facproof

import (
	"crypto/elliptic"
	"io"
	"math/big"
)

const (
	ProofFacBytesParts = 11
)

type (
	ProofFac struct {
		P, Q, A, B, T, Sigma, Z1, Z2, W1, W2, V *big.Int
	}
)

var (
	// rangeParameter l limits the bits of p or q to be in [1024-l, 1024+l]
	rangeParameter = new(big.Int).Lsh(big.NewInt(1), 15)
	one            = big.NewInt(1)
)

// NewProof implements prooffac
func NewProof(Session []byte, ec elliptic.Curve, N0, NCap, s, t, N0p, N0q *big.Int, rand io.Reader) (*ProofFac, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fig 28.1 sample

// Fig 28.1 compute

// SECURITY: Use constant-time exponentiation for secret exponents N0p, N0q
// See: https://github.com/golang/go/issues/20654

// P = s^N0p * t^mu mod NCap
// mu is random, not secret

// Q = s^N0q * t^nu mod NCap
// nu is random, not secret

// A, B, T use random exponents (alpha, beta, x, y, r) - non-secret, regular exp is fine

// Fig 28.2 e

// Fig 28.3
// z1 = e * N0p + alpha (reveals N0p in the output, but that's part of the protocol)

// z2 = e * N0q + beta

func NewProofFromBytes(bzs [][]byte) (*ProofFac, error) { _ = "STUB: not implemented"; return nil, nil }

func (pf *ProofFac) Verify(Session []byte, ec elliptic.Curve, N0, NCap, s, t *big.Int) bool {
	_ = "STUB: not implemented"
	return false
}

// Fig 28. Range Check

// Fig 28. Equality Check

func (pf *ProofFac) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (pf *ProofFac) Bytes() [ProofFacBytesParts][]byte { _ = "STUB: not implemented"; return nil }
