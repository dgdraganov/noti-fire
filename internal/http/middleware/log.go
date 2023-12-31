package middleware

import (
	"net/http"

	"github.com/dgdraganov/noti-fire/internal/model"
	"go.uber.org/zap"
)

type loggerMiddleware struct {
	logs *zap.SugaredLogger
}

func NewLoggerMiddleware(logger *zap.SugaredLogger) *loggerMiddleware {
	return &loggerMiddleware{
		logs: logger,
	}
}

// Log implements the middleware logic to log incoming request and outgoing responses
func (logger *loggerMiddleware) Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Context().Value(model.RequestID).(string)

		logger.logs.Infow(
			"request received",
			"request_method", r.Method,
			"response_id", requestID,
			"request_url", r.URL,
		)
		handler.ServeHTTP(w, r)

		logger.logs.Infow(
			"server response",
			"response_id", requestID,
		)
	})
}
