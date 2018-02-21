package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nicklanng/carpark/api/handlers"
)

func TestNotImplemented(t *testing.T) {
	// create a dummy request
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// run the request through the handler
	rr := httptest.NewRecorder()
	handler := handlers.NotImplemented()
	handler.ServeHTTP(rr, req)

	// code should be NOT IMPLEMENTEd
	if rr.Code != http.StatusNotImplemented {
		t.Error("Not implemented handler returned wrong code")
	}
}
