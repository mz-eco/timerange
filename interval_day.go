package timerange

import "time"

type day int

func (m day) Date(year int, month time.Month, day int) time.Time {
	return date(year, month, day)
}

func (m day) Now() TimeRange {
	return now(m)
}

func (m day) Today() TimeRange {
	return now(m)
}

func (m day) To(b time.Time, size int) TimeRange {
	return RangeTo(
		Truncate(b, m),
		day(size))
}

func (m day) At(now time.Time) TimeRange {
	return RangeAt(now, m)
}

func (m day) Range(b, e time.Time) TimeRange {
	return Range(
		Truncate(b, m),
		Truncate(e, m),
	)
}

func (m day) Add(now time.Time, size int) time.Time {
	return day(size).AddTo(now)
}

func Days(size int) Interval {
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

func (m day) AddTo(o time.Time) time.Time {
	return o.AddDate(0, 0, int(m))
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
