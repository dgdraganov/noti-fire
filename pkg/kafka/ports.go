package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Writer interface {
	WriteMessage(ctx context.Context, msgs kafka.Message) error
}
