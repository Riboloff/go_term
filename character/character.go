package character

type Char struct {
	Head, Body, Legs [3]byte
	Cord             Cord
	Height           int
}

type Cord struct {
	X, Y int
}

func CharArmUp(m *Char) {
	m.Head[2] = '/'
	m.Body[2] = ' '
}

func CharArmDown(m *Char) {
	m.Head[2] = ' '
	m.Body[2] = '\\'
}

func CharInit() Char {
	return Char{
		Head:   [3]byte{' ', '0', ' '},
		Body:   [3]byte{'<', '#', '\\'},
		Legs:   [3]byte{'/', ' ', '\\'},
		Cord:   Cord{1, 1},
		Height: 3}
}
