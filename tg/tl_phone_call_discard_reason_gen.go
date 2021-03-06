// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// PhoneCallDiscardReasonMissed represents TL type `phoneCallDiscardReasonMissed#85e42301`.
// The phone call was missed
//
// See https://core.telegram.org/constructor/phoneCallDiscardReasonMissed for reference.
type PhoneCallDiscardReasonMissed struct {
}

// PhoneCallDiscardReasonMissedTypeID is TL type id of PhoneCallDiscardReasonMissed.
const PhoneCallDiscardReasonMissedTypeID = 0x85e42301

func (p *PhoneCallDiscardReasonMissed) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhoneCallDiscardReasonMissed) String() string {
	if p == nil {
		return "PhoneCallDiscardReasonMissed(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneCallDiscardReasonMissed")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PhoneCallDiscardReasonMissed) TypeID() uint32 {
	return PhoneCallDiscardReasonMissedTypeID
}

// Encode implements bin.Encoder.
func (p *PhoneCallDiscardReasonMissed) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallDiscardReasonMissed#85e42301 as nil")
	}
	b.PutID(PhoneCallDiscardReasonMissedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhoneCallDiscardReasonMissed) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallDiscardReasonMissed#85e42301 to nil")
	}
	if err := b.ConsumeID(PhoneCallDiscardReasonMissedTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCallDiscardReasonMissed#85e42301: %w", err)
	}
	return nil
}

// construct implements constructor of PhoneCallDiscardReasonClass.
func (p PhoneCallDiscardReasonMissed) construct() PhoneCallDiscardReasonClass { return &p }

// Ensuring interfaces in compile-time for PhoneCallDiscardReasonMissed.
var (
	_ bin.Encoder = &PhoneCallDiscardReasonMissed{}
	_ bin.Decoder = &PhoneCallDiscardReasonMissed{}

	_ PhoneCallDiscardReasonClass = &PhoneCallDiscardReasonMissed{}
)

// PhoneCallDiscardReasonDisconnect represents TL type `phoneCallDiscardReasonDisconnect#e095c1a0`.
// The phone call was disconnected
//
// See https://core.telegram.org/constructor/phoneCallDiscardReasonDisconnect for reference.
type PhoneCallDiscardReasonDisconnect struct {
}

// PhoneCallDiscardReasonDisconnectTypeID is TL type id of PhoneCallDiscardReasonDisconnect.
const PhoneCallDiscardReasonDisconnectTypeID = 0xe095c1a0

func (p *PhoneCallDiscardReasonDisconnect) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhoneCallDiscardReasonDisconnect) String() string {
	if p == nil {
		return "PhoneCallDiscardReasonDisconnect(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneCallDiscardReasonDisconnect")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PhoneCallDiscardReasonDisconnect) TypeID() uint32 {
	return PhoneCallDiscardReasonDisconnectTypeID
}

// Encode implements bin.Encoder.
func (p *PhoneCallDiscardReasonDisconnect) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallDiscardReasonDisconnect#e095c1a0 as nil")
	}
	b.PutID(PhoneCallDiscardReasonDisconnectTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhoneCallDiscardReasonDisconnect) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallDiscardReasonDisconnect#e095c1a0 to nil")
	}
	if err := b.ConsumeID(PhoneCallDiscardReasonDisconnectTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCallDiscardReasonDisconnect#e095c1a0: %w", err)
	}
	return nil
}

// construct implements constructor of PhoneCallDiscardReasonClass.
func (p PhoneCallDiscardReasonDisconnect) construct() PhoneCallDiscardReasonClass { return &p }

// Ensuring interfaces in compile-time for PhoneCallDiscardReasonDisconnect.
var (
	_ bin.Encoder = &PhoneCallDiscardReasonDisconnect{}
	_ bin.Decoder = &PhoneCallDiscardReasonDisconnect{}

	_ PhoneCallDiscardReasonClass = &PhoneCallDiscardReasonDisconnect{}
)

// PhoneCallDiscardReasonHangup represents TL type `phoneCallDiscardReasonHangup#57adc690`.
// The phone call was ended normally
//
// See https://core.telegram.org/constructor/phoneCallDiscardReasonHangup for reference.
type PhoneCallDiscardReasonHangup struct {
}

// PhoneCallDiscardReasonHangupTypeID is TL type id of PhoneCallDiscardReasonHangup.
const PhoneCallDiscardReasonHangupTypeID = 0x57adc690

func (p *PhoneCallDiscardReasonHangup) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhoneCallDiscardReasonHangup) String() string {
	if p == nil {
		return "PhoneCallDiscardReasonHangup(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneCallDiscardReasonHangup")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PhoneCallDiscardReasonHangup) TypeID() uint32 {
	return PhoneCallDiscardReasonHangupTypeID
}

// Encode implements bin.Encoder.
func (p *PhoneCallDiscardReasonHangup) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallDiscardReasonHangup#57adc690 as nil")
	}
	b.PutID(PhoneCallDiscardReasonHangupTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhoneCallDiscardReasonHangup) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallDiscardReasonHangup#57adc690 to nil")
	}
	if err := b.ConsumeID(PhoneCallDiscardReasonHangupTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCallDiscardReasonHangup#57adc690: %w", err)
	}
	return nil
}

// construct implements constructor of PhoneCallDiscardReasonClass.
func (p PhoneCallDiscardReasonHangup) construct() PhoneCallDiscardReasonClass { return &p }

