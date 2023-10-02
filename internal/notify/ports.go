package notify

import (
	"context"

	"github.com/dgdraganov/noti-fire/internal/model"
)

type Consumer interface {
	Consume(ctx context.Context) (*model.ConsumedMessage, error)
	MarkConsumed(ctx context.Context, msg *model.ConsumedMessage) error
}

type Dispatcher interface {
	Dispatch(model.NotificationMessage)
}
