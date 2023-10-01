package notifyer

import "context"

type Publisher interface {
	Publish(ctx context.Context, topic string, message string) error
}
