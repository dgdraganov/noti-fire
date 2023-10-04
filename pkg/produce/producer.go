package produce

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgdraganov/noti-fire/internal/model"
)

type messageProducer struct {
	writer MessageWriter
}

// NewMessageProducer is sa constructor function for the messageProducer type
func NewMessageProducer(writer MessageWriter) *messageProducer {
	return &messageProducer{
		writer: writer,
	}
}

// Publish uses the message writer in order to publish a message
func (pr *messageProducer) Publish(ctx context.Context, msg model.EventMessage) error {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	if err := pr.writer.WriteMessage(ctx, bytes); err != nil {
		return fmt.Errorf("write message: %w", err)
	}

	return nil
}
