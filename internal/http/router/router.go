package router

import "net/http"

type notificationRouter struct {
	mux *http.ServeMux
}

// NewNotificationRouter is a constructor function for the notificationRouter type
func NewNotificationRouter() *notificationRouter {
	serveMux := http.NewServeMux()
	return &notificationRouter{
		mux: serveMux,
	}
}

// Register registers the given handler to the underlying serve mux
func (router *notificationRouter) Register(pattern string, handler http.Handler) {
	router.mux.Handle(pattern, handler)
}

// ServeMux is used to return the underlying serve mux
func (router *notificationRouter) ServeMux() *http.ServeMux {
	return router.mux
}
