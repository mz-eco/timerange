package timerange

import "time"

type today int

func (m today) Day() time.Time {
	return Truncate(time.Now(), Day)
}

func (m today) Month() time.Time {
	return Truncate(time.Now(), Month)
}

func (m today) Year() time.Time {
	return Truncate(time.Now(), Year)
}

func (m today) Begin() time.Time {
	return Begin(time.Now(), Day)
}

func (m today) End() time.Time {
	return End(time.Now(), Day)
}

func (m today) Time(hour, minute, second int) time.Time {
	var (
		now      = time.Now()
		y, mm, d = now.Date()
		loc      = now.Location()
	)

	return time.Date(
		y, mm, d,
		hour, minute, second,
		0,
		loc,
	)
}

func (m today) Now() time.Time {
	return time.Now()
}

const (
	Today today = 1
)
