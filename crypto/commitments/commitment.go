// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

// partly ported from:
// https://github.com/KZen-networks/curv/blob/78a70f43f5eda376e5888ce33aec18962f572bbe/src/cryptographic_primitives/commitments/hash_commitment.rs

package commitments

import (
	"io"
	"math/big"
)

const (
	HashLength = 256
)

type (
	HashCommitment   = *big.Int
	HashDeCommitment = []*big.Int

	HashCommitDecommit struct {
		C HashCommitment
		D HashDeCommitment
	}
)

func NewHashCommitmentWithRandomness(r *big.Int, secrets ...*big.Int) *HashCommitDecommit {
	_ = "STUB: not implemented"
	return nil
}

func NewHashCommitment(rand io.Reader, secrets ...*big.Int) *HashCommitDecommit {
	_ = "STUB: not implemented"
	return nil
}

// r

func NewHashDeCommitmentFromBytes(marshalled [][]byte) HashDeCommitment {
	_ = "STUB: not implemented"
	return *new(HashDeCommitment)
}

func (cmt *HashCommitDecommit) Verify() bool { _ = "STUB: not implemented"; return false }

func (cmt *HashCommitDecommit) DeCommit() (bool, HashDeCommitment) {
	_ = "STUB: not implemented"

	// [1:] skips random element r in D
	return false, *new(HashDeCommitment)
}
