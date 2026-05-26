// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"crypto/elliptic"
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/crypto"
)

// PrepareForSigning(), GG18Spec (11) Fig. 14
func PrepareForSigning(ec elliptic.Curve, i, pax int, xi *big.Int, ks []*big.Int, bigXs []*crypto.ECPoint) (wi *big.Int, bigWs []*crypto.ECPoint) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 2-4.

// big.Int Div is calculated as: a/b = a * modInv(b,q)

// 5-10.

// big.Int Div is calculated as: a/b = a * modInv(b,q)
