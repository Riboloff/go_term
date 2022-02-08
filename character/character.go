package character

type Char struct {
	Head, Body, Legs [3]byte
	Cord             Cord
	Height           int
}

type Cord struct {
	X, Y int
}

func (c *Char) ArmUp() {
	c.Head[2] = '/'
	c.Body[2] = ' '
}

func (c *Char) ArmDown() {
	c.Head[2] = ' '
	c.Body[2] = '\\'
}

func NewChar() Char {
	return Char{
		Head:   [3]byte{' ', '0', ' '},
		Body:   [3]byte{'<', '#', '\\'},
		Legs:   [3]byte{'/', ' ', '\\'},
		Cord:   Cord{1, 1},
		Height: 3}
}
