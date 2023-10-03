package notify

import (
	"context"

	"github.com/dgdraganov/noti-fire/internal/model"
	"go.uber.org/zap"
)

type notifyer struct {
	consumer   Consumer
	dispatcher Dispatcher
	logs       *zap.SugaredLogger
}

// NewNotifyer is a constructor function for the notifyer type
func NewNotifyer(consumer Consumer, dispatcher Dispatcher, logger *zap.SugaredLogger) *notifyer {
	return &notifyer{
		consumer:   consumer,
		dispatcher: dispatcher,
		logs:       logger,
	}
}

// Process will listen asynchronously for new messages from the consumer and will dispatch them through the dispatcher
func (n *notifyer) Process(ctx context.Context) {
	go n.process(ctx)
}

func (n *notifyer) process(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		consumed, err := n.consumer.Consume(ctx)
		if err != nil {
			n.logs.Errorf(
				"consumer failed to fetch message",
				"error", err,
			)
		}

		dispatchMessage := model.NotificationMessage{
			Message: string(consumed.Message.Value),
		}
		n.dispatcher.Dispatch(dispatchMessage)

		err = n.consumer.MarkConsumed(ctx, consumed)
		if err != nil {
			n.logs.Errorf(
				"consumer failed to acknowledge message",
				"error", err,
			)
		}
	}
}
