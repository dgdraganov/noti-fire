package common

import (
	"encoding/json"
	"fmt"
	"io"
)

func JSONDecode(body io.ReadCloser, obj any) error {
	if err := json.NewDecoder(body).Decode(obj); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}
	return nil
}
