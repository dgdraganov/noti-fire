package notifyer

import "context"

type notifyerAction struct {
	publisher Publisher
}

// NewNotifyerAction i sa constructor function for the notifyerAction type
func NewNotifyerAction(publisher Publisher) *notifyerAction {
	return &notifyerAction{
		publisher: publisher,
	}
}

// Execute is the function that implements the responsibility of the notifyerAction
func (n *notifyerAction) Execute(ctx context.Context, message string) error {
	// not implemented
	return nil
}
