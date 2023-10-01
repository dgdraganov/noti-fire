package kafka

import "context"

type MessageWriter interface {
	WriteMessage(context.Context, []byte)
}
