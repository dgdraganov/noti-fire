package kafka

import (
	"context"
	"fmt"
	"strings"

	"github.com/dgdraganov/noti-fire/internal/model"
	"github.com/dgdraganov/noti-fire/pkg/config"
	"github.com/segmentio/kafka-go"
)

type kafkaReader struct {
	reader *kafka.Reader
}

// NewKafkaReader is sa constructor function for the kafkaReader type
func NewKafkaReader(config config.KafkaConsumerConfig) *kafkaReader {

	brokers := strings.Split(config.Brokers, ",")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  "notifications_consumer",
		Topic:    config.Topic,
		MaxBytes: 10e6, // 10MB
	})

	return &kafkaReader{
		reader: reader,
	}
}

// ReadMessage consumes a single message from p.reader's topic
func (p *kafkaReader) ReadMessage(ctx context.Context) (*model.ConsumedMessage, error) {
	msg, err := p.reader.FetchMessage(ctx)
	if err != nil {
		return nil, fmt.Errorf("reader fetch message: %w", err)
	}

	return &model.ConsumedMessage{msg}, nil
}

// CommitMessage commits the given message by updating the offset value of the respective partition
func (p *kafkaReader) CommitMessage(ctx context.Context, msg *model.ConsumedMessage) error {
	if err := p.reader.CommitMessages(ctx, msg.Message); err != nil {
		return fmt.Errorf("comitting message: %w", err)
	}
	return nil
}
