// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package common

import (
	"io"
	"math/big"
)

const (
	mustGetRandomIntMaxBits = 5000
)

// MustGetRandomInt panics if it is unable to gather entropy from `io.Reader` or when `bits` is <= 0
func MustGetRandomInt(rand io.Reader, bits int) *big.Int { _ = "STUB: not implemented"; return nil }

// Max random value e.g. 2^256 - 1

// Generate cryptographically strong pseudo-random int between 0 - max

func GetRandomPositiveInt(rand io.Reader, lessThan *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func GetRandomPrimeInt(rand io.Reader, bits int) *big.Int { _ = "STUB: not implemented"; return nil }

// fallback to older method

// Generate a random element in the group of all the elements in Z/nZ that
// has a multiplicative inverse.
func GetRandomPositiveRelativelyPrimeInt(rand io.Reader, n *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func IsNumberInMultiplicativeGroup(n, v *big.Int) bool { _ = "STUB: not implemented"; return false }

//	Return a random generator of RQn with high probability.
//	THIS METHOD ONLY WORKS IF N IS THE PRODUCT OF TWO SAFE PRIMES!
//
// https://github.com/didiercrunch/paillier/blob/d03e8850a8e4c53d04e8016a2ce8762af3278b71/utils.go#L39
func GetRandomGeneratorOfTheQuadraticResidue(rand io.Reader, n *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// GetRandomQuadraticNonResidue returns a quadratic non residue of odd n.
func GetRandomQuadraticNonResidue(rand io.Reader, n *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// GetRandomBytes returns random bytes of length.
func GetRandomBytes(rand io.Reader, length int) ([]byte, error) {
	_ = "STUB: not implemented"
	// Per [BIP32], the seed must be in range [MinSeedBytes, MaxSeedBytes].
	return nil, nil
}
