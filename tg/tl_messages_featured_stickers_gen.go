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

// MessagesFeaturedStickersNotModified represents TL type `messages.featuredStickersNotModified#c6dc0c66`.
// Featured stickers haven't changed
//
// See https://core.telegram.org/constructor/messages.featuredStickersNotModified for reference.
type MessagesFeaturedStickersNotModified struct {
	// Total number of featured stickers
	Count int
}

// MessagesFeaturedStickersNotModifiedTypeID is TL type id of MessagesFeaturedStickersNotModified.
const MessagesFeaturedStickersNotModifiedTypeID = 0xc6dc0c66

func (f *MessagesFeaturedStickersNotModified) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFeaturedStickersNotModified) String() string {
	if f == nil {
		return "MessagesFeaturedStickersNotModified(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesFeaturedStickersNotModified")
	sb.WriteString("{\n")
	sb.WriteString("\tCount: ")
	sb.WriteString(fmt.Sprint(f.Count))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *MessagesFeaturedStickersNotModified) TypeID() uint32 {
	return MessagesFeaturedStickersNotModifiedTypeID
}

// Encode implements bin.Encoder.
func (f *MessagesFeaturedStickersNotModified) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.featuredStickersNotModified#c6dc0c66 as nil")
	}
	b.PutID(MessagesFeaturedStickersNotModifiedTypeID)
	b.PutInt(f.Count)
	return nil
}

// GetCount returns value of Count field.
func (f *MessagesFeaturedStickersNotModified) GetCount() (value int) {
	return f.Count
}

// Decode implements bin.Decoder.
func (f *MessagesFeaturedStickersNotModified) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.featuredStickersNotModified#c6dc0c66 to nil")
	}
	if err := b.ConsumeID(MessagesFeaturedStickersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.featuredStickersNotModified#c6dc0c66: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickersNotModified#c6dc0c66: field count: %w", err)
		}
		f.Count = value
	}
	return nil
}

// construct implements constructor of MessagesFeaturedStickersClass.
func (f MessagesFeaturedStickersNotModified) construct() MessagesFeaturedStickersClass { return &f }

// Ensuring interfaces in compile-time for MessagesFeaturedStickersNotModified.
var (
	_ bin.Encoder = &MessagesFeaturedStickersNotModified{}
	_ bin.Decoder = &MessagesFeaturedStickersNotModified{}

	_ MessagesFeaturedStickersClass = &MessagesFeaturedStickersNotModified{}
)

// MessagesFeaturedStickers represents TL type `messages.featuredStickers#b6abc341`.
// Featured stickersets
//
// See https://core.telegram.org/constructor/messages.featuredStickers for reference.
type MessagesFeaturedStickers struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// Total number of featured stickers
	Count int
	// Featured stickersets
	Sets []StickerSetCoveredClass
	// IDs of new featured stickersets
	Unread []int64
}

// MessagesFeaturedStickersTypeID is TL type id of MessagesFeaturedStickers.
const MessagesFeaturedStickersTypeID = 0xb6abc341

