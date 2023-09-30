package server

import "net/http"

type httpServer struct {
	mux *http.ServeMux
}

// NewHTTPServer is a constructor function for the httpServer type
func NewHTTPServer() *httpServer {
	serveMux := http.NewServeMux()
	return &httpServer{
		mux: serveMux,
	}
}
