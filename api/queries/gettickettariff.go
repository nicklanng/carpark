package queries

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nicklanng/carpark/projection"
)

type getTicketTariffResponse struct {
	ID     projection.TicketID `json:"id"`
	Tariff int                 `json:"tariff"`
}

// GetTicketTariff returns the current price of a parking ticket.
func GetTicketTariff(state *projection.State) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idParam := vars["ID"]
		ticketID := projection.TicketID(idParam)

		ticket, ok := state.GetTicket(ticketID)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		tariff, err := ticket.GetTariff(time.Now())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		response := getTicketTariffResponse{
			ID:     ticket.ID,
			Tariff: tariff,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})
}
