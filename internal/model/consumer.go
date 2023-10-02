package model

import "github.com/segmentio/kafka-go"

type ConsumedMessage struct {
	Message kafka.Message
}
