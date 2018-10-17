package timerange

import "time"

type hour time.Duration

func (m hour) Date(year int, month time.Month, day int, hour int) time.Time {
	return time.Date(
		year,
		month,
		day,
		hour,
		0,
		0,
		0,
		time.Local)
}

func Hours(size int) Interval {
	return hour(size)
}

func (m hour) To(b time.Time, size int) TimeRange {
	return RangeTo(
		Truncate(b, m),
		hour(size))
}

func (m hour) At(now time.Time) TimeRange {
	return RangeAt(now, m)
}

func (m hour) Range(b, e time.Time) TimeRange {
	return Range(
		Truncate(b, m),
		Truncate(e, m),
	)
}

func (m hour) Now() TimeRange {
	return now(m)
}

func (m hour) GetValue() time.Duration {
	return time.Duration(m)
}

func (m hour) GetUnit() time.Duration {
	return time.Hour
}

func (m hour) Next(o time.Time) time.Time {
	return next(o, m)
}

func (m hour) Current(o time.Time) time.Time {
	return current(o, m)
}

func (m hour) Preview(o time.Time) time.Time {
	return preview(o, m)
}

func (m hour) IsWhole(now time.Time) bool {
	return m.Current(now).Equal(now)
}

func (m hour) AddTo(o time.Time) time.Time {
	return o.Add(time.Duration(m) * time.Hour)
}

func (m hour) Allow() Allow {

	switch {
	case m > 0:
		return AllowForward
	case m < 0:
		return AllowRevert
	default:
		return AllowStop
	}
}
