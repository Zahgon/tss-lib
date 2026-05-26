// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package mta

import (
	"crypto/elliptic"
	"io"
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/crypto/paillier"
)

const (
	RangeProofAliceBytesParts = 6
)

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
)

type (
	RangeProofAlice struct {
		Z, U, W, S, S1, S2 *big.Int
	}
)

// ProveRangeAlice implements Alice's range proof used in the MtA and MtAwc protocols from GG18Spec (9) Fig. 9.
func ProveRangeAlice(Session []byte, ec elliptic.Curve, pk *paillier.PublicKey, c, NTilde, h1, h2, m, r *big.Int, rand io.Reader) (*RangeProofAlice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1.

// 2.

// 3.

// 4.

// 5.

// SECURITY: Use constant-time exponentiation for secret message m
// See: https://github.com/golang/go/issues/20654

// rho is random, not secret

// 6.

// 7.

// 8-9. e'

// must use RejectionSample

// s1 = e * m + alpha

// s2 = e * rho + gamma

func RangeProofAliceFromBytes(bzs [][]byte) (*RangeProofAlice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pf *RangeProofAlice) Verify(Session []byte, ec elliptic.Curve, pk *paillier.PublicKey, NTilde, h1, h2, c *big.Int) bool {
	_ = "STUB: not implemented"
	return false
}

// Reject c where gcd(c, N) != 1 to prevent nil pointer dereference from c^(-e) in modular exponentiation.
// When gcd(c, N) != 1, the modular inverse doesn't exist and big.Int.Exp returns nil.
// This also covers the c == 0 case.

// 3.

// 1-2. e'

// must use RejectionSample

// for the following conditionals

// 4. gamma^s_1 * s^N * c^-e

// u != (4)

// 5. h_1^s_1 * h_2^s_2 * z^-e

// w != (5)

func (pf *RangeProofAlice) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (pf *RangeProofAlice) Bytes() [RangeProofAliceBytesParts][]byte {
	_ = "STUB: not implemented"
	return nil
}
