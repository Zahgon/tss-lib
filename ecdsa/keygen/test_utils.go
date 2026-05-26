// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package keygen

import (
	"math/big"

	"github.com/bnb-chain/tss-lib/v3/test"
	"github.com/bnb-chain/tss-lib/v3/tss"
)

const (
	// To change these parameters, you must first delete the text fixture files in test/_fixtures/ and then run the keygen test alone.
	// Then the signing and resharing tests will work with the new n, t configuration using the newly written fixture files.
	TestParticipants = test.TestParticipants
	TestThreshold    = test.TestParticipants / 2
)
const (
	testFixtureDirFormat  = "%s/../../test/_ecdsa_fixtures"
	testFixtureFileFormat = "keygen_data_%d.json"
)

func LoadKeygenTestFixtures(qty int, optionalStart ...int) ([]LocalPartySaveData, tss.SortedPartyIDs, error) {
	_ = "STUB: not implemented"
	return nil, *new(tss.SortedPartyIDs), nil
}

func LoadKeygenTestFixturesRandomSet(qty, fixtureCount int) ([]LocalPartySaveData, tss.SortedPartyIDs, error) {
	_ = "STUB: not implemented"
	return nil, *new(tss.SortedPartyIDs), nil
}

func LoadNTildeH1H2FromTestFixture(idx int) (NTildei, h1i, h2i *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func makeTestFixtureFilePath(partyIndex int) string { _ = "STUB: not implemented"; return "" }
