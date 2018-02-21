package handlers

import "net/http"

// NotImplemented simply returns a 501 Not Implemented
func NotImplemented() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	})
}
