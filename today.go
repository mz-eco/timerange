package timerange

import "time"

type today int


func (m today) Day() TimeRange {
	return now(Day)
}

func (m today) Month() TimeRange {
	return now(Month)
}

func (m today) Year() TimeRange {
	return now(Year)
}

func (m today) Begin() time.Time {
	return Begin(time.Now(), Day)
}

func (m today) End() time.Time {
	return End(time.Now(), Day)
}

func (m today) Time(hour,minute,second int) time.Time {
	var (
		now = time.Now()
		y,mm,d = now.Date()
		loc = now.Location()
	)

	return time.Date(
		y,mm,d,
		hour,minute,second,
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
