package config

import (
	"fmt"
	"os"
)

type KafkaConsumerConfig struct {
	Brokers string // comma-separated broker addresses
	Topic   string
}

type ConsumerConfig struct {
	AppEnv       string
	ConsumerName string
	KafkaConsumerConfig
}

// NewConsumerConfig is a constructor function for the ConsumerConfig type
func NewConsumerConfig() (ConsumerConfig, error) {
	appEnv, ok := os.LookupEnv("APP_ENV")
	if !ok {
		return ConsumerConfig{}, fmt.Errorf("%w: APP_ENV", errMissingEnvVariable)
	}

	consumerName, ok := os.LookupEnv("CONSUMER_NAME")
	if !ok {
		return ConsumerConfig{}, fmt.Errorf("%w: CONSUMER_NAME", errMissingEnvVariable)
	}

	kafkaBrokers, ok := os.LookupEnv("KAFKA_BROKERS")
	if !ok {
		return ConsumerConfig{}, fmt.Errorf("%w: KAFKA_BROKERS", errMissingEnvVariable)
	}

	kafkaTopic, ok := os.LookupEnv("KAFKA_TOPIC")
	if !ok {
		return ConsumerConfig{}, fmt.Errorf("%w: KAFKA_TOPIC", errMissingEnvVariable)
	}

	kafkaConfig := KafkaConsumerConfig{
		Brokers: kafkaBrokers,
		Topic:   kafkaTopic,
	}

	return ConsumerConfig{
		AppEnv:              appEnv,
		ConsumerName:        consumerName,
		KafkaConsumerConfig: kafkaConfig,
	}, nil
}
