// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"fmt"
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/common"
	"github.com/bnb-chain/tss-lib/v3/crypto"
	cmt "github.com/bnb-chain/tss-lib/v3/crypto/commitments"
	"github.com/bnb-chain/tss-lib/v3/eddsa/keygen"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

// Implements Party
// Implements Stringer
var _ tss.Party = (*LocalParty)(nil)
var _ fmt.Stringer = (*LocalParty)(nil)

type (
	LocalParty struct {
		*tss.BaseParty
		params *tss.Parameters

		keys keygen.LocalPartySaveData
		temp localTempData
		data *common.SignatureData

		// outbound messaging
		out chan<- tss.Message
		end chan<- *common.SignatureData
	}

	localMessageStore struct {
		signRound1Messages,
		signRound2Messages,
		signRound3Messages []tss.ParsedMessage
	}

	localTempData struct {
		localMessageStore

		// temp data (thrown away after sign) / round 1
		wi,
		m,
		ri *big.Int
		fullBytesLen int
		pointRi      *crypto.ECPoint
		deCommit     cmt.HashDeCommitment

		// round 2
		cjs []*big.Int
		si  *[32]byte

		// round 3
		r *big.Int

		ssid      []byte
		ssidNonce *big.Int
	}
)

func NewLocalParty(
	msg *big.Int,
	params *tss.Parameters,
	key keygen.LocalPartySaveData,
	out chan<- tss.Message,
	end chan<- *common.SignatureData,
	fullBytesLen ...int,
) tss.Party {
	_ = "STUB: not implemented"
	return *new(tss.Party)
}

// msgs init

// temp data init

func (p *LocalParty) FirstRound() tss.Round { _ = "STUB: not implemented"; return *new(tss.Round) }

func (p *LocalParty) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

func (p *LocalParty) Update(msg tss.ParsedMessage) (ok bool, err *tss.Error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *LocalParty) UpdateFromBytes(wireBytes []byte, from *tss.PartyID, isBroadcast bool) (bool, *tss.Error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *LocalParty) ValidateMessage(msg tss.ParsedMessage) (bool, *tss.Error) {
	_ = "STUB: not implemented"
	return false, nil
}

// check that the message's "from index" will fit into the array

func (p *LocalParty) StoreMessage(msg tss.ParsedMessage) (bool, *tss.Error) {
	_ = "STUB: not implemented"
	// ValidateBasic is cheap; double-check the message here in case the public StoreMessage was called externally
	return false, nil
}

// switch/case is necessary to store any messages beyond current round
// this does not handle message replays. we expect the caller to apply replay and spoofing protection.

// unrecognised message, just ignore!

func (p *LocalParty) PartyID() *tss.PartyID { _ = "STUB: not implemented"; return nil }

func (p *LocalParty) String() string { _ = "STUB: not implemented"; return "" }
