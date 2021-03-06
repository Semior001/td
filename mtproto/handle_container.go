package mtproto

import (
	"golang.org/x/xerrors"

	"github.com/gotd/td/internal/proto"

	"github.com/gotd/td/bin"
)

func (c *Conn) handleContainer(b *bin.Buffer) error {
	var container proto.MessageContainer
	if err := container.Decode(b); err != nil {
		return xerrors.Errorf("container: %w", err)
	}
	for _, msg := range container.Messages {
		if err := c.processContainerMessage(msg); err != nil {
			return err
		}
	}
	return nil
}

func (c *Conn) processContainerMessage(msg proto.Message) error {
	b := &bin.Buffer{Buf: msg.Body}
	return c.handleMessage(b)
}
