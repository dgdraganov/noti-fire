package kafka

import (
	"context"
	"strings"

	"github.com/dgdraganov/noti-fire/pkg/config"
	"github.com/segmentio/kafka-go"
)

type kafkaProducer struct {
	writer *kafka.Writer
}

// NewKafkaProducer i sa constructor function for the kafkaProducer type
func NewKafkaProducer(config config.KafkaProducerConfig) *kafkaProducer {
	// dialer := &kafka.Dialer{
	//     Timeout:   10 * time.Second,
	//     DualStack: true,
	//     TLS:       &tls.Config{...tls config...},
	// }

	brokers := strings.Split(config.Brokers, ",")

	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		//Balancer: &kafka.Hash{},
	})
	return &kafkaProducer{
		writer: kafkaWriter,
	}
}

func (pr *kafkaProducer) Publish(ctx context.Context, topic string, msgs ...string) error {

	// not implemented
	return nil
}
