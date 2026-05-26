// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"crypto/elliptic"
	"math/big"
)

// PrepareForSigning(), Fig. 7
func PrepareForSigning(ec elliptic.Curve, i, pax int, xi *big.Int, ks []*big.Int) (wi *big.Int) {
	_ = "STUB: not implemented"
	return nil
}

// 1-4.

// big.Int Div is calculated as: a/b = a * modInv(b,q)
