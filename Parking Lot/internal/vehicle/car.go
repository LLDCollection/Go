package vehicle

const CAR VehicleType = "car"

type Car struct {
	licensePlateNumber string
}

func (b *Car) LicensePlateNumber() string {
	return b.licensePlateNumber
}

func (b *Car) Type() VehicleType {
	return CAR
}

func NewCar(licensePlateNumber string) *Car {
	return &Car{licensePlateNumber: licensePlateNumber}
}
