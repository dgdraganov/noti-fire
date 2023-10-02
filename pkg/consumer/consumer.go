package consumer

import (
	"context"

	"github.com/dgdraganov/noti-fire/internal/model"
)

type Reader interface {
	ReadMessage(ctx context.Context) (*model.ConsumedMessage, error)
	CommitMessage(ctx context.Context, msg *model.ConsumedMessage) error
}
