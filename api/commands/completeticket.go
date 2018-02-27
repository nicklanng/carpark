package commands

import (
	"net/http"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/gorilla/mux"
	"github.com/nicklanng/carpark/events"
	"github.com/nicklanng/carpark/projection"
)

// CompleteTicket as an HTTP handler that sends a TicketComplete to the event dispatcher.
func CompleteTicket(eventChan chan<- events.Event, state *projection.State) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idParam := vars["ID"]
		ticketID := projection.TicketID(idParam)

		now := time.Now()
		protoTimestamp, err := ptypes.TimestampProto(now)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ticket, ok := state.GetTicket(ticketID)
		if !ok {
			// error codes would be useful here
			w.WriteHeader(http.StatusNotFound)
			return
		}

		requestedEvent := &events.TicketComplete{
			At:       protoTimestamp,
			TicketID: string(ticketID),
		}

		if valid := ticket.IsValidTransition(requestedEvent); !valid {
			// error codes would be useful here
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// put the event on the channel to be persisted
		eventChan <- requestedEvent

		w.WriteHeader(http.StatusOK)
	})
}
