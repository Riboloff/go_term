package main

import (
	"time"

	tm "github.com/buger/goterm"
)

type man struct {
	head, body, legs string
}

func printMan(m man) {
	tm.Print(m.head)
	tm.MoveCursorDown(1)
	tm.MoveCursorBackward(3)
	tm.Print(m.body)
	tm.MoveCursorDown(1)
	tm.MoveCursorBackward(3)
	tm.Print(m.legs)
	tm.Flush()
}

func upArm(m *man) {
	m.head = " 0/"
	m.body = "<# "
}

func downArm(m *man) {
	m.head = " 0 "
	m.body = "<#\\"
}

func main() {
	tm.Clear()
	m := man{" 0 ", "<#\\", "/ \\"}
	tm.MoveCursor(1, 1)
	i := 0
	for {
		tm.MoveCursor(1, 1)
		printMan(m)
		if i%2 == 1 {
			upArm(&m)
		} else {
			downArm(&m)
		}
		time.Sleep(time.Second)
		i = i + 1
	}
}
