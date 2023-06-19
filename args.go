package main

import (
	"errors"
	"os"
	"strconv"
)

type WorkTime struct {
	Work Date
	Off  Date
}

func DefaultWorkTime() *WorkTime {
	return &WorkTime{
		Work: DefaultDate(),
		Off:  DefaultDate(),
	}
}

func (w *WorkTime) IsDefault() bool {
	return w.Work.IsDefault() && w.Off.IsDefault()
}

func parseWorkTimeArgs() (*WorkTime, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("no arguments")
	}

	workTime := os.Args[1][:4]
	getOffTime := os.Args[1][4:]

	work, err := parseDate(workTime)
	if err != nil {
		return nil, err
	}

	off, err := parseDate(getOffTime)
	if err != nil {
		return nil, err
	}

	return &WorkTime{
		Work: work,
		Off:  off,
	}, nil
}

func parseDate(str string) (Date, error) {
	if len(str) < 4 {
		return Date{}, errors.New("invalid date")
	}

	hour, err := strconv.Atoi(str[:2])
	if err != nil {
		return Date{}, err
	}

	min, err := strconv.Atoi(str[2:])
	if err != nil {
		return Date{}, err
	}

	return Date{
		Day:     0,
		Hours:   hour,
		Minutes: min,
	}, nil
}
