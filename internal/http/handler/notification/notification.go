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

func NewNotificationHandler(method string, action Action, logger *zap.SugaredLogger) *notificationHandler {
	return &notificationHandler{
		action: action,
		logs:   logger,
		method: method,
	}
}

// ServeHTTP implements the http.Handler interface for the notificationHandler type
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
		err := common.WriteResponse(
			w,
			fmt.Sprintf("invalid request method %s, expected method is %s", r.Method, n.method),
			http.StatusBadRequest,
		)
		if err != nil {
			n.logs.Errorw(
				"write response",
				"request_id", requestID,
				"error", err,
			)
		}
		return
	}
	notification := model.NotificationRequest{}
	// decode json request body
	err := common.JSONDecode(r.Body, &notification)
	if err != nil {
		err := common.WriteResponse(
			w,
			"invalid JSON request body",
			http.StatusBadRequest,
		)
		if err != nil {
			n.logs.Errorw(
				"write response",
				"request_id", requestID,
				"error", err,
			)
		}
		return
	}

	// perform handler action
	err = n.action.Execute(r.Context(), notification.Message)
	if err != nil {
		n.logs.Errorw(
			"execute action",
			"request_id", requestID,
			"error", err.Error(),
		)
		err := common.WriteResponse(
			w,
			"Something went wrong on our end!",
			http.StatusInternalServerError)
		if err != nil {
			n.logs.Errorw(
				"write response",
				"request_id", requestID,
				"error", err,
			)
		}
		return
	}

	// OK
	err = common.WriteResponse(
		w,
		"notification request submitted successfully",
		http.StatusOK,
	)
	if err != nil {
		n.logs.Errorw(
			"write response",
			"request_id", requestID,
			"error", err,
		)
	}
}
