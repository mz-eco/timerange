package timerange

import (
	"time"
)

func current(o time.Time, d durationUnitWhole) time.Time {
	return o.Truncate(d.GetUnit())
}

func next(o time.Time, d durationUnitWhole) time.Time {
	return current(o, d).Add(d.GetUnit())
}

func preview(o time.Time, d durationUnitWhole) time.Time {
	return current(o, d).Add(-1 * d.GetUnit())
}

func date(y int, m time.Month, d int) time.Time {
	return time.Date(
		y, m, d,
		0, 0, 0, 0,
		time.Local,
	)
}

func dayAt(y int, m time.Month, d int) int {
	m = (m + 9) % 12
	y = y - int(m)/10
	return 365*y + y/4 - y/100 + y/400 + (int(m)*306+5)/10 + (d - 1)
}

func weekday(now time.Time) int {

	var (
		wd = int(now.Weekday()) - 1
	)

	if wd >= 0 {
		return wd
	}

	return 6
}

func now(w Whole) TimeRange {
	var (
		now = time.Now()
	)

	return Range(
		w.Current(now),
		w.Next(now))
}
