package notification

import (
	"fmt"
	"net/http"

	"github.com/dgdraganov/noti-fire/internal/http/handler/common"
	"github.com/dgdraganov/noti-fire/internal/model"
	"go.uber.org/zap"
)

type notificationHandler struct {
	method string
	action Action
	logs   *zap.SugaredLogger
}

// NewNotificationHandler is a constructor function for the notificationHandler type
func NewNotificationHandler(method string, action Action, logger *zap.SugaredLogger) *notificationHandler {
	return &notificationHandler{
		action: action,
		logs:   logger,
	}
}
func (n *notificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value(model.RequestID).(string)
	// validate request method
	if r.Method != n.method {
		n.logs.Warnf(
			"invalid request method for notification handler",
			"request_method", r.Method,
			"expected_request_method", n.method,
			"request_id", requestID,
		)
		common.WriteResponse(
			w,
			fmt.Sprintf("invalid request method %q, expected method is %q", r.Method, n.method),
			http.StatusBadRequest,
		)
		return
	}

	// decode json request body
	notificReq := model.NotificationRequest{}
	err := common.JSONDecode(r.Body, notificReq)
	if err != nil {
		n.logs.Errorw(
			"request decode",
			"request_id", requestID,
			"error", err.Error(),
		)
		common.WriteResponse(
			w,
			"Something went wrong on our end!",
			http.StatusInternalServerError,
		)
		return
	}

	// perform handler action
	err = n.action.Execute(r.Context(), notificReq.Message)
	if err != nil {
		n.logs.Errorw(
			"execute action",
			"request_id", requestID,
			"error", err.Error(),
		)
		common.WriteResponse(
			w,
			"Something went wrong on our end!",
			http.StatusInternalServerError)
		return
	}

	// OK
	common.WriteResponse(
		w,
		"notification request submitted successfully",
		http.StatusOK,
	)
}
