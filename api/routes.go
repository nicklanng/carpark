package api

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/nicklanng/carpark/api/handlers"
)

// BuildRoutes returns a mux router with all API routes
func BuildRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Health check
	r.Handle("/_status", handlers.NotImplemented()).Methods("GET")

	return r
}
