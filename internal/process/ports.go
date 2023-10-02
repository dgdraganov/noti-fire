package process

import (
	"context"

	"github.com/dgdraganov/noti-fire/internal/model"
)

type Publisher interface {
	Publish(ctx context.Context, msg model.EventMessage) error
}
