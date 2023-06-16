package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func Run(s tcell.Screen, d Date, ch chan Date) {
	for {
		select {
		case d = <-ch:
		default:
		}

		draw(s, d)
		time.Sleep(time.Second)
	}
}

func handleKey(s tcell.Screen) {
	for {
		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit(s)
			}
		}
	}
}

func quit(s tcell.Screen) {
	s.Fini()
	os.Exit(0)
}

func NewScreen() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}

	s.SetStyle(tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset))
	s.Clear()

	return s
}
