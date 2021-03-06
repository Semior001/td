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

// StickerSetCoveredClassVector is a box for Vector<StickerSetCovered>
type StickerSetCoveredClassVector struct {
	// Elements of Vector<StickerSetCovered>
	Elems []StickerSetCoveredClass
}

// StickerSetCoveredClassVectorTypeID is TL type id of StickerSetCoveredClassVector.
const StickerSetCoveredClassVectorTypeID = bin.TypeVector

func (vec *StickerSetCoveredClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *StickerSetCoveredClassVector) String() string {
	if vec == nil {
		return "StickerSetCoveredClassVector(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StickerSetCoveredClassVector")
	sb.WriteByte('[')
	for _, e := range vec.Elems {
		sb.WriteString(fmt.Sprint(e) + ",\n")
	}
	sb.WriteByte(']')
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (vec *StickerSetCoveredClassVector) TypeID() uint32 {
	return StickerSetCoveredClassVectorTypeID
}

// Encode implements bin.Encoder.
func (vec *StickerSetCoveredClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<StickerSetCovered> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<StickerSetCovered>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<StickerSetCovered>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *StickerSetCoveredClassVector) GetElems() (value []StickerSetCoveredClass) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *StickerSetCoveredClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<StickerSetCovered> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<StickerSetCovered>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStickerSetCovered(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<StickerSetCovered>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for StickerSetCoveredClassVector.
var (
	_ bin.Encoder = &StickerSetCoveredClassVector{}
	_ bin.Decoder = &StickerSetCoveredClassVector{}
)
