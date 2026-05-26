// Copyright © 2019-2020 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package commitments

import (
	"math/big"
)

const (
	PartsCap    = 3
	MaxPartSize = int64(1 * 1024 * 1024) // 1 MB - rather liberal
)

type builder struct {
	parts [][]*big.Int
}

func NewBuilder() *builder { _ = "STUB: not implemented"; return nil }

func (b *builder) Parts() [][]*big.Int { _ = "STUB: not implemented"; return nil }

func (b *builder) AddPart(part []*big.Int) *builder { _ = "STUB: not implemented"; return nil }

func (b *builder) Secrets() ([]*big.Int, error) { _ = "STUB: not implemented"; return nil, nil }

// +1 to accommodate length prefix element

func ParseSecrets(secrets []*big.Int) ([][]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// are we looking at a length prefix element? (first one is)
