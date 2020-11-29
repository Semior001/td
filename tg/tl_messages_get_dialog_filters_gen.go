// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// MessagesGetDialogFiltersRequest represents TL type `messages.getDialogFilters#f19ed96d`.
type MessagesGetDialogFiltersRequest struct {
}

// MessagesGetDialogFiltersRequestTypeID is TL type id of MessagesGetDialogFiltersRequest.
const MessagesGetDialogFiltersRequestTypeID = 0xf19ed96d

// Encode implements bin.Encoder.
func (g *MessagesGetDialogFiltersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogFilters#f19ed96d as nil")
	}
	b.PutID(MessagesGetDialogFiltersRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDialogFiltersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogFilters#f19ed96d to nil")
	}
	if err := b.ConsumeID(MessagesGetDialogFiltersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDialogFilters#f19ed96d: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetDialogFiltersRequest.
var (
	_ bin.Encoder = &MessagesGetDialogFiltersRequest{}
	_ bin.Decoder = &MessagesGetDialogFiltersRequest{}
)