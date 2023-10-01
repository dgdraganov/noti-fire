package middleware

import (
	"context"
	"net/http"

	"github.com/dgdraganov/noti-fire/internal/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type requestIdMiddleware struct {
	logs *zap.SugaredLogger
}

func NewRequestIdMiddleware(logger *zap.SugaredLogger) *requestIdMiddleware {
	return &requestIdMiddleware{
		logs: logger,
	}
}

func (r *requestIdMiddleware) Id(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New()
		ctx := context.WithValue(r.Context(), model.RequestID, requestID.String())
		r = r.WithContext(ctx)

		handler.ServeHTTP(w, r)
	})
}
