package main

import (
	"bufio"
	"os"
	"github.com/nsf/termbox-go"
)

type screen struct {
	buffer 	[][]rune
	pos			int

}


func (s *screen) height() int {
	_, height := termbox.Size()
	if len(s.buffer) < height {
		return len(s.buffer)
	}
	return height
}

func (s *screen) width() int {
	width, _ := termbox.Size()
	return width
}

func (s *screen) draw() {
	lines := s.buffer[s.pos:(s.height()+s.pos)]
	y := 0
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	
	for _, runes := range lines {
		for x := 0; x < len(runes); x += 1 {
			termbox.SetCell(x, y, runes[x], termbox.ColorDefault, termbox.ColorDefault)
		}
		termbox.Flush()
		y += 1
	}
}

func (s *screen) load() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s.buffer = append(s.buffer, []rune(scanner.Text()))
	}
}

func (s *screen) moveUp() {
	if s.pos > 0 {
		s.pos -= 1
		s.draw()
	} else {
		return
	}
}

func (s *screen) moveDown() {
	if len(s.buffer) > s.height() && s.pos < (len(s.buffer) - s.height()) {
		s.pos += 1
		s.draw()
	}
}

func main() {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	s := screen{ pos: 0 }

	s.load()
	s.draw()

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
