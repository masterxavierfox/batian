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

	s := Screen{ pos: 0 }

	s.Load()
	s.Draw()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowUp:
				s.MoveUp()
			case termbox.KeyArrowDown:
				s.MoveDown()
			default:
				s.Draw()
			}
		}
	}
}
