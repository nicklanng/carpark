package events

import (
	"database/sql"
	"fmt"

	"github.com/nicklanng/carpark/logging"
)

func NewDispatcher(db *sql.DB) chan<- Event {
	eventChan := make(chan Event, 10)

	go func() {
		seq := int64(0)

		for {
			event := <-eventChan

			if err := insertEvent(db, seq, event); err != nil {
				logging.Error(fmt.Sprintf("Failed to insert event: %s", err.Error()))
			}

			seq = seq + 1
		}
	}()

	return eventChan
}
