package parking

type Size int

const (
	S Size = iota
	M
	L
	XL
)

type Spot struct {
	isOccupied bool
	size       Size
}

func (s *Spot) IsOccupied() bool {
	return s.isOccupied
}

func (s *Spot) Occupy() {
	s.isOccupied = true
}

func (s *Spot) Vacate() {
	s.isOccupied = false
}

func NewSpot(size Size) *Spot {
	return &Spot{
		isOccupied: false,
		size:       size,
	}
}
