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

// LangpackGetLanguageRequest represents TL type `langpack.getLanguage#6a596502`.
// Get information about a language in a localization pack
//
// See https://core.telegram.org/method/langpack.getLanguage for reference.
type LangpackGetLanguageRequest struct {
	// Language pack name
	LangPack string
	// Language code
	LangCode string
}

// LangpackGetLanguageRequestTypeID is TL type id of LangpackGetLanguageRequest.
const LangpackGetLanguageRequestTypeID = 0x6a596502

func (g *LangpackGetLanguageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangPack == "") {
		return false
	}
	if !(g.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *LangpackGetLanguageRequest) String() string {
	if g == nil {
		return "LangpackGetLanguageRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("LangpackGetLanguageRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tLangPack: ")
	sb.WriteString(fmt.Sprint(g.LangPack))
	sb.WriteString(",\n")
	sb.WriteString("\tLangCode: ")
	sb.WriteString(fmt.Sprint(g.LangCode))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *LangpackGetLanguageRequest) TypeID() uint32 {
	return LangpackGetLanguageRequestTypeID
}

// Encode implements bin.Encoder.
func (g *LangpackGetLanguageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode langpack.getLanguage#6a596502 as nil")
	}
	b.PutID(LangpackGetLanguageRequestTypeID)
	b.PutString(g.LangPack)
	b.PutString(g.LangCode)
	return nil
}

// GetLangPack returns value of LangPack field.
func (g *LangpackGetLanguageRequest) GetLangPack() (value string) {
	return g.LangPack
}

// GetLangCode returns value of LangCode field.
func (g *LangpackGetLanguageRequest) GetLangCode() (value string) {
	return g.LangCode
}

// Decode implements bin.Decoder.
func (g *LangpackGetLanguageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode langpack.getLanguage#6a596502 to nil")
	}
	if err := b.ConsumeID(LangpackGetLanguageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode langpack.getLanguage#6a596502: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getLanguage#6a596502: field lang_pack: %w", err)
		}
		g.LangPack = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getLanguage#6a596502: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for LangpackGetLanguageRequest.
var (
	_ bin.Encoder = &LangpackGetLanguageRequest{}
	_ bin.Decoder = &LangpackGetLanguageRequest{}
)

// LangpackGetLanguage invokes method langpack.getLanguage#6a596502 returning error if any.
// Get information about a language in a localization pack
//
// See https://core.telegram.org/method/langpack.getLanguage for reference.
func (c *Client) LangpackGetLanguage(ctx context.Context, request *LangpackGetLanguageRequest) (*LangPackLanguage, error) {
	var result LangPackLanguage

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
