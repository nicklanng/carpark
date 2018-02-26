package projection

import (
	"errors"

	"github.com/golang/protobuf/ptypes"
	"github.com/nicklanng/carpark/events"
)

// TODO: Need some sort of concurrency protection around the state

type State struct {
	tickets map[TicketID]*Ticket
}

func NewState() *State {
	state := new(State)
	state.tickets = make(map[TicketID]*Ticket)
	return state
}

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
