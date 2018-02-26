package projection

import (
	"errors"
	"math"
	"time"
)

type TicketID string

type Ticket struct {
	ID       TicketID
	IssuedAt time.Time
}

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
