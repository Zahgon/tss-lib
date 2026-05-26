// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"crypto/elliptic"
	"math/big"

	cmt "github.com/bnb-chain/tss-lib/v3/crypto/commitments"
	"github.com/bnb-chain/tss-lib/v3/crypto/schnorr"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

// These messages were generated from Protocol Buffers definitions into eddsa-signing.pb.go
// The following messages are registered on the Protocol Buffers "wire"

var (
	// Ensure that signing messages implement ValidateBasic
	_ = []tss.MessageContent{
		(*SignRound1Message)(nil),
		(*SignRound2Message)(nil),
		(*SignRound3Message)(nil),
	}
)

// ----- //

func NewSignRound1Message(
	from *tss.PartyID,
	commitment cmt.HashCommitment,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound1Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *SignRound1Message) UnmarshalCommitment() *big.Int { _ = "STUB: not implemented"; return nil }

// ----- //

func NewSignRound2Message(
	from *tss.PartyID,
	deCommitment cmt.HashDeCommitment,
	proof *schnorr.ZKProof,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound2Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *SignRound2Message) UnmarshalDeCommitment() []*big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (m *SignRound2Message) UnmarshalZKProof(ec elliptic.Curve) (*schnorr.ZKProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ----- //

func NewSignRound3Message(
	from *tss.PartyID,
	si *big.Int,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound3Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *SignRound3Message) UnmarshalS() *big.Int { _ = "STUB: not implemented"; return nil }
