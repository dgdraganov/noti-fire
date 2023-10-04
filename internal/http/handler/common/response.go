package common

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgdraganov/noti-fire/internal/model"
)

// WriteResponse is used by http handlers in order to write the specific message and status code to the response writer object
func WriteResponse(w http.ResponseWriter, message string, statusCode int) error {
	respMsg := model.ResponseMessage{Message: message}
	resp, err := json.Marshal(respMsg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Something went wrong on our end!")); err != nil {
			return fmt.Errorf("response write: %w", err)
		}
		return fmt.Errorf("json marshal: %w", err)
	}
	w.WriteHeader(statusCode)
	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("response write: %w", err)
	}
	return nil
}
