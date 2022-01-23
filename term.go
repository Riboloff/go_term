package main

import (
	"time"

	tm "github.com/buger/goterm"
)

type char struct {
	head, body, legs string
	cord             cord
	height           int
}

type cord struct {
	x, y int
}

type _map struct {
	//grid [2][8]byte
	grid [100][4]byte
}

func printChar(char char) {
	tm.Print(char.head)
	tm.MoveCursorDown(1)
	tm.MoveCursorBackward(3)
	tm.Print(char.body)
	tm.MoveCursorDown(1)
	tm.MoveCursorBackward(3)
	tm.Print(char.legs)
	tm.Flush()
}

func charArmUp(m *char) {
	m.head = " 0/"
	m.body = "<# "
	m.legs = "/ \\"
}

func charArmDown(m *char) {
	m.head = " 0 "
	m.body = "<#\\"
	m.legs = "/ \\"
}

func charInit() char {
	return char{
		head:   " 0 ",
		body:   "<#\\",
		legs:   "/ \\",
		cord:   cord{1, 1},
		height: 3}
}

func mapInit() _map {
	var g [100][4]byte
	for i := 1; i < len(g); i++ {
		for j := 1; j < len(g[0]); j++ {
			g[i][j] = ' '
		}
	}
	return _map{
		grid: g}
}

func putCharOnMap(m *_map, char char) {
	cord := char.cord

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
		}
	}
}

func printMap(m _map, char char) {
	putCharOnMap(&m, char)
	g := m.grid
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
	char := charInit()
	_map := mapInit()
	printMap(_map, char)
	//charDance(char)
	//charMoveRight(char)
}

func charMoveRight(char char) {
	i := 0
	for {
		char.cord.x = char.cord.x + 3
		tm.MoveCursor(char.cord.x, char.cord.y)

		printChar(char)
		//if i%2 == 1 {
		//} else {
		//}
		time.Sleep(time.Second)
		i = i + 1
	}
}

func charDance(char char) {
	i := 0
	for {
		tm.MoveCursor(char.cord.x, char.cord.y)
		if i%2 == 1 {
			charArmUp(&char)
		} else {
			charArmDown(&char)
		}
		printChar(char)
		time.Sleep(time.Second)
		i = i + 1
	}
}
