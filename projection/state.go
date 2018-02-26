package projection

import (
	"errors"

	"github.com/golang/protobuf/ptypes"
	"github.com/nicklanng/carpark/events"
)

// TODO: Need some sort of concurrency protection around the state

// State is the projection of the events into the state of the application.
type State struct {
	tickets map[TicketID]*Ticket
}

// NewState returns a new state object to keep track of all aggregates.
func NewState() *State {
	state := new(State)
	state.tickets = make(map[TicketID]*Ticket)
	return state
}

// ProcessEvent folds the supplied event in to the current aggregates.
func (s *State) ProcessEvent(event events.Event) error {
	switch v := event.(type) {
	case *events.TicketIssued:
		// turn timestamp back into go type
		issuedAt, err := ptypes.Timestamp(v.GetAt())
		if err != nil {
			return err
		}

		t := &Ticket{
			ID:       TicketID(v.GetTicketID()),
			IssuedAt: issuedAt,
		}
		if err := s.createTicket(t); err != nil {
			return err
		}
	}

	return nil
}

// GetTicket returns a ticket in the state.
// The boolean returns whether the ticket was found.
func (s *State) GetTicket(id TicketID) (*Ticket, bool) {
	ticket, ok := s.tickets[id]
	return ticket, ok
}

func (s *State) createTicket(t *Ticket) error {
	if _, ok := s.tickets[t.ID]; ok {
		return errors.New("Ticket already exists")
	}

	s.tickets[t.ID] = t

	return nil
}
