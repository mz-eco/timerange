package timerange

import "time"

type second time.Duration

func (m second) Date(year int, month time.Month, day int, hour int, minute int, second int) time.Time {
	return time.Date(
		year,
		month,
		day,
		hour,
		minute,
		second,
		0,
		time.Local)
}

func Seconds(size int) Interval {
	return second(size)
}

func (m second) To(b time.Time, size int) TimeRange {
	return RangeTo(
		Truncate(b, m),
		second(size))
}

func (m second) At(now time.Time) TimeRange {
	return RangeAt(now, m)
}

func (m second) Range(b, e time.Time) TimeRange {
	return Range(
		Truncate(b, m),
		Truncate(e, m),
	)
}

func (m second) GetValue() time.Duration {
	return time.Duration(m)
}

func (m second) GetUnit() time.Duration {
	return time.Second
}

func (m second) Next(o time.Time) time.Time {
	return next(o, m)
}

func (m second) Current(o time.Time) time.Time {
	return current(o, m)
}

func (m second) Preview(o time.Time) time.Time {
	return preview(o, m)
}

func (m second) IsWhole(now time.Time) bool {
	return m.Current(now).Equal(now)
}

func (m second) AddTo(o time.Time) time.Time {
	return o.Add(time.Duration(m) * time.Second)
}

func (m second) GetSize() (days int, duration time.Duration) {
	return 0, time.Duration(m) * time.Second
}

func (m second) Allow() Allow {

	switch {
	case m > 0:
		return AllowForward
	case m < 0:
		return AllowRevert
	default:
		return AllowStop
	}
}
