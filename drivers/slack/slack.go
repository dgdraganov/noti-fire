package slack

import "github.com/dgdraganov/noti-fire/internal/model"

type SlackDriver struct {
}

func NewSlackDriver() *SlackDriver {
	return &SlackDriver{}
}

// Send implements the dispatch.Driver interface
func (s *SlackDriver) Send(msg model.NotificationMessage) error {
	return nil
}
