// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package tss

import (
	"sync"
)

type Party interface {
	Start() *Error
	// The main entry point when updating a party's state from the wire.
	// isBroadcast should represent whether the message was received via a reliable broadcast
	UpdateFromBytes(wireBytes []byte, from *PartyID, isBroadcast bool) (ok bool, err *Error)
	// You may use this entry point to update a party's state when running locally or in tests
	Update(msg ParsedMessage) (ok bool, err *Error)
	Running() bool
	WaitingFor() []*PartyID
	ValidateMessage(msg ParsedMessage) (bool, *Error)
	StoreMessage(msg ParsedMessage) (bool, *Error)
	FirstRound() Round
	WrapError(err error, culprits ...*PartyID) *Error
	PartyID() *PartyID
	String() string

	// Private lifecycle methods
	setRound(Round) *Error
	round() Round
	advance()
	lock()
	unlock()
}

type BaseParty struct {
	mtx        sync.Mutex
	rnd        Round
	FirstRound Round
}

func (p *BaseParty) Running() bool { _ = "STUB: not implemented"; return false }

func (p *BaseParty) WaitingFor() []*PartyID { _ = "STUB: not implemented"; return nil }

func (p *BaseParty) WrapError(err error, culprits ...*PartyID) *Error {
	_ = "STUB: not implemented"
	return nil
}

// an implementation of ValidateMessage that is shared across the different types of parties (keygen, signing, dynamic groups)
func (p *BaseParty) ValidateMessage(msg ParsedMessage) (bool, *Error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *BaseParty) String() string { _ = "STUB: not implemented"; return "" }

// -----
// Private lifecycle methods

func (p *BaseParty) setRound(round Round) *Error { _ = "STUB: not implemented"; return nil }

func (p *BaseParty) round() Round { _ = "STUB: not implemented"; return *new(Round) }

func (p *BaseParty) advance() { _ = "STUB: not implemented"; return }

func (p *BaseParty) lock() { _ = "STUB: not implemented"; return }

func (p *BaseParty) unlock() {
	_ = "STUB: not implemented"

	// ----- //
	return
}

func BaseStart(p Party, task string, prepare ...func(Round) *Error) *Error {
	_ = "STUB: not implemented"
	return nil
}

// an implementation of Update that is shared across the different types of parties (keygen, signing, dynamic groups)
func BaseUpdate(p Party, msg ParsedMessage, task string) (ok bool, err *Error) {
	_ = "STUB: not implemented"
	// fast-fail on an invalid message; do not lock the mutex yet
	return false, nil
}

// lock the mutex. need this mtx unlock hook; L108 is recursive so cannot use defer

// data is written to P state below

// finished! the round implementation will have sent the data through the `end` channel.

// recursive so can't defer after return
// re-run round update or finish)
