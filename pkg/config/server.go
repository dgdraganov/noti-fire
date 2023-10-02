package config

import (
	"fmt"
	"os"
)

type KafkaProducerConfig struct {
	Brokers string // comma-separated broker addresses
	Topic   string
}

type ServerConfig struct {
	AppEnv     string
	ServerPort string
	ServerName string
	KafkaProducerConfig
}

// NewServerConfig i sa constructor function for the ServerConfig type
func NewServerConfig() (ServerConfig, error) {
	appEnv, ok := os.LookupEnv("APP_ENV")
	if !ok {
		return ServerConfig{}, fmt.Errorf("%w: APP_ENV", errMissingEnvVariable)
	}

	serverPort, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		return ServerConfig{}, fmt.Errorf("%w: SERVER_PORT", errMissingEnvVariable)
	}

	serverName, ok := os.LookupEnv("SERVER_NAME")
	if !ok {
		return ServerConfig{}, fmt.Errorf("%w: SERVER_NAME", errMissingEnvVariable)
	}

	brokers, ok := os.LookupEnv("KAFKA_BROKERS")
	if !ok {
		return ServerConfig{}, fmt.Errorf("%w: KAFKA_BROKERS", errMissingEnvVariable)
	}

	topic, ok := os.LookupEnv("KAFKA_TOPIC")
	if !ok {
		return ServerConfig{}, fmt.Errorf("%w: KAFKA_TOPIC", errMissingEnvVariable)
	}

	kafkaConfig := KafkaProducerConfig{
		Brokers: brokers,
		Topic:   topic,
	}

	return ServerConfig{
		AppEnv:              appEnv,
		ServerPort:          serverPort,
		ServerName:          serverName,
		KafkaProducerConfig: kafkaConfig,
	}, nil
}
