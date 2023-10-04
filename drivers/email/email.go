package email

import "github.com/dgdraganov/noti-fire/internal/model"

type EmailDriver struct {
}

func NewEmailDriver() *EmailDriver {
	return &EmailDriver{}
}

// Send implements the dispatch.Driver interface
func (s *EmailDriver) Send(msg model.NotificationMessage) error {
	return nil
}