// Ensuring interfaces in compile-time for PhoneCallDiscardReasonHangup.
var (
	_ bin.Encoder = &PhoneCallDiscardReasonHangup{}
	_ bin.Decoder = &PhoneCallDiscardReasonHangup{}

	_ PhoneCallDiscardReasonClass = &PhoneCallDiscardReasonHangup{}
)

// PhoneCallDiscardReasonBusy represents TL type `phoneCallDiscardReasonBusy#faf7e8c9`.
// The phone call was discared because the user is busy in another call
//
// See https://core.telegram.org/constructor/phoneCallDiscardReasonBusy for reference.
type PhoneCallDiscardReasonBusy struct {
}

// PhoneCallDiscardReasonBusyTypeID is TL type id of PhoneCallDiscardReasonBusy.
const PhoneCallDiscardReasonBusyTypeID = 0xfaf7e8c9

func (p *PhoneCallDiscardReasonBusy) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhoneCallDiscardReasonBusy) String() string {
	if p == nil {
		return "PhoneCallDiscardReasonBusy(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneCallDiscardReasonBusy")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PhoneCallDiscardReasonBusy) TypeID() uint32 {
	return PhoneCallDiscardReasonBusyTypeID
}

// Encode implements bin.Encoder.
func (p *PhoneCallDiscardReasonBusy) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallDiscardReasonBusy#faf7e8c9 as nil")
	}
	b.PutID(PhoneCallDiscardReasonBusyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhoneCallDiscardReasonBusy) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallDiscardReasonBusy#faf7e8c9 to nil")
	}
	if err := b.ConsumeID(PhoneCallDiscardReasonBusyTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCallDiscardReasonBusy#faf7e8c9: %w", err)
	}
	return nil
}

// construct implements constructor of PhoneCallDiscardReasonClass.
func (p PhoneCallDiscardReasonBusy) construct() PhoneCallDiscardReasonClass { return &p }

// Ensuring interfaces in compile-time for PhoneCallDiscardReasonBusy.
var (
	_ bin.Encoder = &PhoneCallDiscardReasonBusy{}
	_ bin.Decoder = &PhoneCallDiscardReasonBusy{}

	_ PhoneCallDiscardReasonClass = &PhoneCallDiscardReasonBusy{}
)

// PhoneCallDiscardReasonClass represents PhoneCallDiscardReason generic type.
//
// See https://core.telegram.org/type/PhoneCallDiscardReason for reference.
//
// Example:
//  g, err := DecodePhoneCallDiscardReason(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *PhoneCallDiscardReasonMissed: // phoneCallDiscardReasonMissed#85e42301
//  case *PhoneCallDiscardReasonDisconnect: // phoneCallDiscardReasonDisconnect#e095c1a0
//  case *PhoneCallDiscardReasonHangup: // phoneCallDiscardReasonHangup#57adc690
//  case *PhoneCallDiscardReasonBusy: // phoneCallDiscardReasonBusy#faf7e8c9
//  default: panic(v)
//  }
type PhoneCallDiscardReasonClass interface {
	bin.Encoder
	bin.Decoder
	construct() PhoneCallDiscardReasonClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodePhoneCallDiscardReason implements binary de-serialization for PhoneCallDiscardReasonClass.
func DecodePhoneCallDiscardReason(buf *bin.Buffer) (PhoneCallDiscardReasonClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PhoneCallDiscardReasonMissedTypeID:
		// Decoding phoneCallDiscardReasonMissed#85e42301.
		v := PhoneCallDiscardReasonMissed{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneCallDiscardReasonClass: %w", err)
		}
		return &v, nil
	case PhoneCallDiscardReasonDisconnectTypeID:
		// Decoding phoneCallDiscardReasonDisconnect#e095c1a0.
		v := PhoneCallDiscardReasonDisconnect{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneCallDiscardReasonClass: %w", err)
		}
		return &v, nil
	case PhoneCallDiscardReasonHangupTypeID:
		// Decoding phoneCallDiscardReasonHangup#57adc690.
		v := PhoneCallDiscardReasonHangup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneCallDiscardReasonClass: %w", err)
		}
		return &v, nil
	case PhoneCallDiscardReasonBusyTypeID:
		// Decoding phoneCallDiscardReasonBusy#faf7e8c9.
		v := PhoneCallDiscardReasonBusy{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneCallDiscardReasonClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PhoneCallDiscardReasonClass: %w", bin.NewUnexpectedID(id))
	}
}

// PhoneCallDiscardReason boxes the PhoneCallDiscardReasonClass providing a helper.
type PhoneCallDiscardReasonBox struct {
	PhoneCallDiscardReason PhoneCallDiscardReasonClass
}

// Decode implements bin.Decoder for PhoneCallDiscardReasonBox.
func (b *PhoneCallDiscardReasonBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PhoneCallDiscardReasonBox to nil")
	}
	v, err := DecodePhoneCallDiscardReason(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PhoneCallDiscardReason = v
	return nil
}

// Encode implements bin.Encode for PhoneCallDiscardReasonBox.
func (b *PhoneCallDiscardReasonBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PhoneCallDiscardReason == nil {
		return fmt.Errorf("unable to encode PhoneCallDiscardReasonClass as nil")
	}
	return b.PhoneCallDiscardReason.Encode(buf)
}
