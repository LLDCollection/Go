package vehicle

type VehicleType string

type Vehicle interface {
	LicensePlateNumber() string
	Type() VehicleType
}
