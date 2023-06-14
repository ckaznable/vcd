package main

import (
	"fmt"
	"time"

	"github.com/ckaznable/vcd/nv"
)

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

	cmp(now, nextVacation)
}

func cmp(time1 time.Time, time2 time.Time) {
	duration := time2.Sub(time1)

	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	fmt.Printf("%d:%d:%d\n", days, hours, minutes)
}
