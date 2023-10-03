package process_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dgdraganov/noti-fire/internal/model"
	"github.com/dgdraganov/noti-fire/internal/process"
)

type publisherMock struct {
	publish func(ctx context.Context, msg model.EventMessage) error
}

func (p *publisherMock) Publish(ctx context.Context, msg model.EventMessage) error {
	return p.publish(ctx, msg)
}

func Test_ProcessAction_Execute_Success(t *testing.T) {
	publisher := &publisherMock{
		publish: func(ctx context.Context, msg model.EventMessage) error {
			return nil
		},
	}

	action := process.NewProcessAction(publisher)
	errGot := action.Execute(context.Background(), "test message")

	errExpected := error(nil)

	if errExpected != errGot {
		t.Fatalf("error does not match, expected: %s, got: %s", errExpected, errGot)
	}

}

func Test_ProcessAction_Execute_Error(t *testing.T) {
	var errTest error = errors.New("test error")
	publisher := &publisherMock{
		publish: func(ctx context.Context, msg model.EventMessage) error {
			return errTest
		},
	}

	action := process.NewProcessAction(publisher)
	errGot := action.Execute(context.Background(), "test message")

	if !errors.Is(errGot, errTest) {
		t.Fatalf("error does not match, expected error %q, to be %q", errGot, errTest)
	}
}
