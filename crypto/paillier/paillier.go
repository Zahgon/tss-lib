// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

// The Paillier Crypto-system is an additive crypto-system. This means that given two ciphertexts, one can perform operations equivalent to adding the respective plain texts.
// Additionally, Paillier Crypto-system supports further computations:
//
// * Encrypted integers can be added together
// * Encrypted integers can be multiplied by an unencrypted integer
// * Encrypted integers and unencrypted integers can be added together
//
// Implementation adheres to GG18Spec (6)

package paillier

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"sync"

	"github.com/otiai10/primes"

	crypto2 "github.com/bnb-chain/tss-lib/v3/crypto"
)

const (
	ProofIters         = 13
	verifyPrimesUntil  = 1000 // Verify uses primes <1000
	pQBitLenDifference = 3    // >1020-bit P-Q
)

type (
	PublicKey struct {
		N *big.Int
	}

	PrivateKey struct {
		PublicKey
		LambdaN, // lcm(p-1, q-1)
		PhiN *big.Int // (p-1) * (q-1)
		P, Q *big.Int

		// cached M = N^(-1) mod PhiN, lazily computed
		m     *big.Int
		mOnce sync.Once
	}

	// Proof uses the new GenerateXs method in GG18Spec (6)
	Proof [ProofIters]*big.Int
)

var (
	ErrMessageTooLong   = fmt.Errorf("the message is too large or < 0")
	ErrMessageMalFormed = fmt.Errorf("the message is mal-formed")

	zero = big.NewInt(0)
	one  = big.NewInt(1)
)

func init() {
	// init primes cache
	_ = primes.Globally.Until(verifyPrimesUntil)
}

// len is the length of the modulus (each prime = len / 2)
func GenerateKeyPair(ctx context.Context, rand io.Reader, modulusBitLen int, optionalConcurrency ...int) (privateKey *PrivateKey, publicKey *PublicKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// KS-BTL-F-03: use two safe primes for P, Q

// KS-BTL-F-03: check that p-q is also very large in order to avoid square-root attacks

// phiN = P-1 * Q-1

// lambdaN = lcm(P−1, Q−1)

// ----- //

func (publicKey *PublicKey) EncryptAndReturnRandomness(rand io.Reader, m *big.Int) (c *big.Int, x *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// m < 0 || m >= N ?

// 1. gamma^m mod N2

// 2. x^N mod N2

// 3. (1) * (2) mod N2

func (publicKey *PublicKey) Encrypt(rand io.Reader, m *big.Int) (c *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (publicKey *PublicKey) HomoMult(m, c1 *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// m < 0 || m >= N ?

// c1 < 0 || c1 >= N2 ?

// cipher^m mod N2

func (publicKey *PublicKey) HomoAdd(c1, c2 *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// c1 < 0 || c1 >= N2 ?

// c2 < 0 || c2 >= N2 ?

// c1 * c2 mod N2

func (publicKey *PublicKey) NSquare() *big.Int { _ = "STUB: not implemented"; return nil }

// AsInts returns the PublicKey serialised to a slice of *big.Int for hashing
func (publicKey *PublicKey) AsInts() []*big.Int { _ = "STUB: not implemented"; return nil }

// Gamma returns N+1
func (publicKey *PublicKey) Gamma() *big.Int { _ = "STUB: not implemented"; return nil }

// ----- //

func (privateKey *PrivateKey) Decrypt(c *big.Int) (m *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// c < 0 || c >= N2 ?

// SECURITY: Use constant-time exponentiation to prevent timing side-channels.
// The original code used math/big.Exp which leaks information about the secret
// exponent LambdaN through execution time variations.
// See: https://github.com/golang/go/issues/20654

// Standard (non-constant-time) implementation for better performance

// 1. L(u) = (c^LambdaN-1 mod N2) / N

// 2. L(u) = (Gamma^LambdaN-1 mod N2) / N

// 3. (1) * modInv(2) mod N

// SECURITY: Use constant-time ModInverse to prevent timing side-channels.
// Lg is derived from secret LambdaN exponentiation.
// N = P*Q is composite, so we must provide phi(N) for Euler's theorem.

// M returns the cached value of N^(-1) mod PhiN, computing it on first call.
func (privateKey *PrivateKey) M() *big.Int { _ = "STUB: not implemented"; return nil }

// ----- //

// Proof is an implementation of Gennaro, R., Micciancio, D., Rabin, T.:
// An efficient non-interactive statistical zero-knowledge proof system for quasi-safe prime products.
// In: In Proc. of the 5th ACM Conference on Computer and Communications Security (CCS-98. Citeseer (1998)

func (privateKey *PrivateKey) Proof(k *big.Int, ecdsaPub *crypto2.ECPoint) Proof {
	_ = "STUB: not implemented"
	return *new(Proof)
}

// M = N^(-1) mod PhiN is precomputed and cached on the private key.

// SECURITY: Use constant-time exponentiation for xs[i]^M mod N.
// M is derived from secret PhiN, so we must use constant-time Exp.
// N is odd, so bigmod works correctly here.

// Standard (non-constant-time) implementation for better performance

func (pf Proof) Verify(pkN, k *big.Int, ecdsaPub *crypto2.ECPoint) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// buffered to allow early exit
// uses cache primed in init()

// If prm divides N then Return 0

// is divisible

// ----- utils

func L(u, N *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

// GenerateXs generates the challenges used in Paillier key Proof
func GenerateXs(m int, k, N *big.Int, ecdsaPub *crypto2.ECPoint) []*big.Int {
	_ = "STUB: not implemented"
	return nil
}

// must be in order

// this should never happen. see: https://golang.org/pkg/hash/#Hash

// xi1||···||xib
