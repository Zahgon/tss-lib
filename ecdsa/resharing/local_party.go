// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package resharing

import (
	"fmt"
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/crypto"
	cmt "github.com/bnb-chain/tss-lib/v3/crypto/commitments"
	"github.com/bnb-chain/tss-lib/v3/crypto/vss"
	"github.com/bnb-chain/tss-lib/v3/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

// Implements Party
// Implements Stringer
var _ tss.Party = (*LocalParty)(nil)
var _ fmt.Stringer = (*LocalParty)(nil)

type (
	LocalParty struct {
		*tss.BaseParty
		params *tss.ReSharingParameters

		temp        localTempData
		input, save keygen.LocalPartySaveData

		// outbound messaging
		out chan<- tss.Message
		end chan<- *keygen.LocalPartySaveData
	}

	localMessageStore struct {
		dgRound1Messages,
		dgRound2Message1s,
		dgRound2Message2s,
		dgRound3Message1s,
		dgRound3Message2s,
		dgRound4Message1s,
		dgRound4Message2s []tss.ParsedMessage
	}

	localTempData struct {
		localMessageStore

		// temp data (thrown away after rounds)
		NewVs     vss.Vs
		NewShares vss.Shares
		VD        cmt.HashDeCommitment

		// temporary storage of data that is persisted by the new party in round 5 if all "ACK" messages are received
		newXi     *big.Int
		newKs     []*big.Int
		newBigXjs []*crypto.ECPoint // Xj to save in round 5

		ssid      []byte
		ssidNonce *big.Int
	}
)

// Exported, used in `tss` client
// The `key` is read from and/or written to depending on whether this party is part of the old or the new committee.
// You may optionally generate and set the LocalPreParams if you would like to use pre-generated safe primes and Paillier secret.
// (This is similar to providing the `optionalPreParams` to `keygen.LocalParty`).
func NewLocalParty(
	params *tss.ReSharingParameters,
	key keygen.LocalPartySaveData,
	out chan<- tss.Message,
	end chan<- *keygen.LocalPartySaveData,
) tss.Party {
	_ = "STUB: not implemented"
	return *new(tss.Party)
}

// msgs init
// from t+1 of Old Committee
// from n of New Committee
// "
// from t+1 of Old Committee
// "
// from n of New Committee
// from n of New Committee
// save data init

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
