// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package keygen

import (
	"crypto/elliptic"
	"math/big"

	cmt "github.com/bnb-chain/tss-lib/v3/crypto/commitments"
	"github.com/bnb-chain/tss-lib/v3/crypto/schnorr"
	"github.com/bnb-chain/tss-lib/v3/crypto/vss"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

// These messages were generated from Protocol Buffers definitions into eddsa-keygen.pb.go
// The following messages are registered on the Protocol Buffers "wire"

var (
	// Ensure that keygen messages implement ValidateBasic
	_ = []tss.MessageContent{
		(*KGRound1Message)(nil),
		(*KGRound2Message1)(nil),
		(*KGRound2Message2)(nil),
	}
)

// ----- //

func NewKGRound1Message(from *tss.PartyID, ct cmt.HashCommitment) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *KGRound1Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *KGRound1Message) UnmarshalCommitment() *big.Int { _ = "STUB: not implemented"; return nil }

// ----- //

func NewKGRound2Message1(
	to, from *tss.PartyID,
	share *vss.Share,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *KGRound2Message1) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *KGRound2Message1) UnmarshalShare() *big.Int { _ = "STUB: not implemented"; return nil }

// ----- //

func NewKGRound2Message2(
	from *tss.PartyID,
	deCommitment cmt.HashDeCommitment,
	proof *schnorr.ZKProof,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *KGRound2Message2) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *KGRound2Message2) UnmarshalDeCommitment() []*big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (m *KGRound2Message2) UnmarshalZKProof(ec elliptic.Curve) (*schnorr.ZKProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
