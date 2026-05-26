// Copyright © Swingby

package ckd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"hash"
	"math/big"
)

type ExtendedKey struct {
	ecdsa.PublicKey
	Depth      uint8
	ChildIndex uint32
	ChainCode  []byte // 32 bytes
	ParentFP   []byte // parent fingerprint
	Version    []byte
}

// For more information about child key derivation see https://github.com/binance-chain/tss-lib/issues/104
// https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki .
// The functions below do not implement the full BIP-32 specification. As mentioned in the Jira ticket above,
// we only use non-hardened derived keys.

const (

	// HardenedKeyStart hardened key starts.
	HardenedKeyStart = 0x80000000 // 2^31

	// max Depth
	maxDepth = 1<<8 - 1

	PubKeyBytesLenCompressed = 33

	pubKeyCompressed byte = 0x2

	serializedKeyLen = 78

	// MinSeedBytes is the minimum number of bytes allowed for a seed to
	// a master node.
	MinSeedBytes = 16 // 128 bits

	// MaxSeedBytes is the maximum number of bytes allowed for a seed to
	// a master node.
	MaxSeedBytes = 64 // 512 bits
)

// Extended public key serialization, defined in BIP32
func (k *ExtendedKey) String() string {
	_ = "STUB: not implemented"
	// version(4) || depth(1) || parentFP (4) || childinde(4) || chaincode (32) || key(33) || checksum(4)
	return ""
}

// NewExtendedKeyFromString returns a new extended key from a base58-encoded extended key
func NewExtendedKeyFromString(key string, curve elliptic.Curve) (*ExtendedKey, error) {
	_ = "STUB: not implemented"
	// version(4) || depth(1) || parentFP (4) || childinde(4) || chaincode (32) || key(33) || checksum(4)
	return nil, nil
}

// Split the payload and checksum up and ensure the checksum matches.

// Deserialize each of the payload fields.

func doubleHashB(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func calcHash(buf []byte, hasher hash.Hash) []byte { _ = "STUB: not implemented"; return nil }

func hash160(buf []byte) []byte { _ = "STUB: not implemented"; return nil }

func isOdd(a *big.Int) bool { _ = "STUB: not implemented"; return false }

// PaddedAppend append src to dst, if less than size padding 0 at start
func paddedAppend(dst []byte, srcPaddedSize int, src []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// PaddedBytes padding byte array to size length
func paddedBytes(size int, src []byte) []byte { _ = "STUB: not implemented"; return nil }

// SerializeCompressed serializes a public key 33-byte compressed format
func serializeCompressed(publicKeyX *big.Int, publicKeyY *big.Int) []byte {
	_ = "STUB: not implemented"
	return nil
}

func DeriveChildKeyFromHierarchy(indicesHierarchy []uint32, pk *ExtendedKey, mod *big.Int, curve elliptic.Curve) (*big.Int, *ExtendedKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// DeriveChildKey Derive a child key from the given parent key. The function returns "IL" ("I left"), per BIP-32 spec. It also
// returns the derived child key.
func DeriveChildKey(index uint32, pk *ExtendedKey, curve elliptic.Curve) (*big.Int, *ExtendedKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// I = HMAC-SHA512(Key = chainCode, Data=data)

// falling outside of the valid range for curve private keys
