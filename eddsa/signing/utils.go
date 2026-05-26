// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"crypto/elliptic"
	"io"
	"math/big"

	"github.com/agl/ed25519/edwards25519"
)

func encodedBytesToBigInt(s *[32]byte) *big.Int {
	_ = "STUB: not implemented"
	// Use a copy so we don't screw up our original
	// memory.
	return nil
}

func bigIntToEncodedBytes(a *big.Int) *[32]byte { _ = "STUB: not implemented"; return nil }

// Caveat: a can be longer than 32 bytes.

// Reverse the byte string --> little endian after
// encoding.

func copyBytes(aB []byte) *[32]byte { _ = "STUB: not implemented"; return nil }

// If we have a short byte string, expand
// it so that it's long enough.

func ecPointToEncodedBytes(x *big.Int, y *big.Int) *[32]byte { _ = "STUB: not implemented"; return nil }

func reverse(s *[32]byte) { _ = "STUB: not implemented"; return }

func addExtendedElements(p, q edwards25519.ExtendedGroupElement) edwards25519.ExtendedGroupElement {
	_ = "STUB: not implemented"
	return *new(edwards25519.ExtendedGroupElement)
}

func ecPointToExtendedElement(ec elliptic.Curve, x *big.Int, y *big.Int, rand io.Reader) edwards25519.ExtendedGroupElement {
	_ = "STUB: not implemented"
	return *new(edwards25519.ExtendedGroupElement)
}
