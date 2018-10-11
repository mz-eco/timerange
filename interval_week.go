package timerange

import (
	"time"
)

type week int

func (m week) Current(o time.Time) time.Time {

	var (
		wd       = weekday(o)
		y, mm, d = o.Date()
	)

	return date(y, mm, d-wd)

}

func (m week) Now() TimeRange {
	return now(m)
}

func (m week) To(b time.Time, size int) TimeRange {
	return RangeTo(
		Truncate(b,m),
		week(size))
}

func (m week) At(now time.Time) TimeRange {
	return RangeAt(now,m)
}

func (m week) Range(b,e time.Time) TimeRange {
	return Range(
		Truncate(b,m),
		Truncate(e,m),
	)
}

func Weeks(size int) Interval {
	return week(size)
}

func (m week) Next(o time.Time) time.Time {
	return m.Current(o).AddDate(0, 0, 7)
}

func (m week) IsWhole(now time.Time) bool {
	return m.Current(now).Equal(now)
}

func (m week) Preview(o time.Time) time.Time {
	return m.Current(o).AddDate(0,0,-7)
}

func (m week) AddTo(o time.Time) time.Time {

	var (
		wd = weekday(o)
	)

	switch {
	case m == 0:
		return o
	case m == 1:
		return o.AddDate(0, 0, 7-wd)
	default:
		o = o.AddDate(0, 0, 7-wd)
		return o.AddDate(0, 0, int(m-1)*7)
	}

}

func (m week) Allow() Allow {

	switch {
	case m > 0:
		return AllowForward
	case m < 0:
		return AllowRevert
	default:
		return AllowStop
	}
}
