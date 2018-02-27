package projection

import (
	"errors"
	"math"
	"time"

	"github.com/nicklanng/carpark/events"
)

// TicketStatus is an enum for tracking state of a ticket.
// A stringer is generated for this type.
type TicketStatus int

//go:generate stringer -type=TicketStatus

const (
	// TicketStatusIssued is the initial state of a new ticket
	TicketStatusIssued TicketStatus = iota
	// TicketStatusPaid means that the ticket has been paid for
	TicketStatusPaid
)

// TicketID is a type alias of a string
type TicketID string

// Ticket is an issued ticket for a car in a carpark
type Ticket struct {
	ID       TicketID
	Status   TicketStatus
	IssuedAt time.Time
}

// IsValidTransition determines whether the event could be applied, disregarding ID
func (t *Ticket) IsValidTransition(event events.Event) bool {
	switch event.(type) {
	case *events.TicketPaid:
		return t.Status == TicketStatusIssued
	case *events.TicketComplete:
		return t.Status == TicketStatusPaid
	}

	return false
}

// GetTariff gets the current price of a ticket
func (t *Ticket) GetTariff(now time.Time) (int, error) {
	elapsed := now.Sub(t.IssuedAt)

	if elapsed < 0 {
		return 0, errors.New("Invalid time")
	}

	if elapsed < 1*time.Hour {
		return 150, nil
	}

	if elapsed < 3*time.Hour {
		return 300, nil
	}

	if elapsed < 6*time.Hour {
		return 1000, nil
	}

	if elapsed < 24*time.Hour {
		return 2000, nil
	}

	numberOfDaysParked := int(math.Ceil(elapsed.Hours() / 24))
	return numberOfDaysParked * 2500, nil
}
