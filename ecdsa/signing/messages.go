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
	"github.com/bnb-chain/tss-lib/v3/crypto/mta"
	"github.com/bnb-chain/tss-lib/v3/crypto/schnorr"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

// These messages were generated from Protocol Buffers definitions into ecdsa-signing.pb.go
// The following messages are registered on the Protocol Buffers "wire"

var (
	// Ensure that signing messages implement ValidateBasic
	_ = []tss.MessageContent{
		(*SignRound1Message1)(nil),
		(*SignRound1Message2)(nil),
		(*SignRound2Message)(nil),
		(*SignRound3Message)(nil),
		(*SignRound4Message)(nil),
		(*SignRound5Message)(nil),
		(*SignRound6Message)(nil),
		(*SignRound7Message)(nil),
		(*SignRound8Message)(nil),
		(*SignRound9Message)(nil),
	}
)

// ----- //

func NewSignRound1Message1(
	to, from *tss.PartyID,
	c *big.Int,
	proof *mta.RangeProofAlice,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound1Message1) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *SignRound1Message1) UnmarshalC() *big.Int { _ = "STUB: not implemented"; return nil }

func (m *SignRound1Message1) UnmarshalRangeProofAlice() (*mta.RangeProofAlice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ----- //

func NewSignRound1Message2(
	from *tss.PartyID,
	commitment cmt.HashCommitment,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound1Message2) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *SignRound1Message2) UnmarshalCommitment() *big.Int { _ = "STUB: not implemented"; return nil }

// ----- //

func NewSignRound2Message(
	to, from *tss.PartyID,
	c1Ji *big.Int,
	pi1Ji *mta.ProofBob,
	c2Ji *big.Int,
	pi2Ji *mta.ProofBobWC,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound2Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *SignRound2Message) UnmarshalProofBob() (*mta.ProofBob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *SignRound2Message) UnmarshalProofBobWC(ec elliptic.Curve) (*mta.ProofBobWC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ----- //

func NewSignRound3Message(
	from *tss.PartyID,
	theta *big.Int,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound3Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

// ----- //

func NewSignRound4Message(
	from *tss.PartyID,
	deCommitment cmt.HashDeCommitment,
	proof *schnorr.ZKProof,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound4Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *SignRound4Message) UnmarshalDeCommitment() []*big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (m *SignRound4Message) UnmarshalZKProof(ec elliptic.Curve) (*schnorr.ZKProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ----- //

func NewSignRound5Message(
	from *tss.PartyID,
	commitment cmt.HashCommitment,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound5Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *SignRound5Message) UnmarshalCommitment() *big.Int { _ = "STUB: not implemented"; return nil }

// ----- //

func NewSignRound6Message(
	from *tss.PartyID,
	deCommitment cmt.HashDeCommitment,
	proof *schnorr.ZKProof,
	vProof *schnorr.ZKVProof,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound6Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *SignRound6Message) UnmarshalDeCommitment() []*big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (m *SignRound6Message) UnmarshalZKProof(ec elliptic.Curve) (*schnorr.ZKProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *SignRound6Message) UnmarshalZKVProof(ec elliptic.Curve) (*schnorr.ZKVProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ----- //

func NewSignRound7Message(
	from *tss.PartyID,
	commitment cmt.HashCommitment,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound7Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *SignRound7Message) UnmarshalCommitment() *big.Int { _ = "STUB: not implemented"; return nil }

// ----- //

func NewSignRound8Message(
	from *tss.PartyID,
	deCommitment cmt.HashDeCommitment,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound8Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *SignRound8Message) UnmarshalDeCommitment() []*big.Int {
	_ = "STUB: not implemented"
	return nil
}

// ----- //

func NewSignRound9Message(
	from *tss.PartyID,
	si *big.Int,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *SignRound9Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *SignRound9Message) UnmarshalS() *big.Int { _ = "STUB: not implemented"; return nil }
