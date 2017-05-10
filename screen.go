package main

import (
  "github.com/nsf/termbox-go"
  "bufio"
	"os"
)

type Screen struct {
  buffer  [][]rune
  pos     int
}

func (s *Screen) Draw() {
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

func (s *Screen) Load() {
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    s.buffer = append(s.buffer, []rune(scanner.Text()))
  }
}

func (s *Screen) MoveUp() {
  if s.pos > 0 {
    s.pos -= 1
    s.Draw()
  }
}

func (s *Screen) MoveDown() {
  _, height := termbox.Size()
  if len(s.buffer) > height && s.pos < (len(s.buffer) - height) {
    s.pos += 1
    s.Draw()
  }
}
