package elevator

import (
	"container/heap"
	pq "elevator/pkg/priorityqueue"
	"fmt"
)

type Direction int

const (
	Idle Direction = 0
	Up   Direction = 1
	Down Direction = 2
)

type Elevator struct {
	currentFloor int
	direction    Direction
	upQueue      pq.MinIntHeap
	downQueue    pq.MinIntHeap
}

func New() *Elevator {
	e := &Elevator{
		currentFloor: 0,
		direction:    Idle,
		upQueue:      pq.MinIntHeap{},
		downQueue:    pq.MinIntHeap{},
	}

	heap.Init(&e.upQueue)
	heap.Init(&e.downQueue)

	return e
}

func (e *Elevator) RequestFloor(floor int) {
	if floor > e.currentFloor {
		heap.Push(&e.upQueue, floor)
		if e.direction == Idle {
			e.direction = Up
		}
	} else if floor < e.currentFloor {
		heap.Push(&e.downQueue, floor)
		if e.direction == Idle {
			e.direction = Down
		}
	}
}

func (e *Elevator) Move() {
	if e.direction == Idle {
		fmt.Println("Elevator is Idle")
	}

	var destination int
	if e.direction == Up && e.upQueue.Len() > 0 {
		destination = heap.Pop(&e.upQueue).(int)
	} else if e.direction == Down && e.downQueue.Len() > 0 {
		destination = heap.Pop(&e.downQueue).(int)
	}

	fmt.Printf("Moving from Floor %v to Floor %v\n", e.currentFloor, destination)
	e.currentFloor = destination
	e.arrival()
}

func (e *Elevator) CurrentDirenction() Direction {
	return e.direction
}

func (e *Elevator) arrival() {
	fmt.Printf("Arrived at Floor %v\n", e.currentFloor)
	
	if e.direction == Up && e.upQueue.Len() == 0 && e.downQueue.Len() > 0 {
		e.direction = Down
	} else if e.direction == Down && e.downQueue.Len() == 0 && e.upQueue.Len() > 0 {
		e.direction = Up
	} else if e.upQueue.Len() == 0 && e.downQueue.Len() == 0 {
		e.direction = Idle
	}
}
