package commands

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/nicklanng/carpark/events"
	uuid "github.com/satori/go.uuid"
)

type createTicketResponse struct {
	ID       string `json:"id"`
	IssuedAt string `json:"issuedAt"`
}

// CreateTicket as an HTTP handler that sends a TicketIssued to the event dispatcher.
func CreateTicket(eventChan chan<- events.Event) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ticketID := uuid.NewV4().String()

		// get time the ticket was issued at
		now := time.Now()
		protoTimestamp, err := ptypes.TimestampProto(now)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// put the event on the channel to be persisted
		eventChan <- &events.TicketIssued{
			At:       protoTimestamp,
			TicketID: ticketID,
		}

		// return the ticketID and issued at so that the machine cant print a matching ticket
		response := createTicketResponse{
			ID:       ticketID,
			IssuedAt: now.Format(time.RFC3339),
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})
}
