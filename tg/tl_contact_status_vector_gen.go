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

// ContactStatusVector is a box for Vector<ContactStatus>
type ContactStatusVector struct {
	// Elements of Vector<ContactStatus>
	Elems []ContactStatus
}

// ContactStatusVectorTypeID is TL type id of ContactStatusVector.
const ContactStatusVectorTypeID = bin.TypeVector

func (vec *ContactStatusVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *ContactStatusVector) String() string {
	if vec == nil {
		return "ContactStatusVector(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ContactStatusVector")
	sb.WriteByte('[')
	for _, e := range vec.Elems {
		sb.WriteString(fmt.Sprint(e) + ",\n")
	}
	sb.WriteByte(']')
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (vec *ContactStatusVector) TypeID() uint32 {
	return ContactStatusVectorTypeID
}

// Encode implements bin.Encoder.
func (vec *ContactStatusVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<ContactStatus> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<ContactStatus>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *ContactStatusVector) GetElems() (value []ContactStatus) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *ContactStatusVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<ContactStatus> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<ContactStatus>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ContactStatus
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<ContactStatus>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactStatusVector.
var (
	_ bin.Encoder = &ContactStatusVector{}
	_ bin.Decoder = &ContactStatusVector{}
)
