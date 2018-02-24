package events

import (
	"database/sql"
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/lib/pq"
)

const (
	// PqUniqueViolation is a Postgres error code
	PqUniqueViolation pq.ErrorCode = "23505"
)

var (
	// ErrSequenceOutOfOrder means that the event the caller tried to insert is been superceded by another
	ErrSequenceOutOfOrder = errors.New("sequence out of order")
)

// insertEvent will attempt to append an event to the end of the log.
// Seq should be the last seq + 1. This allows the caller to check that another event which is incompatible with another that has been inserted before.
func insertEvent(db *sql.DB, seq int64, event Event) error {
	eventType := EventType(event)

	data, err := proto.Marshal(event)
	if err != nil {
		return err
	}

	if _, err := db.Exec("INSERT INTO events (seq, type, data) VALUES ($1,$2,$3)", seq, eventType, data); err != nil {
		pqErr := err.(*pq.Error)
		switch pqErr.Code {
		case PqUniqueViolation:
			return ErrSequenceOutOfOrder
		default:
			return err
		}
	}

	return nil
}
