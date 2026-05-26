// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package keygen

import (
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/crypto/dlnproof"
)

type DlnProofVerifier struct {
	semaphore chan interface{}
}

type message interface {
	UnmarshalDLNProof1() (*dlnproof.Proof, error)
	UnmarshalDLNProof2() (*dlnproof.Proof, error)
}

func NewDlnProofVerifier(concurrency int) *DlnProofVerifier { _ = "STUB: not implemented"; return nil }

func (dpv *DlnProofVerifier) VerifyDLNProof1(
	Session []byte,
	m message,
	h1, h2, n *big.Int,
	onDone func(bool),
) {
	_ = "STUB: not implemented"
	return
}

func (dpv *DlnProofVerifier) VerifyDLNProof2(
	Session []byte,
	m message,
	h1, h2, n *big.Int,
	onDone func(bool),
) {
	_ = "STUB: not implemented"
	return
}
