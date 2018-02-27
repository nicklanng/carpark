package commands

import (
	"net/http"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/gorilla/mux"
	"github.com/nicklanng/carpark/events"
	"github.com/nicklanng/carpark/projection"
)

// PayForTicket as an HTTP handler that sends a TicketPaid to the event dispatcher.
func PayForTicket(eventChan chan<- events.Event, state *projection.State) http.Handler {
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

		requestedEvent := &events.TicketPaid{
			At:       protoTimestamp,
			TicketID: string(ticketID),
		}

		if valid := ticket.IsValidTransition(requestedEvent); !valid {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// put the event on the channel to be persisted
		eventChan <- requestedEvent

		w.WriteHeader(http.StatusOK)
	})
}
