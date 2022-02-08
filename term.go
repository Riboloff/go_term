package main

import (
	"time"

	"term/character"
	"term/gamemap"
	"term/printer"
)

func main() {
	printer.Clear()
	char := character.NewChar()
	gamemap := gamemap.NewMap()

	charDancing(gamemap, char, 3)
	charMoveRight(gamemap, char)
}

func charMoveRight(m gamemap.Map, char character.Char) {
	for {
		char.MoveRight()
		printer.PrintMap(m, char)
		time.Sleep(time.Second)
	}
}

func charDancing(m gamemap.Map, char character.Char, count int) {
	for i := 0; i < count; i++ {
		printer.PrintMap(m, char)
		if i%2 == 0 {
			char.ArmUp()
		} else {
			char.ArmDown()
		}
		time.Sleep(time.Second)
	}
}
