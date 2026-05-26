// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package keygen

import (
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/crypto/facproof"
	"github.com/bnb-chain/tss-lib/v3/crypto/modproof"

	cmt "github.com/bnb-chain/tss-lib/v3/crypto/commitments"
	"github.com/bnb-chain/tss-lib/v3/crypto/dlnproof"
	"github.com/bnb-chain/tss-lib/v3/crypto/paillier"
	"github.com/bnb-chain/tss-lib/v3/crypto/vss"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

// These messages were generated from Protocol Buffers definitions into ecdsa-keygen.pb.go
// The following messages are registered on the Protocol Buffers "wire"

var (
	// Ensure that keygen messages implement ValidateBasic
	_ = []tss.MessageContent{
		(*KGRound1Message)(nil),
		(*KGRound2Message1)(nil),
		(*KGRound2Message2)(nil),
		(*KGRound3Message)(nil),
	}
)

// ----- //

func NewKGRound1Message(
	from *tss.PartyID,
	ct cmt.HashCommitment,
	paillierPK *paillier.PublicKey,
	nTildeI, h1I, h2I *big.Int,
	dlnProof1, dlnProof2 *dlnproof.Proof,
) (tss.ParsedMessage, error) {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage), nil
}

func (m *KGRound1Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

// expected len of dln proof = sizeof(int64) + len(alpha) + len(t)

func (m *KGRound1Message) UnmarshalCommitment() *big.Int { _ = "STUB: not implemented"; return nil }

func (m *KGRound1Message) UnmarshalPaillierPK() *paillier.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

func (m *KGRound1Message) UnmarshalNTilde() *big.Int { _ = "STUB: not implemented"; return nil }

func (m *KGRound1Message) UnmarshalH1() *big.Int { _ = "STUB: not implemented"; return nil }

func (m *KGRound1Message) UnmarshalH2() *big.Int { _ = "STUB: not implemented"; return nil }

func (m *KGRound1Message) UnmarshalDLNProof1() (*dlnproof.Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *KGRound1Message) UnmarshalDLNProof2() (*dlnproof.Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ----- //

func NewKGRound2Message1(
	to, from *tss.PartyID,
	share *vss.Share,
	proof *facproof.ProofFac,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *KGRound2Message1) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

// This is commented for backward compatibility, which msg has no proof
// && common.NonEmptyMultiBytes(m.GetFacProof(), facproof.ProofFacBytesParts)

func (m *KGRound2Message1) UnmarshalShare() *big.Int { _ = "STUB: not implemented"; return nil }

func (m *KGRound2Message1) UnmarshalFacProof() (*facproof.ProofFac, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ----- //

func NewKGRound2Message2(
	from *tss.PartyID,
	deCommitment cmt.HashDeCommitment,
	proof *modproof.ProofMod,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *KGRound2Message2) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

// This is commented for backward compatibility, which msg has no proof
// && common.NonEmptyMultiBytes(m.GetModProof(), modproof.ProofModBytesParts)

func (m *KGRound2Message2) UnmarshalDeCommitment() []*big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (m *KGRound2Message2) UnmarshalModProof() (*modproof.ProofMod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ----- //

func NewKGRound3Message(
	from *tss.PartyID,
	proof paillier.Proof,
) tss.ParsedMessage {
	_ = "STUB: not implemented"
	return *new(tss.ParsedMessage)
}

func (m *KGRound3Message) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (m *KGRound3Message) UnmarshalProofInts() paillier.Proof {
	_ = "STUB: not implemented"
	return *new(paillier.Proof)
}
