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

	for {
		charDancing(gamemap, &char, 3)
		charMoveRight(gamemap, &char, 5)
		charMoveLeft(gamemap, &char, 5)
	}
}

func charMoveRight(m gamemap.Map, char *character.Char, count int) {
	for i := 0; i < count; i++ {
		char.MoveRight()
		printer.PrintMap(m, *char)
		time.Sleep(time.Second)
	}
}

func charMoveLeft(m gamemap.Map, char *character.Char, count int) {
	for i := 0; i < count; i++ {
		char.MoveLeft()
		printer.PrintMap(m, *char)
		time.Sleep(time.Second)
	}
}

func charDancing(m gamemap.Map, char *character.Char, count int) {
	for i := 0; i < count; i++ {
		printer.PrintMap(m, *char)
		if i%2 == 0 {
			char.ArmUp()
		} else {
			char.ArmDown()
		}
		time.Sleep(time.Second)
	}
}
