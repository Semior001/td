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

// UploadReuploadCdnFileRequest represents TL type `upload.reuploadCdnFile#9b2754a8`.
// Request a reupload of a certain file to a CDN DC¹.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/method/upload.reuploadCdnFile for reference.
type UploadReuploadCdnFileRequest struct {
	// File token
	FileToken []byte
	// Request token
	RequestToken []byte
}

// UploadReuploadCdnFileRequestTypeID is TL type id of UploadReuploadCdnFileRequest.
const UploadReuploadCdnFileRequestTypeID = 0x9b2754a8

func (r *UploadReuploadCdnFileRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.FileToken == nil) {
		return false
	}
	if !(r.RequestToken == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *UploadReuploadCdnFileRequest) String() string {
	if r == nil {
		return "UploadReuploadCdnFileRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UploadReuploadCdnFileRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFileToken: ")
	sb.WriteString(fmt.Sprint(r.FileToken))
	sb.WriteString(",\n")
	sb.WriteString("\tRequestToken: ")
	sb.WriteString(fmt.Sprint(r.RequestToken))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *UploadReuploadCdnFileRequest) TypeID() uint32 {
	return UploadReuploadCdnFileRequestTypeID
}

// Encode implements bin.Encoder.
func (r *UploadReuploadCdnFileRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode upload.reuploadCdnFile#9b2754a8 as nil")
	}
	b.PutID(UploadReuploadCdnFileRequestTypeID)
	b.PutBytes(r.FileToken)
	b.PutBytes(r.RequestToken)
	return nil
}

// GetFileToken returns value of FileToken field.
func (r *UploadReuploadCdnFileRequest) GetFileToken() (value []byte) {
	return r.FileToken
}

// GetRequestToken returns value of RequestToken field.
func (r *UploadReuploadCdnFileRequest) GetRequestToken() (value []byte) {
	return r.RequestToken
}

// Decode implements bin.Decoder.
func (r *UploadReuploadCdnFileRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode upload.reuploadCdnFile#9b2754a8 to nil")
	}
	if err := b.ConsumeID(UploadReuploadCdnFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.reuploadCdnFile#9b2754a8: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.reuploadCdnFile#9b2754a8: field file_token: %w", err)
		}
		r.FileToken = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.reuploadCdnFile#9b2754a8: field request_token: %w", err)
		}
		r.RequestToken = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UploadReuploadCdnFileRequest.
var (
	_ bin.Encoder = &UploadReuploadCdnFileRequest{}
	_ bin.Decoder = &UploadReuploadCdnFileRequest{}
)

// UploadReuploadCdnFile invokes method upload.reuploadCdnFile#9b2754a8 returning error if any.
// Request a reupload of a certain file to a CDN DC¹.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// Possible errors:
//  400 RSA_DECRYPT_FAILED: Internal RSA decryption failed
//
// See https://core.telegram.org/method/upload.reuploadCdnFile for reference.
// Can be used by bots.
func (c *Client) UploadReuploadCdnFile(ctx context.Context, request *UploadReuploadCdnFileRequest) ([]FileHash, error) {
	var result FileHashVector

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}
