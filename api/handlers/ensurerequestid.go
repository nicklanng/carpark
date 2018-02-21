package handlers

import (
	"net/http"

	"github.com/nicklanng/carpark/api/ctx"
	"github.com/satori/go.uuid"
)

// EnsureRequestID will create a Request ID header if one is not found.
//
// The EnsureRequestID handler is added automatically if you use the api.StartHTTPSServer function.
func EnsureRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("x-request-id") == "" {
			requestID := uuid.Must(uuid.NewV4())
			r.Header.Add("x-request-id", requestID.String())
		}

		ctxWithReqID := ctx.NewContextWithRequestID(r.Context(), r)

		next.ServeHTTP(w, r.WithContext(ctxWithReqID))
	})
}
