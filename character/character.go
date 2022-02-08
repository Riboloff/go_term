package character

const (
	FRONT = iota + 1
	RIGHT_ONE
	RIGHT_TWO
)

type Char struct {
	Head, Body, Legs [3]byte
	Cord             Cord
	Height           int
	State            int
}

type Cord struct {
	X, Y int
}

func NewChar() Char {
	return Char{
		Head:   [3]byte{' ', '0', ' '},
		Body:   [3]byte{'<', '#', '\\'},
		Legs:   [3]byte{'/', ' ', '\\'},
		Cord:   Cord{1, 1},
		State:  FRONT,
		Height: 3}
}

func (c *Char) ArmUp() {
	c.Head[2] = '/'
	c.Body[2] = ' '
}

func (c *Char) ArmDown() {
	c.Head[2] = ' '
	c.Body[2] = '\\'
}

func (c *Char) MoveRight() {
	c.Cord.X = c.Cord.X + 1

	if c.State != RIGHT_ONE {
		c.State = RIGHT_ONE
		c.Body[0] = '.'
		c.Body[2] = '\\'
		c.Legs[0] = '/'
		c.Legs[1] = '>'
		c.Legs[2] = ' '
	} else {
		c.State = RIGHT_TWO
		c.Body[0] = '<'
		c.Body[2] = '.'
		c.Legs[0] = '-'
		c.Legs[1] = '|'
		c.Legs[2] = ' '
	}
}
