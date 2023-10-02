package consumer

import (
	"context"
	"fmt"

	"github.com/dgdraganov/noti-fire/internal/model"
)

type messageConsumer struct {
	reader ReadCommitter
}

// NewMessageConsumer is sa constructor function for the messageConsumer type
func NewMessageConsumer(reader ReadCommitter) *messageConsumer {
	return &messageConsumer{
		reader: reader,
	}
}

// Consume uses the reader in order to consume a pending message
func (c *messageConsumer) Consume(ctx context.Context) (*model.ConsumedMessage, error) {
	msg, err := c.reader.ReadMessage(ctx)
	if err != nil {
		return nil, fmt.Errorf("reader read message: %w", err)
	}
	return msg, nil
}

// MarkConsumed makes sure that the message is accepted and will not be consumed again
func (c *messageConsumer) MarkConsumed(ctx context.Context, msg *model.ConsumedMessage) error {
	err := c.reader.CommitMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("reader commit message: %w", err)
	}
	return nil
}
