package vehicle

const BIKE VehicleType = "bike"

type Bike struct {
	licensePlateNumber string
}

func (b *Bike) LicensePlateNumber() string {
	return b.licensePlateNumber
}

func (b *Bike) Type() VehicleType {
	return BIKE
}

func NewBike(licensePlateNumber string) *Bike {
	return &Bike{licensePlateNumber: licensePlateNumber}
}
