// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package common

import (
	"math/big"
)

// modInt is a *big.Int that performs all of its arithmetic with modular reduction.
type modInt big.Int

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
	two  = big.NewInt(2)
)

func ModInt(mod *big.Int) *modInt { _ = "STUB: not implemented"; return nil }

func (mi *modInt) Add(x, y *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (mi *modInt) Sub(x, y *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (mi *modInt) Div(x, y *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (mi *modInt) Mul(x, y *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (mi *modInt) Exp(x, y *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (mi *modInt) ModInverse(g *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (mi *modInt) i() *big.Int { _ = "STUB: not implemented"; return nil }

func IsInInterval(b *big.Int, bound *big.Int) bool { _ = "STUB: not implemented"; return false }

func AppendBigIntToBytesSlice(commonBytes []byte, appended *big.Int) []byte {
	_ = "STUB: not implemented"
	return nil
}
