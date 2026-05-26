// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package keygen

import (
	"context"
	"io"
	"time"
)

const (
	// Using a modulus length of 2048 is recommended in the GG18 spec
	paillierModulusLen = 2048
	// Two 1024-bit safe primes to produce NTilde
	safePrimeBitLen = 1024
	// Ticker for printing log statements while generating primes/modulus
	logProgressTickInterval = 8 * time.Second
	// Safe big len using random for ssid
	SafeBitLen = 1024
)

// GeneratePreParams finds two safe primes and computes the Paillier secret required for the protocol.
// This can be a time consuming process so it is recommended to do it out-of-band.
// If not specified, a concurrency value equal to the number of available CPU cores will be used.
// If pre-parameters could not be generated before the timeout, an error is returned.
func GeneratePreParams(timeout time.Duration, optionalConcurrency ...int) (*LocalPreParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GeneratePreParams finds two safe primes and computes the Paillier secret required for the protocol.
// This can be a time consuming process so it is recommended to do it out-of-band.
// If not specified, a concurrency value equal to the number of available CPU cores will be used.
// If pre-parameters could not be generated before the context is done, an error is returned.
func GeneratePreParamsWithContext(ctx context.Context, optionalConcurrency ...int) (*LocalPreParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GeneratePreParams finds two safe primes and computes the Paillier secret required for the protocol.
// This can be a time consuming process so it is recommended to do it out-of-band.
// If not specified, a concurrency value equal to the number of available CPU cores will be used.
// If pre-parameters could not be generated before the context is done, an error is returned.
func GeneratePreParamsWithContextAndRandom(ctx context.Context, rand io.Reader, optionalConcurrency ...int) (*LocalPreParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// prepare for concurrent Paillier and safe prime generation

// 4. generate Paillier public key E_i, private key and proof

// more concurrency weight is assigned here because the paillier primes have a requirement of having "large" P-Q

// 5-7. generate safe primes for ZKPs used later on

// this ticker will print a log statement while the generating is still in progress

// errors can be thrown in the following code; consume chans to end goroutines here
