package produce

import "context"

type MessageWriter interface {
	WriteMessage(context.Context, []byte) error
}
