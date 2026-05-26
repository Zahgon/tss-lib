// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package test

import (
	"github.com/bnb-chain/tss-lib/v3/tss"
)

func SharedPartyUpdater(party tss.Party, msg tss.Message, errCh chan<- *tss.Error) {
	_ = "STUB: not implemented"
	// do not send a message from this party back to itself
	return
}
