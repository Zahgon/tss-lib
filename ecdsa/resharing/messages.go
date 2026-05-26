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
	"github.com/bnb-chain/tss-lib/v3/crypto/dlnproof"
	"github.com/bnb-chain/tss-lib/v3/crypto/facproof"
	"github.com/bnb-chain/tss-lib/v3/crypto/modproof"
	"github.com/bnb-chain/tss-lib/v3/crypto/paillier"
	"github.com/bnb-chain/tss-lib/v3/crypto/vss"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

// These messages were generated from Protocol Buffers definitions into ecdsa-resharing.pb.go

var (
	// Ensure that signing messages implement ValidateBasic
	_ = []tss.MessageContent{
		(*DGRound1Message)(nil),
		(*DGRound2Message1)(nil),
		(*DGRound2Message2)(nil),
		(*DGRound3Message1)(nil),
		(*DGRound3Message2)(nil),
		(*DGRound4Message1)(nil),
		(*DGRound4Message2)(nil),
	}
)

// ----- //

func NewDGRound1Message(
	to []*tss.PartyID,
	from *tss.PartyID,
	ecdsaPub *crypto.ECPoint,
	vct cmt.HashCommitment,
	ssid []byte,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *DGRound1Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *DGRound1Message) UnmarshalECDSAPub(ec elliptic.Curve) (*crypto.ECPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *DGRound1Message) UnmarshalVCommitment() *big.Int { _ = "STUB: not implemented"; return nil }

func (m *DGRound1Message) UnmarshalSSID() []byte {
	_ = "STUB: not implemented"

	// ----- //
	return nil
}

func NewDGRound2Message1(
	to []*tss.PartyID,
	from *tss.PartyID,
	paillierPK *paillier.PublicKey,
	modProof *modproof.ProofMod,
	NTildei, H1i, H2i *big.Int,
	dlnProof1, dlnProof2 *dlnproof.Proof,
) (tss.ParsedMessage, error) {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage), nil
}

func (m *DGRound2Message1) ValidateBasic() bool {
	_ = "STUB: not implemented"

	// use with NoProofFac()
	// common.NonEmptyMultiBytes(m.ModProof, modproof.ProofModBytesParts) &&
	return false
}

// expected len of dln proof = sizeof(int64) + len(alpha) + len(t)

func (m *DGRound2Message1) UnmarshalPaillierPK() *paillier.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

func (m *DGRound2Message1) UnmarshalNTilde() *big.Int { _ = "STUB: not implemented"; return nil }

func (m *DGRound2Message1) UnmarshalH1() *big.Int { _ = "STUB: not implemented"; return nil }

func (m *DGRound2Message1) UnmarshalH2() *big.Int { _ = "STUB: not implemented"; return nil }

func (m *DGRound2Message1) UnmarshalModProof() (*modproof.ProofMod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *DGRound2Message1) UnmarshalDLNProof1() (*dlnproof.Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *DGRound2Message1) UnmarshalDLNProof2() (*dlnproof.Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ----- //

func NewDGRound2Message2(
	to []*tss.PartyID,
	from *tss.PartyID,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *DGRound2Message2) ValidateBasic() bool {
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

func NewDGRound4Message2(
	to []*tss.PartyID,
	from *tss.PartyID,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *DGRound4Message2) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func NewDGRound4Message1(
	to *tss.PartyID,
	from *tss.PartyID,
	proof *facproof.ProofFac,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *DGRound4Message1) ValidateBasic() bool {
	_ = "STUB: not implemented"

	// use with NoProofFac()
	// && common.NonEmptyMultiBytes(m.GetFacProof(), facproof.ProofFacBytesParts)
	return false
}

func (m *DGRound4Message1) UnmarshalFacProof() (*facproof.ProofFac, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
