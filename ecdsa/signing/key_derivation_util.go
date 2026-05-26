// Copyright © 2021 Swingby

package signing

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/crypto"
	"github.com/bnb-chain/tss-lib/v3/crypto/ckd"
	"github.com/bnb-chain/tss-lib/v3/ecdsa/keygen"
)

func UpdatePublicKeyAndAdjustBigXj(keyDerivationDelta *big.Int, keys []keygen.LocalPartySaveData, extendedChildPk *ecdsa.PublicKey, ec elliptic.Curve) error {
	_ = "STUB: not implemented"
	return nil
}

// Suppose X_j has shamir shares X_j0,     X_j1,     ..., X_jn
// So X_j + D has shamir shares  X_j0 + D, X_j1 + D, ..., X_jn + D

func derivingPubkeyFromPath(masterPub *crypto.ECPoint, chainCode []byte, path []uint32, ec elliptic.Curve) (*big.Int, *ckd.ExtendedKey, error) {
	_ = "STUB: not implemented"
	// build ecdsa key pair
	return nil, nil, nil
}
