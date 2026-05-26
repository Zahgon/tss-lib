// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"github.com/bnb-chain/tss-lib/v3/tss"
)

func (round *finalization) Start() *tss.Error { _ = "STUB: not implemented"; return nil }

// byte v = if(R.X > curve.N) then 2 else 0) | (if R.Y.IsEven then 0 else 1);

// This is copied from:
// https://github.com/btcsuite/btcd/blob/c26ffa870fd817666a857af1bf6498fabba1ffe3/btcec/signature.go#L442-L444
// This is needed because of tendermint checks here:
// https://github.com/tendermint/tendermint/blob/d9481e3648450cb99e15c6a070c1fb69aa0c255b/crypto/secp256k1/secp256k1_nocgo.go#L43-L47

// save the signature for final output

func (round *finalization) CanAccept(msg tss.ParsedMessage) bool {
	_ = "STUB: not implemented"
	// not expecting any incoming messages in this round
	return false
}

func (round *finalization) Update() (bool, *tss.Error) {
	_ = "STUB: not implemented"
	// not expecting any incoming messages in this round
	return false, nil
}

func (round *finalization) NextRound() tss.Round {
	_ = "STUB: not implemented"
	// finished!
	return *new(tss.Round)
}

func padToLengthBytesInPlace(src []byte, length int) []byte { _ = "STUB: not implemented"; return nil }
