package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	s := screen{ pos: 0 }

	s.load()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowUp:
				s.moveUp()
			case termbox.KeyArrowDown:
				s.moveDown()
			default:
				s.draw()
			}
		}
	}
}
