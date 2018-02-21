package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nicklanng/carpark/api/ctx"
	"github.com/nicklanng/carpark/api/handlers"
)

func TestEnsureRequestID_CreateNewID(t *testing.T) {
	// create a dummy request with no requestID
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// run the request through the handler
	rr := httptest.NewRecorder()
	handler := handlers.EnsureRequestID(pingRequestIDFromContext())
	handler.ServeHTTP(rr, req)

	// check for new request ID
	var newRequestID string
	if newRequestID = req.Header.Get("x-request-id"); newRequestID == "" {
		t.Error("Request ID has not been set")
	}

	// requestID should have been added to the request Context
	if requestID := rr.Body.String(); requestID != newRequestID {
		t.Error("Request ID has not been put in the request context")
	}
}

func TestEnsureRequestID_LeaveExistingRequestID(t *testing.T) {
	originalRequestID := "this-value-should-remain"

	// create a dummy request with the original requestID
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("x-request-id", originalRequestID)

	// run the request through the handler
	rr := httptest.NewRecorder()
	handler := handlers.EnsureRequestID(pingRequestIDFromContext())
	handler.ServeHTTP(rr, req)

	// check for new request ID
	if requestID := req.Header.Get("x-request-id"); requestID != originalRequestID {
		t.Errorf("RequestID should have been untouched but was set to %s", requestID)
	}

	// requestID should have been added to the request Context
	if requestID := rr.Body.String(); requestID != originalRequestID {
		t.Error("Request ID has not been put in the request context")
	}
}

// gets the request ID from the request context and echoes it back in the response body
func pingRequestIDFromContext() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ctx.RequestIDFromContext(r.Context())))
	})
}
