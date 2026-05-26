// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package common

import (
	"math/big"
)

func BigIntsToBytes(bigInts []*big.Int) [][]byte { _ = "STUB: not implemented"; return nil }

func MultiBytesToBigInts(bytes [][]byte) []*big.Int { _ = "STUB: not implemented"; return nil }

// Returns true when the byte slice is non-nil and non-empty
func NonEmptyBytes(bz []byte) bool { _ = "STUB: not implemented"; return false }

// Returns true when all of the slices in the multi-dimensional byte slice are non-nil and non-empty
func NonEmptyMultiBytes(bzs [][]byte, expectLen ...int) bool {
	_ = "STUB: not implemented"
	return false
}

// variadic (optional) arg test

// PadToLengthBytesInPlace pad {0, ...} to the front of src if len(src) < length
// output length is equal to the parameter length
func PadToLengthBytesInPlace(src []byte, length int) []byte { _ = "STUB: not implemented"; return nil }
