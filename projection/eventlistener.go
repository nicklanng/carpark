package projection

import (
	"encoding/hex"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/lib/pq"
	"github.com/nicklanng/carpark/events"
	"github.com/nicklanng/carpark/logging"
)

func CreateEventListener(state *State, eventListener *pq.Listener) {
	go func() {
		for {
			// listen to postgres notify for new event
			notification := <-eventListener.NotificationChannel()

			// parse the notification string
			notificationFields := strings.Split(notification.Extra, ",")
			eventType := notificationFields[1]
			eventHex := notificationFields[2]
			eventBytes, err := hex.DecodeString(eventHex)
			if err != nil {
				logging.Error("Error decoding hex event: " + err.Error())
				return
			}

			// TODO: check the sequence is not out of order

			// unmarshall event
			switch eventType {
			case "TicketIssued":
				event := &events.TicketIssued{}
				if err := proto.Unmarshal(eventBytes, event); err != nil {
					logging.Error("Error unmarshalling bytes to event: " + err.Error())
					return
				}
				state.ProcessEvent(event)
			default:
				logging.Warn("Unknown error type")
				return
			}

		}
	}()
}
