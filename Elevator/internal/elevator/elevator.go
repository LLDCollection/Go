package elevator

import (
	"fmt"
	"sort"
)

type Direction int

const (
	Up   Direction = 1
	Down Direction = 2
)

type Elevator struct {
	currentFloor int
	direction    Direction
	destinations []int
}

func (e *Elevator) RequestFloor(floor int) {
	for _, dest := range e.destinations {
		if dest == floor {
			return
		}
	}

	e.destinations = append(e.destinations, floor)
	sort.Ints(e.destinations)
}

func (e *Elevator) Move() {
	if len(e.destinations) == 0 {
		return
	}

	nextDestination := e.destinations[0]
	for _, dest := range e.destinations {
		if e.direction == Up && e.currentFloor < dest {
			nextDestination = dest
			break
		} else if e.direction == Down && e.currentFloor > dest {
			nextDestination = dest
			break
		}
	}

	fmt.Printf("Moving from Floor %v to Floor %v\n", e.currentFloor, nextDestination)
	e.currentFloor = nextDestination
	e.arrival()
}

func (e *Elevator) arrival() {
	fmt.Printf("Arrived at Floor %v", e.currentFloor)
	e.destinations = remove(e.destinations, e.currentFloor)
}

func remove(list []int, item int) []int {
	idx := -1

	for i, v := range list {
		if v == item {
			idx = i
			break
		}
	}

	if idx == -1 {
		return list
	}

	return append(list[:idx], list[idx + 1:]...)
}
