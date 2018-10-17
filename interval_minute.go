package timerange

import "time"

type minute time.Duration

func (m minute) Date(year int, month time.Month, day int, hour int, minute int) time.Time {
	return time.Date(
		year,
		month,
		day,
		hour,
		minute,
		0,
		0,
		time.Local)
}

func Minutes(size int) Interval {
	return minute(size)
}

func (m minute) Now() TimeRange {
	return now(m)
}

func (m minute) To(b time.Time, size int) TimeRange {
	return RangeTo(
		Truncate(b, m),
		minute(size))
}

func (m minute) At(now time.Time) TimeRange {
	return RangeAt(now, m)
}

func (m minute) Range(b, e time.Time) TimeRange {
	return Range(
		Truncate(b, m),
		Truncate(e, m),
	)
}

func (m minute) GetValue() time.Duration {
	return time.Duration(m)
}

func (m minute) GetUnit() time.Duration {
	return time.Minute
}

func (m minute) Next(o time.Time) time.Time {
	return next(o, m)
}

func (m minute) Current(o time.Time) time.Time {
	return current(o, m)
}

func (m minute) Preview(o time.Time) time.Time {
	return preview(o, m)
}

func (m minute) IsWhole(now time.Time) bool {
	return m.Current(now).Equal(now)
}

func (m minute) AddTo(o time.Time) time.Time {
	return o.Add(time.Duration(m) * time.Minute)
}

func (m minute) GetSize() (days int, duration time.Duration) {
	return 0, time.Duration(m) * time.Minute
}

func (m minute) Allow() Allow {

	switch {
	case m > 0:
		return AllowForward
	case m < 0:
		return AllowRevert
	default:
		return AllowStop
	}
}
