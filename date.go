package main

type Date struct {
	Day     int
	Hours   int
	Minutes int
}

func (d *Date) IsDefault() bool {
	return d.Day == 0 && d.Hours == 0 && d.Minutes == 0
}

func DefaultDate() Date {
	return Date{
		Day:     0,
		Hours:   0,
		Minutes: 0,
	}
}
