package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

var tickFlag = true

func draw(s tcell.Screen, d Date) {
	joinStr := ":"
	if tickFlag {
		tickFlag = false
	} else {
		joinStr = " "
		tickFlag = true
	}

	str := fmt.Sprintf("%d%s%0*d%s%0*d", d.Day, joinStr, 2, d.Hours, joinStr, 2, d.Minutes)
	if d.IsDefault() {
		str = "Have good vacation!"
	}

	width, height := s.Size()
	y := height / 2
	x := width/2 - len(str)/2

	s.Clear()
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	for i, char := range str {
		s.SetContent(x+i, y, char, nil, defStyle)
	}

	s.Show()
}
