// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package tss

import (
	"crypto/elliptic"

	s256k1 "github.com/btcsuite/btcd/btcec/v2"
	"github.com/decred/dcrd/dcrec/edwards/v2"
)

type CurveName string

const (
	Secp256k1 CurveName = "secp256k1"
	Ed25519   CurveName = "ed25519"
)

var (
	ec       elliptic.Curve
	registry map[CurveName]elliptic.Curve
)

// Init default curve (secp256k1)
func init() {
	ec = s256k1.S256()

	registry = make(map[CurveName]elliptic.Curve)
	registry[Secp256k1] = s256k1.S256()
	registry[Ed25519] = edwards.Edwards()
}

func RegisterCurve(name CurveName, curve elliptic.Curve) { _ = "STUB: not implemented"; return }

// return curve, exist(bool)
func GetCurveByName(name CurveName) (elliptic.Curve, bool) {
	_ = "STUB: not implemented"
	return *new(elliptic.Curve), false
}

// return name, exist(bool)
func GetCurveName(curve elliptic.Curve) (CurveName, bool) {
	_ = "STUB: not implemented"
	return *new(CurveName), false
}

// SameCurve returns true if both lhs and rhs are the same known curve
func SameCurve(lhs, rhs elliptic.Curve) bool { _ = "STUB: not implemented"; return false }

// if lhs/rhs not exist, return false

// EC returns the current elliptic curve in use. The default is secp256k1
func EC() elliptic.Curve {
	_ = "STUB: not implemented"

	// SetCurve sets the curve used by TSS. Must be called before Start. The default is secp256k1
	// Deprecated
	return *new(elliptic.Curve)
}

func SetCurve(curve elliptic.Curve) { _ = "STUB: not implemented"; return }

// secp256k1
func S256() elliptic.Curve { _ = "STUB: not implemented"; return *new(elliptic.Curve) }

func Edwards() elliptic.Curve { _ = "STUB: not implemented"; return *new(elliptic.Curve) }
