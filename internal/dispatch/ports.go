package dispatch

import "github.com/dgdraganov/noti-fire/internal/model"

type Driver interface {
	Send(model.NotificationMessage) error
}
