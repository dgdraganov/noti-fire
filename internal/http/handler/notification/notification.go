package notification

import "net/http"

type notificationHandler struct {
	operation Operation
}

func (n *notificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// not implemented
}

// NewNotificationHandler is a constructor function for the notificationHandler type
func NewNotificationHandler(operation Operation) *notificationHandler {
	return &notificationHandler{
		operation: operation,
	}
}
