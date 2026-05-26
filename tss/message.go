// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package tss

import (
	"google.golang.org/protobuf/proto"
)

type (
	// Message describes the interface of the TSS Message for all protocols
	Message interface {
		// Type is encoded in the protobuf Any structure
		Type() string
		// The set of parties that this message should be sent to
		GetTo() []*PartyID
		// The party that this message is from
		GetFrom() *PartyID
		// Indicates whether the message should be broadcast to other participants
		IsBroadcast() bool
		// Indicates whether the message is to the old committee during re-sharing; used mainly in tests
		IsToOldCommittee() bool
		// Indicates whether the message is to both committees during re-sharing; used mainly in tests
		IsToOldAndNewCommittees() bool
		// Returns the encoded inner message bytes to send over the wire along with metadata about how the message should be delivered
		WireBytes() ([]byte, *MessageRouting, error)
		// Returns the protobuf message wrapper struct
		// Only its inner content should be sent over the wire, not this struct itself
		WireMsg() *MessageWrapper
		String() string
	}

	// ParsedMessage represents a message with inner ProtoBuf message content
	ParsedMessage interface {
		Message
		Content() MessageContent
		ValidateBasic() bool
	}

	// MessageContent represents a ProtoBuf message with validation logic
	MessageContent interface {
		proto.Message
		ValidateBasic() bool
	}

	// MessageRouting holds the full routing information for the message, consumed by the transport
	MessageRouting struct {
		// which participant this message came from
		From *PartyID
		// when `nil` the message should be broadcast to all parties
		To []*PartyID
		// whether the message should be broadcast to other participants
		IsBroadcast bool
		// whether the message should be sent to old committee participants rather than the new committee
		IsToOldCommittee bool
		// whether the message should be sent to both old and new committee participants
		IsToOldAndNewCommittees bool
	}

	// Implements ParsedMessage; this is a concrete implementation of what messages produced by a LocalParty look like
	MessageImpl struct {
		MessageRouting
		content MessageContent
		wire    *MessageWrapper
	}
)

var (
	_ Message       = (*MessageImpl)(nil)
	_ ParsedMessage = (*MessageImpl)(nil)
)

// ----- //

// NewMessageWrapper constructs a MessageWrapper from routing metadata and content
func NewMessageWrapper(routing MessageRouting, content MessageContent) *MessageWrapper {
	_ = "STUB: not implemented"
	// marshal the content to the ProtoBuf Any type
	return nil
}

// convert given PartyIDs to the wire format

// ----- //

func NewMessage(meta MessageRouting, content MessageContent, wire *MessageWrapper) ParsedMessage {
	_ = "STUB: not implemented"
	return *new(ParsedMessage)
}

func (mm *MessageImpl) Type() string { _ = "STUB: not implemented"; return "" }

func (mm *MessageImpl) GetTo() []*PartyID { _ = "STUB: not implemented"; return nil }

func (mm *MessageImpl) GetFrom() *PartyID { _ = "STUB: not implemented"; return nil }

func (mm *MessageImpl) IsBroadcast() bool { _ = "STUB: not implemented"; return false }

// only `true` in DGRound2Message (resharing)
func (mm *MessageImpl) IsToOldCommittee() bool { _ = "STUB: not implemented"; return false }

// only `true` in DGRound4Message (resharing)
func (mm *MessageImpl) IsToOldAndNewCommittees() bool { _ = "STUB: not implemented"; return false }

func (mm *MessageImpl) WireBytes() ([]byte, *MessageRouting, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (mm *MessageImpl) WireMsg() *MessageWrapper { _ = "STUB: not implemented"; return nil }

func (mm *MessageImpl) Content() MessageContent {
	_ = "STUB: not implemented"
	return *new(MessageContent)
}

func (mm *MessageImpl) ValidateBasic() bool { _ = "STUB: not implemented"; return false }

func (mm *MessageImpl) String() string { _ = "STUB: not implemented"; return "" }
