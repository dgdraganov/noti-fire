package notifyer

import (
	"context"

	"github.com/dgdraganov/noti-fire/internal/model"
)

type Consumer interface {
	Consume(ctx context.Context) (*model.ConsumedMessage, error)
	MarkConsumed(ctx context.Context, msg *model.ConsumedMessage) error
}
