// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package tss

// Used externally to update a LocalParty with a valid ParsedMessage
func ParseWireMessage(wireBytes []byte, from *PartyID, isBroadcast bool) (ParsedMessage, error) {
	_ = "STUB: not implemented"
	return *new(ParsedMessage), nil
}

func parseWrappedMessage(wire *MessageWrapper, from *PartyID) (ParsedMessage, error) {
	_ = "STUB: not implemented"
	return *new(ParsedMessage), nil
}
