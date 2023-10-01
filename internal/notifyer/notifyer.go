package notifyer

import (
	"context"
	"fmt"

	"github.com/dgdraganov/noti-fire/internal/model"
)

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
	eventMessage := model.EventMessage{
		Message: message,
	}
	err := n.publisher.Publish(ctx, eventMessage)
	if err != nil {
		return fmt.Errorf("publish event: %w", err)
	}
	return nil
}
