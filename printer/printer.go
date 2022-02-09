package printer

import (
	"term/character"
	"term/gamemap"

	tm "github.com/buger/goterm"
)

func Clear() {
	tm.Clear()
}

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

func PrintMap(m gamemap.Map, char character.Char) {
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
