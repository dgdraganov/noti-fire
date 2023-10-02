package sms

import "github.com/dgdraganov/noti-fire/internal/model"

type SMSDriver struct {
}

func NewSMSDriver() *SMSDriver {
	return &SMSDriver{}
}

func (s *SMSDriver) Send(msg model.NotificationMessage) error {
	return nil
}
