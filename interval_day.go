package timerange

import "time"

type day int


func (m day) Date(year int, month time.Month, day int) time.Time {
	return date(year,month,day)
}

func Days(size int) day {
	return day(size)
}


func (m day) Preview(o time.Time) time.Time {

	var (
		y, mm, d = o.Date()
	)

	return date(y, mm, d-1)
}

func (m day) Next(o time.Time) time.Time {
	var (
		y, mm, d = o.Date()
	)

	return date(y, mm, d+1)
}

func (m day) Current(o time.Time) time.Time {
	var (
		y, mm, d = o.Date()
	)

	return date(y, mm, d)
}

func (m day) IsWhole(now time.Time) bool {
	return m.Current(now).Equal(now)
}

func (m day) Add(o time.Time) time.Time {
	return o.AddDate(0, 0, int(m))
}

func (m day) GetSize() (days int, duration time.Duration) {
	return int(m), 0
}

func (m day) Allow() Allow {

	switch {
	case m > 0:
		return AllowForward
	case m < 0:
		return AllowRevert
	default:
		return AllowStop
	}
}
