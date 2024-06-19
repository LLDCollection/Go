package main

import (
	"elevator/internal/elevator"
)

func main() {
	myElevator := elevator.New()
	
	myElevator.RequestFloor(5)
	myElevator.RequestFloor(8)
	myElevator.RequestFloor(3)
	myElevator.RequestFloor(7)

	for myElevator.CurrentDirenction() != elevator.Idle {
		myElevator.Move()
	}
}
