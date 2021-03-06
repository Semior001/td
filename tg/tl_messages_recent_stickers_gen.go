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

// MessagesRecentStickersNotModified represents TL type `messages.recentStickersNotModified#b17f890`.
// No new recent sticker was found
//
// See https://core.telegram.org/constructor/messages.recentStickersNotModified for reference.
type MessagesRecentStickersNotModified struct {
}

// MessagesRecentStickersNotModifiedTypeID is TL type id of MessagesRecentStickersNotModified.
const MessagesRecentStickersNotModifiedTypeID = 0xb17f890

func (r *MessagesRecentStickersNotModified) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesRecentStickersNotModified) String() string {
	if r == nil {
		return "MessagesRecentStickersNotModified(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesRecentStickersNotModified")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *MessagesRecentStickersNotModified) TypeID() uint32 {
	return MessagesRecentStickersNotModifiedTypeID
}

// Encode implements bin.Encoder.
func (r *MessagesRecentStickersNotModified) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.recentStickersNotModified#b17f890 as nil")
	}
	b.PutID(MessagesRecentStickersNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesRecentStickersNotModified) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.recentStickersNotModified#b17f890 to nil")
	}
	if err := b.ConsumeID(MessagesRecentStickersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.recentStickersNotModified#b17f890: %w", err)
	}
	return nil
}

// construct implements constructor of MessagesRecentStickersClass.
func (r MessagesRecentStickersNotModified) construct() MessagesRecentStickersClass { return &r }

// Ensuring interfaces in compile-time for MessagesRecentStickersNotModified.
var (
	_ bin.Encoder = &MessagesRecentStickersNotModified{}
	_ bin.Decoder = &MessagesRecentStickersNotModified{}

	_ MessagesRecentStickersClass = &MessagesRecentStickersNotModified{}
)

// MessagesRecentStickers represents TL type `messages.recentStickers#22f3afb3`.
// Recently used stickers
//
// See https://core.telegram.org/constructor/messages.recentStickers for reference.
type MessagesRecentStickers struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// Emojis associated to stickers
	Packs []StickerPack
	// Recent stickers
	Stickers []DocumentClass
	// When was each sticker last used
	Dates []int
}

// MessagesRecentStickersTypeID is TL type id of MessagesRecentStickers.
const MessagesRecentStickersTypeID = 0x22f3afb3

func (r *MessagesRecentStickers) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Hash == 0) {
		return false
	}
	if !(r.Packs == nil) {
		return false
	}
	if !(r.Stickers == nil) {
		return false
	}
	if !(r.Dates == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesRecentStickers) String() string {
	if r == nil {
		return "MessagesRecentStickers(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesRecentStickers")
	sb.WriteString("{\n")
	sb.WriteString("\tHash: ")
	sb.WriteString(fmt.Sprint(r.Hash))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range r.Packs {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range r.Stickers {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range r.Dates {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *MessagesRecentStickers) TypeID() uint32 {
	return MessagesRecentStickersTypeID
}

// Encode implements bin.Encoder.
func (r *MessagesRecentStickers) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.recentStickers#22f3afb3 as nil")
	}
	b.PutID(MessagesRecentStickersTypeID)
	b.PutInt(r.Hash)
	b.PutVectorHeader(len(r.Packs))
	for idx, v := range r.Packs {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.recentStickers#22f3afb3: field packs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(r.Stickers))
	for idx, v := range r.Stickers {
		if v == nil {
			return fmt.Errorf("unable to encode messages.recentStickers#22f3afb3: field stickers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.recentStickers#22f3afb3: field stickers element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(r.Dates))
	for _, v := range r.Dates {
		b.PutInt(v)
	}
	return nil
}

// GetHash returns value of Hash field.
func (r *MessagesRecentStickers) GetHash() (value int) {
	return r.Hash
}

// GetPacks returns value of Packs field.
func (r *MessagesRecentStickers) GetPacks() (value []StickerPack) {
	return r.Packs
}

// GetStickers returns value of Stickers field.
func (r *MessagesRecentStickers) GetStickers() (value []DocumentClass) {
	return r.Stickers
}

// GetDates returns value of Dates field.
func (r *MessagesRecentStickers) GetDates() (value []int) {
	return r.Dates
}

// Decode implements bin.Decoder.
func (r *MessagesRecentStickers) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.recentStickers#22f3afb3 to nil")
	}
	if err := b.ConsumeID(MessagesRecentStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.recentStickers#22f3afb3: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.recentStickers#22f3afb3: field hash: %w", err)
		}
		r.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.recentStickers#22f3afb3: field packs: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StickerPack
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.recentStickers#22f3afb3: field packs: %w", err)
			}
			r.Packs = append(r.Packs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.recentStickers#22f3afb3: field stickers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.recentStickers#22f3afb3: field stickers: %w", err)
			}
			r.Stickers = append(r.Stickers, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.recentStickers#22f3afb3: field dates: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.recentStickers#22f3afb3: field dates: %w", err)
			}
			r.Dates = append(r.Dates, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesRecentStickersClass.
func (r MessagesRecentStickers) construct() MessagesRecentStickersClass { return &r }

// Ensuring interfaces in compile-time for MessagesRecentStickers.
var (
	_ bin.Encoder = &MessagesRecentStickers{}
	_ bin.Decoder = &MessagesRecentStickers{}

	_ MessagesRecentStickersClass = &MessagesRecentStickers{}
)

// MessagesRecentStickersClass represents messages.RecentStickers generic type.
//
// See https://core.telegram.org/type/messages.RecentStickers for reference.
//
// Example:
//  g, err := DecodeMessagesRecentStickers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *MessagesRecentStickersNotModified: // messages.recentStickersNotModified#b17f890
//  case *MessagesRecentStickers: // messages.recentStickers#22f3afb3
//  default: panic(v)
//  }
type MessagesRecentStickersClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesRecentStickersClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeMessagesRecentStickers implements binary de-serialization for MessagesRecentStickersClass.
func DecodeMessagesRecentStickers(buf *bin.Buffer) (MessagesRecentStickersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesRecentStickersNotModifiedTypeID:
		// Decoding messages.recentStickersNotModified#b17f890.
		v := MessagesRecentStickersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesRecentStickersClass: %w", err)
		}
		return &v, nil
	case MessagesRecentStickersTypeID:
		// Decoding messages.recentStickers#22f3afb3.
		v := MessagesRecentStickers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesRecentStickersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesRecentStickersClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesRecentStickers boxes the MessagesRecentStickersClass providing a helper.
type MessagesRecentStickersBox struct {
	RecentStickers MessagesRecentStickersClass
}

// Decode implements bin.Decoder for MessagesRecentStickersBox.
func (b *MessagesRecentStickersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesRecentStickersBox to nil")
	}
	v, err := DecodeMessagesRecentStickers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.RecentStickers = v
	return nil
}

// Encode implements bin.Encode for MessagesRecentStickersBox.
func (b *MessagesRecentStickersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.RecentStickers == nil {
		return fmt.Errorf("unable to encode MessagesRecentStickersClass as nil")
	}
	return b.RecentStickers.Encode(buf)
}
