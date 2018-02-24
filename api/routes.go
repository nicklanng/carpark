package api

import (
	"github.com/gorilla/mux"
	"github.com/nicklanng/carpark/api/commands"
	"github.com/nicklanng/carpark/api/handlers"
	"github.com/nicklanng/carpark/events"
)

// BuildRoutes returns a mux router with all API routes
func BuildRoutes(eventChan chan<- events.Event) *mux.Router {
	r := mux.NewRouter()

	// Health check
	r.Handle("/_status", handlers.NotImplemented()).Methods("GET")

	// commands
	r.Handle("/ticket", commands.CreateTicket(eventChan)).Methods("POST")

	return r
}
