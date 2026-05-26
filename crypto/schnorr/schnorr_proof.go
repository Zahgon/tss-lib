// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package schnorr

import (
	"io"
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/crypto"
)

type (
	ZKProof struct {
		Alpha *crypto.ECPoint
		T     *big.Int
	}

	ZKVProof struct {
		Alpha *crypto.ECPoint
		T, U  *big.Int
	}
)

// NewZKProof constructs a new Schnorr ZK proof of knowledge of the discrete logarithm (GG18Spec Fig. 16)
func NewZKProof(Session []byte, x *big.Int, X *crypto.ECPoint, rand io.Reader) (*ZKProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// already on the curve.

// SECURITY: Use constant-time multiplication for secret x

// NewZKProof verifies a new Schnorr ZK proof of knowledge of the discrete logarithm (GG18Spec Fig. 16)
func (pf *ZKProof) Verify(Session []byte, X *crypto.ECPoint) bool {
	_ = "STUB: not implemented"
	return false
}

func (pf *ZKProof) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

// NewZKProof constructs a new Schnorr ZK proof of knowledge s_i, l_i such that V_i = R^s_i, g^l_i (GG18Spec Fig. 17)
func NewZKVProof(Session []byte, V, R *crypto.ECPoint, s, l *big.Int, rand io.Reader) (*ZKVProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// already on the curve.

// SECURITY: Use constant-time multiplication for secret values s and l

func (pf *ZKVProof) Verify(Session []byte, V, R *crypto.ECPoint) bool {
	_ = "STUB: not implemented"
	return false
}

// already on the curve.

func (pf *ZKVProof) ValidateBasic() bool { _ = "STUB: not implemented"; return false }
