package main

import "net/http"

type httpServer struct {
	mux *http.ServeMux
}

func NewHTTPServer() *httpServer {
	serveMux := http.NewServeMux()
	return &httpServer{
		mux: serveMux,
	}
}