func (f *MessagesFeaturedStickers) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Hash == 0) {
		return false
	}
	if !(f.Count == 0) {
		return false
	}
	if !(f.Sets == nil) {
		return false
	}
	if !(f.Unread == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFeaturedStickers) String() string {
	if f == nil {
		return "MessagesFeaturedStickers(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesFeaturedStickers")
	sb.WriteString("{\n")
	sb.WriteString("\tHash: ")
	sb.WriteString(fmt.Sprint(f.Hash))
	sb.WriteString(",\n")
	sb.WriteString("\tCount: ")
	sb.WriteString(fmt.Sprint(f.Count))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range f.Sets {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range f.Unread {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *MessagesFeaturedStickers) TypeID() uint32 {
	return MessagesFeaturedStickersTypeID
}

// Encode implements bin.Encoder.
func (f *MessagesFeaturedStickers) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.featuredStickers#b6abc341 as nil")
	}
	b.PutID(MessagesFeaturedStickersTypeID)
	b.PutInt(f.Hash)
	b.PutInt(f.Count)
	b.PutVectorHeader(len(f.Sets))
	for idx, v := range f.Sets {
		if v == nil {
			return fmt.Errorf("unable to encode messages.featuredStickers#b6abc341: field sets element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.featuredStickers#b6abc341: field sets element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Unread))
	for _, v := range f.Unread {
		b.PutLong(v)
	}
	return nil
}

// GetHash returns value of Hash field.
func (f *MessagesFeaturedStickers) GetHash() (value int) {
	return f.Hash
}

// GetCount returns value of Count field.
func (f *MessagesFeaturedStickers) GetCount() (value int) {
	return f.Count
}

// GetSets returns value of Sets field.
func (f *MessagesFeaturedStickers) GetSets() (value []StickerSetCoveredClass) {
	return f.Sets
}

// GetUnread returns value of Unread field.
func (f *MessagesFeaturedStickers) GetUnread() (value []int64) {
	return f.Unread
}

// Decode implements bin.Decoder.
func (f *MessagesFeaturedStickers) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.featuredStickers#b6abc341 to nil")
	}
	if err := b.ConsumeID(MessagesFeaturedStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: field hash: %w", err)
		}
		f.Hash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: field count: %w", err)
		}
		f.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: field sets: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStickerSetCovered(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: field sets: %w", err)
			}
			f.Sets = append(f.Sets, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: field unread: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: field unread: %w", err)
			}
			f.Unread = append(f.Unread, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesFeaturedStickersClass.
func (f MessagesFeaturedStickers) construct() MessagesFeaturedStickersClass { return &f }

// Ensuring interfaces in compile-time for MessagesFeaturedStickers.
var (
	_ bin.Encoder = &MessagesFeaturedStickers{}
	_ bin.Decoder = &MessagesFeaturedStickers{}

	_ MessagesFeaturedStickersClass = &MessagesFeaturedStickers{}
)

// MessagesFeaturedStickersClass represents messages.FeaturedStickers generic type.
//
// See https://core.telegram.org/type/messages.FeaturedStickers for reference.
//
// Example:
//  g, err := DecodeMessagesFeaturedStickers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *MessagesFeaturedStickersNotModified: // messages.featuredStickersNotModified#c6dc0c66
//  case *MessagesFeaturedStickers: // messages.featuredStickers#b6abc341
//  default: panic(v)
//  }
type MessagesFeaturedStickersClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesFeaturedStickersClass

	// Total number of featured stickers
	GetCount() (value int)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeMessagesFeaturedStickers implements binary de-serialization for MessagesFeaturedStickersClass.
func DecodeMessagesFeaturedStickers(buf *bin.Buffer) (MessagesFeaturedStickersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesFeaturedStickersNotModifiedTypeID:
		// Decoding messages.featuredStickersNotModified#c6dc0c66.
		v := MessagesFeaturedStickersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFeaturedStickersClass: %w", err)
		}
		return &v, nil
	case MessagesFeaturedStickersTypeID:
		// Decoding messages.featuredStickers#b6abc341.
		v := MessagesFeaturedStickers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFeaturedStickersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesFeaturedStickersClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesFeaturedStickers boxes the MessagesFeaturedStickersClass providing a helper.
type MessagesFeaturedStickersBox struct {
	FeaturedStickers MessagesFeaturedStickersClass
}

// Decode implements bin.Decoder for MessagesFeaturedStickersBox.
func (b *MessagesFeaturedStickersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesFeaturedStickersBox to nil")
	}
	v, err := DecodeMessagesFeaturedStickers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.FeaturedStickers = v
	return nil
}

// Encode implements bin.Encode for MessagesFeaturedStickersBox.
func (b *MessagesFeaturedStickersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.FeaturedStickers == nil {
		return fmt.Errorf("unable to encode MessagesFeaturedStickersClass as nil")
	}
	return b.FeaturedStickers.Encode(buf)
}
