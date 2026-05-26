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

	"github.com/bnb-chain/tss-lib/v3/crypto"
	"github.com/bnb-chain/tss-lib/v3/crypto/paillier"
)

const (
	ProofBobBytesParts   = 10
	ProofBobWCBytesParts = 12
)

type (
	ProofBob struct {
		Z, ZPrm, T, V, W, S, S1, S2, T1, T2 *big.Int
	}

	ProofBobWC struct {
		*ProofBob
		U *crypto.ECPoint
	}
)

// ProveBobWC implements Bob's proof both with or without check "ProveMtawc_Bob" and "ProveMta_Bob" used in the MtA protocol from GG18Spec (9) Figs. 10 & 11.
// an absent `X` generates the proof without the X consistency check X = g^x
func ProveBobWC(Session []byte, ec elliptic.Curve, pk *paillier.PublicKey, NTilde, h1, h2, c1, c2, x, y, r *big.Int, X *crypto.ECPoint, rand io.Reader) (*ProofBobWC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// steps are numbered as shown in Fig. 10, but diverge slightly for Fig. 11
// 1.

// 2.

// 3.

// 4.

// 5.
// initialization suppresses an IDE warning

// 6.

// SECURITY: Use constant-time exponentiation for secret x
// See: https://github.com/golang/go/issues/20654

// rho is random, not secret

// 7.

// 8.

// SECURITY: Use constant-time exponentiation for secret y

// sigma is random, not secret

// 9.

// 10.

// 11-12. e'

// must use RejectionSample

// X is nil if called by ProveBob (Bob's proof "without check")

// 13.

// 14.

// 15.

// 16.

// 17.

// the regular Bob proof ("without check") is extracted and returned by ProveBob

// or the WC ("with check") version is used in round 2 of the signing protocol

// ProveBob implements Bob's proof "ProveMta_Bob" used in the MtA protocol from GG18Spec (9) Fig. 11.
func ProveBob(Session []byte, ec elliptic.Curve, pk *paillier.PublicKey, NTilde, h1, h2, c1, c2, x, y, r *big.Int, rand io.Reader) (*ProofBob, error) {
	_ = "STUB: not implemented"
	// the Bob proof ("with check") contains the ProofBob "without check"; this method extracts and returns it
	// X is supplied as nil to exclude it from the proof hash
	return nil, nil
}

func ProofBobWCFromBytes(ec elliptic.Curve, bzs [][]byte) (*ProofBobWC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ProofBobFromBytes(bzs [][]byte) (*ProofBob, error) { _ = "STUB: not implemented"; return nil, nil }

// ProveBobWC.Verify implements verification of Bob's proof with check "VerifyMtawc_Bob" used in the MtA protocol from GG18Spec (9) Fig. 10.
// an absent `X` verifies a proof generated without the X consistency check X = g^x
func (pf *ProofBobWC) Verify(Session []byte, ec elliptic.Curve, pk *paillier.PublicKey, NTilde, h1, h2, c1, c2 *big.Int, X *crypto.ECPoint) bool {
	_ = "STUB: not implemented"
	return false
}

// q^2
// q^3
// q^6
// q^7

// 3.

// 1-2. e'

// must use RejectionSample

// X is nil if called on a ProveBob (Bob's proof "without check")

// for the following conditionals

// 4. runs only in the "with check" mode from Fig. 10

// 5-6.

// 5.

// 6.

// 7.

// ProveBob.Verify implements verification of Bob's proof without check "VerifyMta_Bob" used in the MtA protocol from GG18Spec (9) Fig. 11.
func (pf *ProofBob) Verify(Session []byte, ec elliptic.Curve, pk *paillier.PublicKey, NTilde, h1, h2, c1, c2 *big.Int) bool {
	_ = "STUB: not implemented"
	return false
}

func (pf *ProofBob) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (pf *ProofBobWC) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (pf *ProofBob) Bytes() [ProofBobBytesParts][]byte { _ = "STUB: not implemented"; return nil }

func (pf *ProofBobWC) Bytes() [ProofBobWCBytesParts][]byte { _ = "STUB: not implemented"; return nil }
