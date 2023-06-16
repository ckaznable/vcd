package main

import (
	"time"

	"github.com/ckaznable/vcd/nv"
)

func main() {
	now := time.Now()
	nextDate := getDiffDate(now, getNextVacation())

	screen := NewScreen()
	ch := make(chan Date)

	go handleKey(screen)
	go countdown(nextDate, ch)
	Run(screen, nextDate, ch)
}

func getDiffDate(time1 time.Time, time2 time.Time) Date {
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

func getNextVacation() time.Time {
	now := time.Now()
	builder := nv.NewVacationBuilder()
	builder.AllWeekendsInYear(now.Year())
	builder.AllWeekendsInYear(now.Year() + 1)
	builder.ExcludeDays(nv.GetWorkDay())
	builder.Days(nv.GetHoliday())

	list := builder.Build()
	nextVacation := time.Now()

	for _, d := range list {
		if d.After(now) {
			nextVacation = d
			break
		}
	}

	return nextVacation
}
