package vehicle

const BUS VehicleType = "bus"

type Bus struct {
	licensePlateNumber string
}

func (b *Bus) LicensePlateNumber() string {
	return b.licensePlateNumber
}

func (b *Bus) Type() VehicleType {
	return BIKE
}

func NewBus(licensePlateNumber string) *Bus {
	return &Bus{licensePlateNumber: licensePlateNumber}
}
