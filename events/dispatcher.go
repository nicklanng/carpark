package events

import (
	"database/sql"
	"fmt"

	"github.com/nicklanng/carpark/logging"
)

func NewDispatcher(db *sql.DB) chan<- Event {
	eventChan := make(chan Event, 10)

	go func() {
		for {
			event := <-eventChan

			if err := insertEvent(db, event); err != nil {
				logging.Error(fmt.Sprintf("Failed to insert event: %s", err.Error()))
			}
		}
	}()

	return eventChan
}
