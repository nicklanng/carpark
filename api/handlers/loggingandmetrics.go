package handlers

import (
	"net/http"

	"github.com/nicklanng/carpark/logging"
	"github.com/nicklanng/carpark/metrics"
)

// LoggingAndMetrics is an HTTP handler that will log all HTTP requests and responses using the logging library
// and increment request and response counts in statsd.
// A good practice is to wrap your top-level mux in this handler.
//
// The LoggingAndMetrics handler is added automatically if you use the api.StartHTTPSServer function.
func LoggingAndMetrics(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.LogHTTPRequest(r)
		metrics.CountHTTPRequest(r)

		res := newResponseCatcher(w, r)
		h.ServeHTTP(res, r)

		metrics.CountHTTPResponse(&res.response)
		logging.LogHTTPResponse(&res.response)
	})
}

func newResponseCatcher(w http.ResponseWriter, r *http.Request) *responseCatcher {
	res := &responseCatcher{}
	res.w = w
	res.response.Request = r
	res.response.Proto = r.Proto
	return res
}

type responseCatcher struct {
	w        http.ResponseWriter
	response http.Response
}

func (rc *responseCatcher) Header() http.Header {
	return rc.w.Header()
}

func (rc *responseCatcher) Write(b []byte) (int, error) {
	size, err := rc.w.Write(b)
	rc.response.ContentLength += int64(size)
	return size, err
}

func (rc *responseCatcher) WriteHeader(s int) {
	rc.w.WriteHeader(s)
	rc.response.StatusCode = s
}

func (rc *responseCatcher) Flush() {
	f, ok := rc.w.(http.Flusher)
	if ok {
		f.Flush()
	}
}
