package shared

type Position struct {
	x int
	y int
}

func NewPosition(x int, y int) *Position {
	return &Position{
		x: x,
		y: y,
	}
}

func (p *Position) X() int {
	return p.x
}

func (p *Position) Y() int {
	return p.y
}
