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

// MessagesAcceptEncryptionRequest represents TL type `messages.acceptEncryption#3dbc0415`.
// Confirms creation of a secret chat
//
// See https://core.telegram.org/method/messages.acceptEncryption for reference.
type MessagesAcceptEncryptionRequest struct {
	// Secret chat ID
	Peer InputEncryptedChat
	// B = g ^ b mod p, see Wikipedia¹
	//
	// Links:
	//  1) https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange
	GB []byte
	// 64-bit fingerprint of the received key
	KeyFingerprint int64
}

// MessagesAcceptEncryptionRequestTypeID is TL type id of MessagesAcceptEncryptionRequest.
const MessagesAcceptEncryptionRequestTypeID = 0x3dbc0415

func (a *MessagesAcceptEncryptionRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Peer.Zero()) {
		return false
	}
	if !(a.GB == nil) {
		return false
	}
	if !(a.KeyFingerprint == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *MessagesAcceptEncryptionRequest) String() string {
	if a == nil {
		return "MessagesAcceptEncryptionRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesAcceptEncryptionRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(a.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tGB: ")
	sb.WriteString(fmt.Sprint(a.GB))
	sb.WriteString(",\n")
	sb.WriteString("\tKeyFingerprint: ")
	sb.WriteString(fmt.Sprint(a.KeyFingerprint))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (a *MessagesAcceptEncryptionRequest) TypeID() uint32 {
	return MessagesAcceptEncryptionRequestTypeID
}

// Encode implements bin.Encoder.
func (a *MessagesAcceptEncryptionRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.acceptEncryption#3dbc0415 as nil")
	}
	b.PutID(MessagesAcceptEncryptionRequestTypeID)
	if err := a.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.acceptEncryption#3dbc0415: field peer: %w", err)
	}
	b.PutBytes(a.GB)
	b.PutLong(a.KeyFingerprint)
	return nil
}

// GetPeer returns value of Peer field.
func (a *MessagesAcceptEncryptionRequest) GetPeer() (value InputEncryptedChat) {
	return a.Peer
}

// GetGB returns value of GB field.
func (a *MessagesAcceptEncryptionRequest) GetGB() (value []byte) {
	return a.GB
}

// GetKeyFingerprint returns value of KeyFingerprint field.
func (a *MessagesAcceptEncryptionRequest) GetKeyFingerprint() (value int64) {
	return a.KeyFingerprint
}

// Decode implements bin.Decoder.
func (a *MessagesAcceptEncryptionRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.acceptEncryption#3dbc0415 to nil")
	}
	if err := b.ConsumeID(MessagesAcceptEncryptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.acceptEncryption#3dbc0415: %w", err)
	}
	{
		if err := a.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.acceptEncryption#3dbc0415: field peer: %w", err)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.acceptEncryption#3dbc0415: field g_b: %w", err)
		}
		a.GB = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.acceptEncryption#3dbc0415: field key_fingerprint: %w", err)
		}
		a.KeyFingerprint = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesAcceptEncryptionRequest.
var (
	_ bin.Encoder = &MessagesAcceptEncryptionRequest{}
	_ bin.Decoder = &MessagesAcceptEncryptionRequest{}
)

// MessagesAcceptEncryption invokes method messages.acceptEncryption#3dbc0415 returning error if any.
// Confirms creation of a secret chat
//
// Possible errors:
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 ENCRYPTION_ALREADY_ACCEPTED: Secret chat already accepted
//  400 ENCRYPTION_ALREADY_DECLINED: The secret chat was already declined
//
// See https://core.telegram.org/method/messages.acceptEncryption for reference.
func (c *Client) MessagesAcceptEncryption(ctx context.Context, request *MessagesAcceptEncryptionRequest) (EncryptedChatClass, error) {
	var result EncryptedChatBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.EncryptedChat, nil
}
