package main

import (
	"time"

	"term/character"
	"term/gamemap"

	tm "github.com/buger/goterm"
)

func putCharOnMap(m *gamemap.Map, char character.Char) {
	x := char.Cord.X
	y := char.Cord.Y

	for i := 0; i < char.Height; i++ {
		for j := 0; j < len(char.Head); j++ {
			switch i {
			case 0:
				m.Grid[x+j][y+i] = char.Head[j]
			case 1:
				m.Grid[x+j][y+i] = char.Body[j]
			case 2:
				m.Grid[x+j][y+i] = char.Legs[j]
			}
		}
	}
}

func printMap(m gamemap.Map, char character.Char) {
	putCharOnMap(&m, char)
	g := m.Grid
	for i := 1; i < len(g); i++ {
		for j := 1; j < len(g[0]); j++ {
			tm.MoveCursor(i, j)
			tm.Printf("%c", g[i][j])
		}
	}
	tm.Flush()
}

func main() {
	tm.Clear()
	char := character.NewChar()
	_map := gamemap.NewMap()
	//printMap(_map, char)

	charMoveRight(_map, char)
}

func charMoveRight(m gamemap.Map, char character.Char) {
	i := 0
	for {
		char.Cord.X = char.Cord.X + 1

		printMap(m, char)
		if i%2 == 1 {
			char.ArmUp()
		} else {
			char.ArmDown()
		}
		time.Sleep(time.Second)
		i = i + 1
	}
}
