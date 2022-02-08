package main

import (
	"time"

	tm "github.com/buger/goterm"
)

type char struct {
	head, body, legs [3]byte
	cord             cord
	height           int
}

type cord struct {
	x, y int
}

type _map struct {
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
	m.head[2] = '/'
	m.body[2] = ' '
}

func charArmDown(m *char) {
	m.head[2] = ' '
	m.body[2] = '\\'
}

func charInit() char {
	return char{
		head:   [3]byte{' ', '0', ' '},
		body:   [3]byte{'<', '#', '\\'},
		legs:   [3]byte{'/', ' ', '\\'},
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

	flag := 0
	for i := 0; i < len(m.grid); i++ {
		for j := 0; j < len(m.grid[0]); j++ {
			if cord.x == i && cord.y == j {
				putChar(i, j, m, char)
				flag = 1
				break
			}
		}
		if flag == 1 {
			break
		}
	}
}

func putChar(x int, y int, m *_map, char char) {

	for i := 0; i < char.height; i++ {
		for j := 0; j < len(char.head); j++ {
			if i == 0 {
				m.grid[x+j][y+i] = char.head[j]
			}
			if i == 1 {
				m.grid[x+j][y+i] = char.body[j]
			}
			if i == 2 {
				m.grid[x+j][y+i] = char.legs[j]
			}
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
	//printMap(_map, char)

	//charDance(char)
	charMoveRight(_map, char)
}

func charMoveRight(m _map, char char) {
	i := 0
	for {
		char.cord.x = char.cord.x + 1

		printMap(m, char)
		if i%2 == 1 {
			charArmUp(&char)
		} else {
			charArmDown(&char)
		}
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
