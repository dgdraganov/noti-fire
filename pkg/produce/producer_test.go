package produce_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dgdraganov/noti-fire/internal/model"
	"github.com/dgdraganov/noti-fire/pkg/produce"
)

// type messageProducer struct {
// 	writer MessageWriter
// }

// type MessageWriter interface {
//     WriteMessage(context.Context, []byte) error
// }

type messageWriterMock struct {
	writeMessage func(context.Context, []byte) error
}

func (w *messageWriterMock) WriteMessage(ctx context.Context, msg []byte) error {
	return w.writeMessage(ctx, msg)
}

func Test_MessageProducer_Publish_Success(t *testing.T) {
	messageWriter := &messageWriterMock{
		writeMessage: func(context.Context, []byte) error {
			return nil
		},
	}

	producer := produce.NewMessageProducer(messageWriter)

	message := model.EventMessage{
		Message: "test message",
	}

	errGot := producer.Publish(context.Background(), message)

	errExpected := error(nil)

	if errGot != errExpected {
		t.Fatalf("error does not match, expected: %v, got: %v", errExpected, errGot)
	}
}

func Test_MessageProducer_Publish_WriteFailed(t *testing.T) {
	errTest := errors.New("test error")
	messageWriter := &messageWriterMock{
		writeMessage: func(context.Context, []byte) error {
			return errTest
		},
	}

	producer := produce.NewMessageProducer(messageWriter)

	message := model.EventMessage{
		Message: "test message",
	}

	errGot := producer.Publish(context.Background(), message)

	if !errors.Is(errGot, errTest) {
		t.Fatalf("error does not match, expect error '%v', to be '%v'", errGot, errTest)
	}
}
