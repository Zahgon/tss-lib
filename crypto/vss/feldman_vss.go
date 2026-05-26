// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

// Feldman VSS, based on Paul Feldman, 1987., A practical scheme for non-interactive verifiable secret sharing.
// In Foundations of Computer Science, 1987., 28th Annual Symposium on. IEEE, 427–43
//

package vss

import (
	"crypto/elliptic"
	"fmt"
	"io"
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/crypto"
)

type (
	Share struct {
		Threshold int
		ID,       // xi
		Share *big.Int // Sigma i
	}

	Vs []*crypto.ECPoint // v0..vt

	Shares []*Share
)

var (
	ErrNumSharesBelowThreshold = fmt.Errorf("not enough shares to satisfy the threshold")

	zero = big.NewInt(0)
	one  = big.NewInt(1)
)

// Check share ids of Shamir's Secret Sharing, return error if duplicate or 0 value found
func CheckIndexes(ec elliptic.Curve, indexes []*big.Int) ([]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a new array of secret shares created by Shamir's Secret Sharing Algorithm,
// requiring a minimum number of shares to recreate, of length shares, from the input secret
func Create(ec elliptic.Curve, threshold int, secret *big.Int, indexes []*big.Int, rand io.Reader) (Vs, Shares, error) {
	_ = "STUB: not implemented"
	return *new(Vs), *new(Shares), nil
}

func (share *Share) Verify(ec elliptic.Curve, threshold int, vs Vs) bool {
	_ = "STUB: not implemented"
	return false
}

// YRO : we need to have our accumulator outside of the loop

// t = k_i^j

// v = v * v_j^t

func (shares Shares) ReConstruct(ec elliptic.Curve) (secret *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// x coords

func samplePolynomial(ec elliptic.Curve, threshold int, secret *big.Int, rand io.Reader) []*big.Int {
	_ = "STUB: not implemented"
	return nil
}

// Evauluates a polynomial with coefficients such that:
// evaluatePolynomial([a, b, c, d], x):
//
//	returns a + bx + cx^2 + dx^3
func evaluatePolynomial(ec elliptic.Curve, threshold int, v []*big.Int, id *big.Int) (result *big.Int) {
	_ = "STUB: not implemented"
	return nil
}
