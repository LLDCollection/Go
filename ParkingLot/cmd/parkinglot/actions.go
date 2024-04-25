// This file contains some helper functions to generate and print data.
package main

import (
	"fmt"

	"parkinglot/internal/parking"
	"parkinglot/internal/vehicle"
)

func initialzeParking() parking.Lot {
	sizes := map[parking.Size]int{
		parking.S:  2,
		parking.M:  2,
		parking.L:  1,
		parking.XL: 1,
	}
	parkingLot := parking.NewLot(sizes)
	return parkingLot
}

func availabilityDisplay(parkingLot parking.Lot) {
	fmt.Printf("Availability of Small Spots: %v\n", parkingLot.Availability(parking.S))
	fmt.Printf("Availability of Medium Spots: %v\n", parkingLot.Availability(parking.M))
	fmt.Printf("Availability of Large Spots: %v\n", parkingLot.Availability(parking.L))
	fmt.Printf("Availability of Xtra Large Spots: %v\n", parkingLot.Availability(parking.XL))
}

func simulateParkings(parkingLot parking.Lot) {
	vehicles := []vehicle.Vehicle{
		vehicle.NewBike("10BH24-2024"),
		vehicle.NewBike("10BH24-2025"),
		vehicle.NewCar("11BH24-2024"),
		vehicle.NewBike("10BH24-2026"),
		vehicle.NewCar("11BH24-2025"),
		vehicle.NewCar("11BH24-2026"),
		vehicle.NewCar("11BH24-2027"),
	}

	for _, v := range vehicles {
		t := parkingLot.Park(v)
		if t != nil {
			fmt.Printf("Ticket No. %v assigned to %v\n", t.ID(), v.LicensePlateNumber())
		} else {
			fmt.Println("No parking spot available.")
		}
	}
}
