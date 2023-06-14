package nv

import (
	"sort"
	"time"
)

type VacationBuilder struct {
	DayList        []time.Time
	ExcludeDayList []time.Time
}

func NewVacationBuilder() *VacationBuilder {
	return &VacationBuilder{
		DayList:        make([]time.Time, 0),
		ExcludeDayList: make([]time.Time, 0),
	}
}

func (v *VacationBuilder) Day(date string) error {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}

	v.DayList = append(v.DayList, d)
	return nil
}

func (v *VacationBuilder) Days(date []string) map[string]error {
	errorMap := make(map[string]error)
	for _, d := range date {
		err := v.Day(d)
		if err != nil {
			errorMap[d] = err
		}
	}

	return errorMap
}

func (v *VacationBuilder) ExcludeDay(date string) error {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}

	v.ExcludeDayList = append(v.ExcludeDayList, d)
	return nil
}

func (v *VacationBuilder) ExcludeDays(date []string) map[string]error {
	errorMap := make(map[string]error)
	for _, d := range date {
		err := v.ExcludeDay(d)
		if err != nil {
			errorMap[d] = err
		}
	}

	return errorMap
}

func (v *VacationBuilder) AllWeekendsInYear(year int) *VacationBuilder {
	startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC)

	for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
			v.DayList = append(v.DayList, d)
		}
	}

	return v
}

func (v *VacationBuilder) AllWeekends() *VacationBuilder {
	now := time.Now()
	year := now.Year()
	return v.AllWeekendsInYear(year)
}

func (v *VacationBuilder) removeDuplicates(times []time.Time) []time.Time {
	encountered := map[time.Time]bool{}
	uniqueTimes := []time.Time{}

	for _, t := range times {
		if !encountered[t] {
			encountered[t] = true
			uniqueTimes = append(uniqueTimes, t)
		}
	}

	return uniqueTimes
}

func (v *VacationBuilder) filterTimes(times1, times2 []time.Time) []time.Time {
	filtered := []time.Time{}

	// 遍歷第一個時間切片
	for _, t1 := range times1 {
		exists := false

		// 檢查是否存在於第二個時間切片中
		for _, t2 := range times2 {
			if t1.Equal(t2) {
				exists = true
				break
			}
		}

		// 如果不存在於第二個時間切片中，則添加到過濾後的切片中
		if !exists {
			filtered = append(filtered, t1)
		}
	}

	return filtered
}

func (v *VacationBuilder) Build() []time.Time {
	days := v.removeDuplicates(v.DayList)
	excludeDays := v.removeDuplicates(v.ExcludeDayList)
	filtered := v.filterTimes(days, excludeDays)

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Before(filtered[j])
	})

	return filtered
}
