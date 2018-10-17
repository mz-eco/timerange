package timerange

import "time"

type year int

func (m year) Date(year int) time.Time {
	return date(year, time.January, 1)
}

func Years(size int) Interval {
	return year(size)
}

func (m year) Now() TimeRange {
	return now(m)
}

func (m year) To(b time.Time, size int) TimeRange {
	return RangeTo(
		Truncate(b, m),
		year(size))
}

func (m year) At(now time.Time) TimeRange {
	return RangeAt(now, m)
}

func (m year) Range(b, e time.Time) TimeRange {
	return Range(
		Truncate(b, m),
		Truncate(e, m),
	)
}

func (m year) Next(o time.Time) time.Time {
	var (
		y, _, _ = o.Date()
	)

	return date(y+1, time.January, 1)
}

func (m year) Current(o time.Time) time.Time {
	var (
		y, _, _ = o.Date()
	)

	return date(y, time.January, 1)
}

func (m year) Preview(o time.Time) time.Time {
	var (
		y, _, _ = o.Date()
	)

	return date(y-1, time.January, 1)

}
func (m year) IsWhole(now time.Time) bool {
	return m.Current(now).Equal(now)
}

func (m year) AddTo(o time.Time) time.Time {
	return o.AddDate(int(m), 0, 0)
}

func (m year) Allow() Allow {

	switch {
	case m > 0:
		return AllowForward
	case m < 0:
		return AllowRevert
	default:
		return AllowStop
	}
}
