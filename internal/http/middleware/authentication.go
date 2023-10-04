package middleware

import (
	"net/http"

	"github.com/dgdraganov/noti-fire/internal/model"
	"go.uber.org/zap"
)

type authenticatorMiddleware struct {
	logs *zap.SugaredLogger
}

func NewAuthenticatorMiddleware(logger *zap.SugaredLogger) *authenticatorMiddleware {
	return &authenticatorMiddleware{
		logs: logger,
	}
}

// Auth implements the middleware logic to authenticate an incoming request
func (auth *authenticatorMiddleware) Auth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Context().Value(model.RequestID).(string)

		// Here goes the logic that will veryfy whether the
		// call is valid in terms of authentication. For
		// example this can be achieved by using JWT authentication.

		auth.logs.Infow(
			"successfully authenticated",
			"request_id", requestID,
		)
		handler.ServeHTTP(w, r)
	})
}
