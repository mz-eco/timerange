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
	return preview(o,m)
}

func (m hour) IsWhole(now time.Time) bool {
	return m.Current(now).Equal(now)
}

func (m hour) Add(o time.Time) time.Time {
	return o.Add(time.Duration(m) * time.Hour)
}

func (m hour) GetSize() (days int, duration time.Duration) {
	return 0, time.Duration(m) * time.Hour
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
