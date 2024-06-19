package ticket

import (
	"math/rand"
	"time"

	"parkinglot/internal/vehicle"
)

type Ticket struct {
	id        int64
	startTime time.Time
	endTime   time.Time
	vehicle   vehicle.Vehicle
}

func (t *Ticket) CalculatePrice() float64 {
	return 5 + (float64(t.endTime.Unix()-t.startTime.Unix()) * float64(0.5))
}

func (t *Ticket) ID() int64 {
	return t.id
}

func New(vehicle vehicle.Vehicle) *Ticket {
	return &Ticket{
		id:        100000 + rand.Int63n(99999),
		startTime: time.Now(),
		vehicle:   vehicle,
	}
}
