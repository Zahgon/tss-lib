// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package keygen

import (
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/crypto"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

type (
	LocalSecrets struct {
		// secret fields (not shared, but stored locally)
		Xi, ShareID *big.Int // xi, kj
	}

	// Everything in LocalPartySaveData is saved locally to user's HD when done
	LocalPartySaveData struct {
		LocalSecrets

		// original indexes (ki in signing preparation phase)
		Ks []*big.Int

		// public keys (Xj = uj*G for each Pj)
		BigXj []*crypto.ECPoint // Xj

		// used for test assertions (may be discarded)
		EDDSAPub *crypto.ECPoint // y
	}
)

func NewLocalPartySaveData(partyCount int) (saveData LocalPartySaveData) {
	_ = "STUB: not implemented"
	return *new(LocalPartySaveData)
}

// BuildLocalSaveDataSubset re-creates the LocalPartySaveData to contain data for only the list of signing parties.
func BuildLocalSaveDataSubset(sourceData LocalPartySaveData, sortedIDs tss.SortedPartyIDs) LocalPartySaveData {
	_ = "STUB: not implemented"
	return *new(LocalPartySaveData)
}
