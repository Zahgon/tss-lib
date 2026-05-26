// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package tss

// fundamental is an error that has a message and a stack, but no caller.
type Error struct {
	cause    error
	task     string
	round    int
	victim   *PartyID
	culprits []*PartyID
}

func NewError(err error, task string, round int, victim *PartyID, culprits ...*PartyID) *Error {
	_ = "STUB: not implemented"
	return nil
}

func (err *Error) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (err *Error) Cause() error { _ = "STUB: not implemented"; return nil }

func (err *Error) Task() string { _ = "STUB: not implemented"; return "" }

func (err *Error) Round() int { _ = "STUB: not implemented"; return 0 }

func (err *Error) Victim() *PartyID { _ = "STUB: not implemented"; return nil }

func (err *Error) Culprits() []*PartyID { _ = "STUB: not implemented"; return nil }

func (err *Error) Error() string { _ = "STUB: not implemented"; return "" }
