// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"

	"github.com/decred/dcrd/dcrec/edwards/v2"
)

// ECPoint convenience helper
type ECPoint struct {
	curve  elliptic.Curve
	coords [2]*big.Int
}

var (
	eight    = big.NewInt(8)
	eightInv = new(big.Int).ModInverse(eight, edwards.Edwards().Params().N)
)

// Creates a new ECPoint and checks that the given coordinates are on the elliptic curve.
func NewECPoint(curve elliptic.Curve, X, Y *big.Int) (*ECPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Creates a new ECPoint without checking that the coordinates are on the elliptic curve.
// Only use this function when you are completely sure that the point is already on the curve.
func NewECPointNoCurveCheck(curve elliptic.Curve, X, Y *big.Int) *ECPoint {
	_ = "STUB: not implemented"
	return nil
}

func (p *ECPoint) X() *big.Int { _ = "STUB: not implemented"; return nil }

func (p *ECPoint) Y() *big.Int { _ = "STUB: not implemented"; return nil }

func (p *ECPoint) Add(p1 *ECPoint) (*ECPoint, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *ECPoint) ScalarMult(k *big.Int) *ECPoint { _ = "STUB: not implemented"; return nil }

// it must be on the curve, no need to check.

func (p *ECPoint) ToECDSAPubKey() *ecdsa.PublicKey { _ = "STUB: not implemented"; return nil }

func (p *ECPoint) IsOnCurve() bool { _ = "STUB: not implemented"; return false }

func (p *ECPoint) Curve() elliptic.Curve { _ = "STUB: not implemented"; return *new(elliptic.Curve) }

func (p *ECPoint) Equals(p2 *ECPoint) bool { _ = "STUB: not implemented"; return false }

func (p *ECPoint) SetCurve(curve elliptic.Curve) *ECPoint { _ = "STUB: not implemented"; return nil }

func (p *ECPoint) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (p *ECPoint) EightInvEight() *ECPoint { _ = "STUB: not implemented"; return nil }

func ScalarBaseMult(curve elliptic.Curve, k *big.Int) *ECPoint {
	_ = "STUB: not implemented"
	return nil
}

// it must be on the curve, no need to check.

func isOnCurve(c elliptic.Curve, x, y *big.Int) bool { _ = "STUB: not implemented"; return false }

// Reject coordinates outside [0, P) to prevent non-canonical point representations
// from bypassing the curve equation check via modular reduction (SRC-2026-573).

// ----- //

func FlattenECPoints(in []*ECPoint) ([]*big.Int, error) { _ = "STUB: not implemented"; return nil, nil }

func UnFlattenECPoints(curve elliptic.Curve, in []*big.Int, noCurveCheck ...bool) ([]*ECPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ----- //
// Gob helpers for if you choose to encode messages with Gob.

func (p *ECPoint) GobEncode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *ECPoint) GobDecode(buf []byte) error { _ = "STUB: not implemented"; return nil }

// ----- //

// crypto.ECPoint is not inherently json marshal-able
func (p *ECPoint) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *ECPoint) UnmarshalJSON(payload []byte) error { _ = "STUB: not implemented"; return nil }

// forward compatible, use global ec as default value
