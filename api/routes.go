package api

import (
	"github.com/gorilla/mux"
	"github.com/nicklanng/carpark/api/commands"
	"github.com/nicklanng/carpark/api/handlers"
	"github.com/nicklanng/carpark/api/queries"
	"github.com/nicklanng/carpark/events"
	"github.com/nicklanng/carpark/projection"
)

// BuildRoutes returns a mux router with all API routes
func BuildRoutes(state *projection.State, eventChan chan<- events.Event) *mux.Router {
	r := mux.NewRouter()

	// Health check
	r.Handle("/_status", handlers.NotImplemented()).Methods("GET")

	// commands
	r.Handle("/ticket", commands.CreateTicket(eventChan)).Methods("POST")
	r.Handle("/ticket/{ID}/pay", commands.PayForTicket(eventChan, state)).Methods("POST")
	r.Handle("/ticket/{ID}/complete", commands.CompleteTicket(eventChan, state)).Methods("POST")

	// queries
	r.Handle("/ticket/{ID}/tariff", queries.GetTicketTariff(state)).Methods("GET")

	return r
}
