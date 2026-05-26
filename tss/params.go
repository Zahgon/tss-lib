// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package tss

import (
	"crypto/elliptic"
	"io"
	"math/big"
	"time"
)

type (
	Parameters struct {
		ec                  elliptic.Curve
		partyID             *PartyID
		parties             *PeerContext
		partyCount          int
		threshold           int
		concurrency         int
		safePrimeGenTimeout time.Duration
		// sessionNonce provides per-session SSID uniqueness for GG20 session binding.
		// For signing, defaults to the message hash if not set.
		// For keygen/resharing, the caller SHOULD set this to a value agreed upon by
		// all parties (e.g., a coordinator-assigned session ID) to prevent cross-session
		// proof replay. If not set, falls back to 0 (no session binding).
		sessionNonce *big.Int
		// for keygen
		noProofMod bool
		noProofFac bool
		// random sources
		partialKeyRand, rand io.Reader
	}

	ReSharingParameters struct {
		*Parameters
		newParties    *PeerContext
		newPartyCount int
		newThreshold  int
	}
)

const (
	defaultSafePrimeGenTimeout = 5 * time.Minute
)

// Exported, used in `tss` client
func NewParameters(ec elliptic.Curve, ctx *PeerContext, partyID *PartyID, partyCount, threshold int) *Parameters {
	_ = "STUB: not implemented"
	return nil
}

func (params *Parameters) EC() elliptic.Curve {
	_ = "STUB: not implemented"
	return *new(elliptic.Curve)
}

func (params *Parameters) Parties() *PeerContext { _ = "STUB: not implemented"; return nil }

func (params *Parameters) PartyID() *PartyID { _ = "STUB: not implemented"; return nil }

func (params *Parameters) PartyCount() int { _ = "STUB: not implemented"; return 0 }

func (params *Parameters) Threshold() int { _ = "STUB: not implemented"; return 0 }

func (params *Parameters) Concurrency() int { _ = "STUB: not implemented"; return 0 }

func (params *Parameters) SafePrimeGenTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// The concurrency level must be >= 1.
func (params *Parameters) SetConcurrency(concurrency int) { _ = "STUB: not implemented"; return }

func (params *Parameters) SetSafePrimeGenTimeout(timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (params *Parameters) NoProofMod() bool { _ = "STUB: not implemented"; return false }

func (params *Parameters) NoProofFac() bool { _ = "STUB: not implemented"; return false }

func (params *Parameters) SetNoProofMod() { _ = "STUB: not implemented"; return }

func (params *Parameters) SetNoProofFac() { _ = "STUB: not implemented"; return }

func (params *Parameters) PartialKeyRand() io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func (params *Parameters) Rand() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (params *Parameters) SetPartialKeyRand(rand io.Reader) { _ = "STUB: not implemented"; return }

func (params *Parameters) SetRand(rand io.Reader) {
	_ = "STUB: not implemented"

	// SessionNonce returns the per-session nonce for SSID uniqueness.
	// Returns nil if not set.
	return
}

func (params *Parameters) SessionNonce() *big.Int { _ = "STUB: not implemented"; return nil }

// SetSessionNonce sets a per-session nonce that all parties must agree on.
// This value is mixed into the SSID to provide GG20 session binding, preventing
// cross-session proof replay attacks. All parties in the same session MUST use
// the same nonce value. The caller is responsible for coordinating this.
func (params *Parameters) SetSessionNonce(nonce *big.Int) { _ = "STUB: not implemented"; return }

// ----- //

// Exported, used in `tss` client
func NewReSharingParameters(ec elliptic.Curve, ctx, newCtx *PeerContext, partyID *PartyID, partyCount, threshold, newPartyCount, newThreshold int) *ReSharingParameters {
	_ = "STUB: not implemented"
	return nil
}

func (rgParams *ReSharingParameters) OldParties() *PeerContext {
	_ = "STUB: not implemented"
	return nil
	// wr use the original method for old parties
}

func (rgParams *ReSharingParameters) OldPartyCount() int { _ = "STUB: not implemented"; return 0 }

func (rgParams *ReSharingParameters) NewParties() *PeerContext {
	_ = "STUB: not implemented"
	return nil
}

func (rgParams *ReSharingParameters) NewPartyCount() int { _ = "STUB: not implemented"; return 0 }

func (rgParams *ReSharingParameters) NewThreshold() int { _ = "STUB: not implemented"; return 0 }

func (rgParams *ReSharingParameters) OldAndNewParties() []*PartyID {
	_ = "STUB: not implemented"
	return nil
}

func (rgParams *ReSharingParameters) OldAndNewPartyCount() int { _ = "STUB: not implemented"; return 0 }

func (rgParams *ReSharingParameters) IsOldCommittee() bool { _ = "STUB: not implemented"; return false }

func (rgParams *ReSharingParameters) IsNewCommittee() bool { _ = "STUB: not implemented"; return false }
