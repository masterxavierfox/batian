package main

import (
  "github.com/nsf/termbox-go"
  "bufio"
	"os"
)

type screen struct {
  buffer        [][]rune
  scrollPos     int
}

func initScreen() screen {
  return screen{ scrollPos: 0 }
}

func (s *screen) draw() {
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  _, height := termbox.Size()
  lines := s.buffer[s.scrollPos:(height+s.scrollPos)]

  for y, runes := range lines {
   for x, char := range runes {
     termbox.SetCell(x, y, char, termbox.ColorDefault, termbox.ColorDefault)
   }
  }

  termbox.Flush()
}

func (s *screen) load() {
  scanner := bufio.NewScanner(os.Stdin)
  _, height := termbox.Size()

  go func() {
    rendered := false
    for scanner.Scan(){
      s.buffer = append(s.buffer, []rune(scanner.Text()))

      if len(s.buffer) == height {
        s.draw()
        rendered = true
      }
    }
    if !rendered {
      s.draw()
    }
  }()
}

func (s *screen) moveUp() {
  if s.scrollPos > 0 {
    s.scrollPos -= 1
    s.draw()
  }
}

func (s *screen) moveDown() {
  _, height := termbox.Size()
  if len(s.buffer) > height && s.scrollPos < (len(s.buffer) - height) {
    s.scrollPos += 1
    s.draw()
  }
}

func (s *screen) loop() {
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
