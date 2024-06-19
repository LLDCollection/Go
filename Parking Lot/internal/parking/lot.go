package parking

import (
	"parkinglot/internal/ticket"
	"parkinglot/internal/vehicle"
)

type Lot struct {
	availableSpots map[Size][]*Spot
	parkingMap     map[*ticket.Ticket]*Spot
}

func (l *Lot) Availability(s Size) int {
	var available int
	for _, s := range l.availableSpots[s] {
		if !s.IsOccupied() {
			available += 1
		}
	}

	return available
}

func (l *Lot) Park(vehicle vehicle.Vehicle) *ticket.Ticket {
	spot := l.findSpot(vehicle)
	if spot == nil {
		return nil
	}
	spot.Occupy()

	tk := ticket.New(vehicle)
	l.parkingMap[tk] = spot

	return tk
}

func (l *Lot) findSpot(vehicle vehicle.Vehicle) *Spot {
	sizes := getPotentialSizes(vehicle)
	for _, s := range sizes {
		spot := l.findSpotForSize(s)
		if spot != nil {
			return spot
		}
	}

	return nil
}

func (l *Lot) findSpotForSize(size Size) *Spot {
	for _, s := range l.availableSpots[size] {
		if !s.IsOccupied() {
			return s
		}
	}

	return nil
}

func getPotentialSizes(v vehicle.Vehicle) []Size {
	switch v.Type() {
	case vehicle.BIKE:
		return []Size{S, M, L, XL}
	case vehicle.CAR:
		return []Size{M, L, XL}
	case vehicle.BUS:
		return []Size{L, XL}
	default:
		return []Size{XL}
	}
}

func NewLot(spotSizes map[Size]int) Lot {
	spotMap := make(map[Size][]*Spot)

	for s, c := range spotSizes {
		for range c {
			spotMap[s] = append(spotMap[s], NewSpot(s))
		}
	}

	return Lot{availableSpots: spotMap, parkingMap: make(map[*ticket.Ticket]*Spot)}
}
