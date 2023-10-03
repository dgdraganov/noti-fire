package server

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type httpServer struct {
	logs   *zap.SugaredLogger
	server *http.Server
}

// NewHTTPServer is a constructor function for the httpServer type
func NewHTTPServer(port string, serveMux *http.ServeMux, logger *zap.SugaredLogger) *httpServer {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: serveMux,
	}

	return &httpServer{
		logs:   logger,
		server: server,
	}
}

// Start runs an http server on the specified port
func (s *httpServer) Start(port string) {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			s.logs.Errorw(
				"server failed unexpectedly",
				"error", err,
			)
		}
	}()
}

// Shutdown stops the server gracefully
func (s *httpServer) Shutdown() {

	if err := s.server.Shutdown(context.Background()); err != nil {
		s.logs.Errorw(
			"server failed unexpectedly",
			"error", err,
		)
	}
}
