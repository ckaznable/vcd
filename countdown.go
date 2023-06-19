package main

import (
	"errors"
	"time"

	"github.com/ckaznable/vcd/nv"
)

func countdown(d Date, ch chan Date) {
	first := true
	now := time.Now()
	remainingSeconds := 60 - now.Second()

	if now.Second() == 0 {
		first = false
	}

	ch <- d
	for {
		if first {
			first = false
			time.Sleep(time.Duration(remainingSeconds) * time.Second)
		} else {
			time.Sleep(time.Minute * 1)
		}

		d, err := tick(d)
		if err != nil {
			ch <- d
			time.Sleep(time.Hour * 24)
			d = getCountDownDate()
		}

		ch <- d
	}
}

func tick(d Date) (Date, error) {
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
		return DefaultDate(), errors.New("time to fun")
	}

	return d, nil
}

func getCountDownDate() Date {
	nv := getNextVacation()
	workTime, err := parseWorkTimeArgs()
	if err == nil {
		nv = time.Date(nv.Year(), nv.Month(), nv.Day()-1, workTime.Work.Hours, workTime.Work.Minutes, 0, 0, nv.Location())
	}

	return getDiffDate(time.Now(), nv)
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
