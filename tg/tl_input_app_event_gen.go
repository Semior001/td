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

// InputAppEvent represents TL type `inputAppEvent#1d1b1245`.
// Event that occured in the application.
//
// See https://core.telegram.org/constructor/inputAppEvent for reference.
type InputAppEvent struct {
	// Client's exact timestamp for the event
	Time float64
	// Type of event
	Type string
	// Arbitrary numeric value for more convenient selection of certain event types, or events referring to a certain object
	Peer int64
	// Details of the event
	Data JSONValueClass
}

// InputAppEventTypeID is TL type id of InputAppEvent.
const InputAppEventTypeID = 0x1d1b1245

func (i *InputAppEvent) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Time == 0) {
		return false
	}
	if !(i.Type == "") {
		return false
	}
	if !(i.Peer == 0) {
		return false
	}
	if !(i.Data == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputAppEvent) String() string {
	if i == nil {
		return "InputAppEvent(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputAppEvent")
	sb.WriteString("{\n")
	sb.WriteString("\tTime: ")
	sb.WriteString(fmt.Sprint(i.Time))
	sb.WriteString(",\n")
	sb.WriteString("\tType: ")
	sb.WriteString(fmt.Sprint(i.Type))
	sb.WriteString(",\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(i.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tData: ")
	sb.WriteString(fmt.Sprint(i.Data))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputAppEvent) TypeID() uint32 {
	return InputAppEventTypeID
}

// Encode implements bin.Encoder.
func (i *InputAppEvent) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputAppEvent#1d1b1245 as nil")
	}
	b.PutID(InputAppEventTypeID)
	b.PutDouble(i.Time)
	b.PutString(i.Type)
	b.PutLong(i.Peer)
	if i.Data == nil {
		return fmt.Errorf("unable to encode inputAppEvent#1d1b1245: field data is nil")
	}
	if err := i.Data.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputAppEvent#1d1b1245: field data: %w", err)
	}
	return nil
}

// GetTime returns value of Time field.
func (i *InputAppEvent) GetTime() (value float64) {
	return i.Time
}

// GetType returns value of Type field.
func (i *InputAppEvent) GetType() (value string) {
	return i.Type
}

// GetPeer returns value of Peer field.
func (i *InputAppEvent) GetPeer() (value int64) {
	return i.Peer
}

// GetData returns value of Data field.
func (i *InputAppEvent) GetData() (value JSONValueClass) {
	return i.Data
}

// Decode implements bin.Decoder.
func (i *InputAppEvent) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputAppEvent#1d1b1245 to nil")
	}
	if err := b.ConsumeID(InputAppEventTypeID); err != nil {
		return fmt.Errorf("unable to decode inputAppEvent#1d1b1245: %w", err)
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode inputAppEvent#1d1b1245: field time: %w", err)
		}
		i.Time = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputAppEvent#1d1b1245: field type: %w", err)
		}
		i.Type = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputAppEvent#1d1b1245: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := DecodeJSONValue(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputAppEvent#1d1b1245: field data: %w", err)
		}
		i.Data = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputAppEvent.
var (
	_ bin.Encoder = &InputAppEvent{}
	_ bin.Decoder = &InputAppEvent{}
)
