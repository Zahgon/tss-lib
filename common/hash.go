// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package common

import (
	_ "crypto/sha512"
	"math/big"
)

const (
	hashInputDelimiter = byte('$')
)

// SHA-512/256 is protected against length extension attacks and is more performant than SHA-256 on 64-bit architectures.
// https://en.wikipedia.org/wiki/Template:Comparison_of_SHA_functions
func SHA512_256(in ...[]byte) []byte { _ = "STUB: not implemented"; return nil }

// prevent hash collisions with this prefix containing the block count

// converting between int and uint64 doesn't change the sign bit, but it may be interpreted as a larger value.
// this prefix is never read/interpreted, so that doesn't matter.

// safety delimiter
// 64-bits

// Security audit: length of each byte buffer should be added after
// each security delimiters in order to enforce proper domain separation

// n < len(data) or an error will never happen.
// see: https://golang.org/pkg/hash/#Hash and https://github.com/golang/go/wiki/Hashing#the-hashhash-interface

func SHA512_256i(in ...*big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

// prevent hash collisions with this prefix containing the block count

// converting between int and uint64 doesn't change the sign bit, but it may be interpreted as a larger value.
// this prefix is never read/interpreted, so that doesn't matter.

// safety delimiter
// 64-bits

// Security audit: length of each byte buffer should be added after
// each security delimiters in order to enforce proper domain separation

// n < len(data) or an error will never happen.
// see: https://golang.org/pkg/hash/#Hash and https://github.com/golang/go/wiki/Hashing#the-hashhash-interface

// SHA512_256i_TAGGED tagged version of SHA512_256i
func SHA512_256i_TAGGED(tag []byte, in ...*big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

// prevent hash collisions with this prefix containing the block count

// converting between int and uint64 doesn't change the sign bit, but it may be interpreted as a larger value.
// this prefix is never read/interpreted, so that doesn't matter.

// safety delimiter
// 64-bits

// Security audit: length of each byte buffer should be added after
// each security delimiters in order to enforce proper domain separation

// n < len(data) or an error will never happen.
// see: https://golang.org/pkg/hash/#Hash and https://github.com/golang/go/wiki/Hashing#the-hashhash-interface

func SHA512_256iOne(in *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

// n < len(data) or an error will never happen.
// see: https://golang.org/pkg/hash/#Hash and https://github.com/golang/go/wiki/Hashing#the-hashhash-interface
