package main

import (
  "github.com/nsf/termbox-go"
  "bufio"
	"os"
)

type screen struct {
  buffer  [][]rune
  pos     int
}

func (s *screen) draw() {
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  _, height := termbox.Size()
  lines := s.buffer[s.pos:(height+s.pos)]

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
  for scanner.Scan(){
    s.buffer = append(s.buffer, []rune(scanner.Text()))
    if len(s.buffer) == height {
      s.draw()
    }
  }
}

func (s *screen) moveUp() {
  if s.pos > 0 {
    s.pos -= 1
    s.draw()
  }
}

func (s *screen) moveDown() {
  _, height := termbox.Size()
  if len(s.buffer) > height && s.pos < (len(s.buffer) - height) {
    s.pos += 1
    s.draw()
  }
}
