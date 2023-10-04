package kafka

import (
	"context"
	"fmt"
	"strings"

	"github.com/dgdraganov/noti-fire/pkg/config"
	"github.com/segmentio/kafka-go"
)

type kafkaWriter struct {
	writer *kafka.Writer
}

// NewKafkaWriter is sa constructor function for the kafkaWriter type
func NewKafkaWriter(config config.KafkaProducerConfig) *kafkaWriter {

	// todo: add tls config and certificates
	// dialer := &kafka.Dialer{
	//     Timeout:   10 * time.Second,
	//     DualStack: true,
	//     TLS:       &tls.Config{...tls config...},
	// }

	brokers := strings.Split(config.Brokers, ",")

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   config.Topic,
	})

	return &kafkaWriter{
		writer: writer,
	}
}

// WriteMessage is publishing a message to a specific kafka topic
func (pr *kafkaWriter) WriteMessage(ctx context.Context, msg []byte) error {
	kafkaMessage := kafka.Message{Value: msg}
	if err := pr.writer.WriteMessages(ctx, kafkaMessage); err != nil {
		return fmt.Errorf("write kafka message: %w", err)
	}
	return nil
}
