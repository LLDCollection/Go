package main

// Main function to test the functionality.
// This can be modified to either support CLI or API.
func main() {
	// Create a parking lot
	parkingLot := initialzeParking()

	// Print availability
	availabilityDisplay(parkingLot)

	// Parking
	simulateParkings(parkingLot)

	// Check availability after parkings
	availabilityDisplay(parkingLot)
}
