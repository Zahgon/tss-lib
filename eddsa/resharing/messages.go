// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package resharing

import (
	"crypto/elliptic"
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/crypto"
	cmt "github.com/bnb-chain/tss-lib/v3/crypto/commitments"
	"github.com/bnb-chain/tss-lib/v3/crypto/vss"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

// These messages were generated from Protocol Buffers definitions into eddsa-resharing.pb.go

var (
	// Ensure that signing messages implement ValidateBasic
	_ = []tss.MessageContent{
		(*DGRound1Message)(nil),
		(*DGRound2Message)(nil),
		(*DGRound3Message1)(nil),
		(*DGRound3Message2)(nil),
		(*DGRound4Message)(nil),
	}
)

// ----- //

func NewDGRound1Message(
	to []*tss.PartyID,
	from *tss.PartyID,
	eddsaPub *crypto.ECPoint,
	vct cmt.HashCommitment,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *DGRound1Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *DGRound1Message) UnmarshalEDDSAPub(ec elliptic.Curve) (*crypto.ECPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *DGRound1Message) UnmarshalVCommitment() *big.Int { _ = "STUB: not implemented"; return nil }

// ----- //

func NewDGRound2Message(
	to []*tss.PartyID,
	from *tss.PartyID,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *DGRound2Message) ValidateBasic() bool {
	_ = "STUB: not implemented"

	// ----- //
	return false
}

func NewDGRound3Message1(
	to *tss.PartyID,
	from *tss.PartyID,
	share *vss.Share,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *DGRound3Message1) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

// ----- //

func NewDGRound3Message2(
	to []*tss.PartyID,
	from *tss.PartyID,
	vdct cmt.HashDeCommitment,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *DGRound3Message2) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *DGRound3Message2) UnmarshalVDeCommitment() cmt.HashDeCommitment {
	_ = "STUB: not implemented"
	return *new(cmt.HashDeCommitment)
}

// ----- //

func NewDGRound4Message(
	to []*tss.PartyID,
	from *tss.PartyID,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *DGRound4Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }
