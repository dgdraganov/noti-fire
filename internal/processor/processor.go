package processor

import (
	"context"
	"fmt"

	"github.com/dgdraganov/noti-fire/internal/model"
)

type processAction struct {
	publisher Publisher
}

// NewProcessAction i sa constructor function for the processAction type
func NewProcessAction(publisher Publisher) *processAction {
	return &processAction{
		publisher: publisher,
	}
}

// Execute is the function that implements the responsibility of the processAction
func (n *processAction) Execute(ctx context.Context, message string) error {
	eventMessage := model.EventMessage{
		Message: message,
	}
	err := n.publisher.Publish(ctx, eventMessage)
	if err != nil {
		return fmt.Errorf("publish event: %w", err)
	}
	return nil
}
