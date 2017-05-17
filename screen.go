package main

import (
  "github.com/nsf/termbox-go"
  "bufio"
	"os"
)

type screen struct {
  buffer        [][]rune
  scrollPos     int
  cursorPos     int
  view          chan []rune
}

func initScreen() screen {
  return screen{ scrollPos: 0, view: make(chan []rune)}
}

func (s *screen) draw() {
  width, height := termbox.Size()
  for {
    runes := <- s.view
    x := 0
    for _, char := range runes {
      if s.cursorPos < (height-1) {
        termbox.SetCell(x, s.cursorPos, char, termbox.ColorDefault, termbox.ColorDefault)
      }
      x += 1
      if x > width {
        s.cursorPos += 1
        x = 2
      }
    }
    s.cursorPos += 1
    termbox.Flush()
  }
}

func (s *screen) reset() {
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  s.cursorPos = 0
  _, height := termbox.Size()

  lines := s.buffer[s.scrollPos:(height+s.scrollPos)-1]

  for _, runes := range lines {
    s.view <- runes
  }
}

func (s *screen) load() {
  scanner := bufio.NewScanner(os.Stdin)
  _, height := termbox.Size()

  go func() {
    for scanner.Scan() {
      s.buffer = append(s.buffer, []rune(scanner.Text()))

      if s.cursorPos < (height-1) {
        s.view <- s.buffer[len(s.buffer)-1]
      }
    }
  }()
}

func (s *screen) moveUp() {
  if s.scrollPos > 0 {
    s.scrollPos -= 1
    s.reset()
  }
}

func (s *screen) moveDown() {
  _, height := termbox.Size()
  if len(s.buffer) > height && s.scrollPos < (len(s.buffer) - height) {
    s.scrollPos += 1
    s.reset()
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
			}
		}
	}
}
