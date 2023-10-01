package notification

import "context"

type Action interface {
	Execute(ctx context.Context, message string) error
}
