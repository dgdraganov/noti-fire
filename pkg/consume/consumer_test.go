package consume_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dgdraganov/noti-fire/internal/model"
	"github.com/dgdraganov/noti-fire/pkg/consume"
)

type readCommitterMock struct {
	readMessage   func(ctx context.Context) (*model.ConsumedMessage, error)
	commitMessage func(ctx context.Context, msg *model.ConsumedMessage) error
}

func (rc *readCommitterMock) ReadMessage(ctx context.Context) (*model.ConsumedMessage, error) {
	return rc.readMessage(ctx)
}
func (rc *readCommitterMock) CommitMessage(ctx context.Context, msg *model.ConsumedMessage) error {
	return rc.commitMessage(ctx, msg)
}

func Test_MessageConsumer_Consume_Success(t *testing.T) {

	expextedMessage := &model.ConsumedMessage{}

	reader := &readCommitterMock{
		readMessage: func(ctx context.Context) (*model.ConsumedMessage, error) {
			return expextedMessage, nil
		},
	}

	consumer := consume.NewMessageConsumer(reader)

	gotMessage, err := consumer.Consume(context.Background())
	if err != nil {
		t.Fatalf("unexpecte error occured")
	}

	if gotMessage != expextedMessage {
		t.Fatalf("unexpected message consumed, expected: %p, got: %p", expextedMessage, gotMessage)
	}
}

func Test_MessageConsumer_Consume_Failed(t *testing.T) {

	errTest := errors.New("test error")
	var expectedMessage *model.ConsumedMessage

	reader := &readCommitterMock{
		readMessage: func(ctx context.Context) (*model.ConsumedMessage, error) {
			return nil, errTest
		},
	}

	consumer := consume.NewMessageConsumer(reader)
	gotMessage, errGot := consumer.Consume(context.Background())
	if errGot == nil {
		t.Fatal("unexpected nil error")
	}

	if !errors.Is(errGot, errTest) {
		t.Fatalf("error does not match, expected error %q, to be %q", errGot, errTest)

	}

	if expectedMessage != gotMessage {
		t.Fatalf("unexpected message, expected: %p, got: %p", expectedMessage, gotMessage)
	}
}

func Test_MessageConsumer_Commit_Success(t *testing.T) {

	reader := &readCommitterMock{
		commitMessage: func(ctx context.Context, msg *model.ConsumedMessage) error {
			return nil
		},
	}

	consumer := consume.NewMessageConsumer(reader)

	consumedMessage := &model.ConsumedMessage{}
	errGot := consumer.MarkConsumed(context.Background(), consumedMessage)

	errExpected := error(nil)
	if errGot != errExpected {
		t.Fatalf("error does not match, expected: %v, got: %v", errExpected, errGot)
	}
}

func Test_MessageConsumer_Commit_Failed(t *testing.T) {

	errTest := errors.New("test error")
	reader := &readCommitterMock{
		commitMessage: func(ctx context.Context, msg *model.ConsumedMessage) error {
			return errTest
		},
	}

	consumer := consume.NewMessageConsumer(reader)

	consumedMessage := &model.ConsumedMessage{}
	errGot := consumer.MarkConsumed(context.Background(), consumedMessage)

	if !errors.Is(errGot, errTest) {
		t.Fatalf("error does not match, expected error '%v', to be '%v'", errGot, errTest)
	}
}
