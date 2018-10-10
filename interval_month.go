package monent

import "time"

type month int

func (m month) Date(year int, month time.Month) time.Time {
	return date(year,month,1)
}

func (m month) Next(o time.Time) time.Time {
	var (
		y, mm, _ = o.Date()
	)

	return date(y, mm+1, 1)
}

func (m month) Current(o time.Time) time.Time {
	var (
		y, mm, _ = o.Date()
	)

	return date(y, mm, 1)
}

func (m month) Preview(o time.Time) time.Time {
	var (
		y, mm, _ = o.Date()
	)

	return date(y, mm-1, 1)
}

func (m month) IsWhole(now time.Time) bool {
	return m.Current(now).Equal(now)
}

func (m month) Add(o time.Time) time.Time {
	return o.AddDate(0, int(m), 0)
}

func (m month) Allow() Allow {

	switch {
	case m > 0:
		return AllowForward
	case m < 0:
		return AllowRevert
	default:
		return AllowStop
	}
}
