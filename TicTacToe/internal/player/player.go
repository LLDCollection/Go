package player

type Player struct {
	symbol     rune
	score      int
	isComputer bool
}

func New(symbol rune) *Player {
	return &Player{
		symbol: symbol,
		score:  0,
	}
}

func (p *Player) Symbol() rune {
	return p.symbol
}
