package ctx_test

import (
	"net/http"
	"testing"

	"github.com/nicklanng/carpark/api/ctx"
)

func TestRequestIDFunctions(t *testing.T) {
	h := http.Header{}
	h.Add("x-request-id", "123456789")
	r := &http.Request{Header: h}

	contextWithRequestID := ctx.NewContextWithRequestID(r.Context(), r)
	correlationID := ctx.RequestIDFromContext(contextWithRequestID)

	if correlationID != "123456789" {
		t.Error("Expected 123456789, got ", correlationID)
	}
}
