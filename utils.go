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
	return current(o,d).Add(-1*d.GetUnit())
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

func clock(h, m, s int) time.Time {
	return time.Date(
		0, time.July, 1, h, m, s, 0, time.Local)
}

func size(b, e time.Time) (days int, duration time.Duration) {

	days = dayAt(e.Date()) - dayAt(b.Date())

	var (
		cb = clock(b.Clock())
		ce = clock(e.Clock())
	)

	duration = ce.Sub(cb)

	return
}

func weekday(now time.Time) int {

	switch now.Weekday() {
	case time.Monday:
		return 0
	case time.Tuesday:
		return 1
	case time.Wednesday:
		return 2
	case time.Thursday:
		return 3
	case time.Friday:
		return 4
	case time.Saturday:
		return 5
	case time.Sunday:
		return 6
	}

	return 0
}

func now(w Whole) TimeRange {
	var (
		now = time.Now()
	)

	return Range(
		w.Current(now),
		w.Next(now))
}
