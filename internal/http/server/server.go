package server

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type httpServer struct {
	mux *http.ServeMux
}

// NewHTTPServer is a constructor function for the httpServer type
func NewHTTPServer(serveMux *http.ServeMux, logger *zap.SugaredLogger) *httpServer {
	return &httpServer{
		mux: serveMux,
	}
}

// StartServer runs an http server on the specified port
func (s *httpServer) Start(port string) error {
	p := fmt.Sprintf(":%s", port)
	return http.ListenAndServe(p, s.mux)
}
