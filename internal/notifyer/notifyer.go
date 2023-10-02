package notifyer

type notifyer struct {
	consumer Consumer
}

// NewNotifyer i sa constructor function for the notifyer type
func NewNotifyer(consumer Consumer) *notifyer {
	return &notifyer{
		consumer: consumer,
	}
}
