package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ckaznable/vcd/nv"
	"github.com/gdamore/tcell/v2"
)

type Date struct {
	Day     int
	Hours   int
	Minutes int
}

func main() {
	now := time.Now()

	builder := nv.NewVacationBuilder()
	builder.AllWeekendsInYear(now.Year())
	builder.AllWeekendsInYear(now.Year() + 1)
	builder.ExcludeDays([]string{
		"2023-06-17",
		"2023-09-23",
	})

	builder.Days([]string{
		"2023-06-22",
		"2023-06-23",
		"2023-09-29",
		"2023-10-09",
		"2023-10-10",
	})

	list := builder.Build()
	nextVacation := time.Now()

	for _, d := range list {
		if d.After(now) {
			nextVacation = d
			break
		}
	}

	nextDate := diff(now, nextVacation)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	// Clear screen
	s.Clear()

	ch := make(chan Date)

	quit := func() {
		s.Fini()
		os.Exit(0)
	}

	go func() {
		first := true
		now := time.Now()
		remainingSeconds := 60 - now.Second()

		if now.Second() == 0 {
			first = false
		}

		ch <- nextDate
		for {
			if first {
				first = false
				time.Sleep(time.Duration(remainingSeconds) * time.Second)
			} else {
				time.Sleep(time.Minute * 1)
			}

			nextDate = tick(nextDate)
			ch <- nextDate
		}
	}()

	go func() {
		for {
			// Poll event
			ev := s.PollEvent()

			// Process event
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					quit()
				}
			}
		}
	}()

	for {
		select {
		case nextDate = <-ch:
		default:
		}

		draw(s, nextDate)
		time.Sleep(time.Second)
	}
}

func diff(time1 time.Time, time2 time.Time) Date {
	duration := time2.Sub(time1)

	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	return Date{
		Day:     days,
		Hours:   hours,
		Minutes: minutes,
	}
}

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

func tick(d Date) Date {
	d.Minutes--
	if d.Minutes == -1 {
		d.Minutes = 59
		d.Hours--
	}

	if d.Hours == -1 {
		d.Hours = 23
		d.Day--
	}

	if d.Day == -1 {
		d.Day = 0
	}

	return d
}
