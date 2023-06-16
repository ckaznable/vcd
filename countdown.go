package main

import "time"

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

		d = tick(d)
		ch <- d
	}

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
