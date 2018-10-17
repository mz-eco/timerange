package timerange

import (
	"time"
)

const (
	Second second = 1
	Minute minute = 1
	Hour   hour   = 1
	Day    day    = 1
	Month  month  = 1
	Year   year   = 1
	Week   week   = 1
)

func Add(now time.Time, ivs ...Interval) time.Time {

	for _, iv := range ivs {
		now = iv.AddTo(now)
	}

	return now
}

func Preview(now time.Time, w Whole) time.Time {
	return w.Preview(now)
}

func Next(now time.Time, w Whole) time.Time {
	return w.Next(now)
}

func Truncate(now time.Time, w Whole) time.Time {
	return w.Current(now)
}

func Begin(now time.Time, w Whole) time.Time {
	return Truncate(now, w)
}

func End(now time.Time, w Whole) time.Time {
	return Next(now, w).Add(-1 * time.Nanosecond)
}
