package timerange

import (
	"fmt"
	"time"
)

type Fixed struct {
	days     int
	duration time.Duration
}

func (m *Fixed) AddTo(now time.Time) time.Time {
	return now.Add(m.duration).AddDate(0, 0, m.days)
}

func (m *Fixed) Allow() Allow {

	switch {
	case m.days == 0:
		switch {
		case m.duration == 0:
			return AllowStop
		case m.duration > 0:
			return AllowForward
		case m.duration < 0:
			return AllowRevert
		default:
			return AllowStop
		}
	case m.days > 0:
		return AllowForward
	case m.days < 0:
		return AllowRevert
	default:
		return AllowStop
	}
}

func (m *Fixed) GetSize() (days int, duration time.Duration) {
	return m.days, m.duration
}

func (m *Fixed) String() string {
	return fmt.Sprintf("[%d day %s)", m.days, m.duration)
}

func NewInterval(intervals ...FixedInterval) *Fixed {

	var (
		fix = &Fixed{
			days:     0,
			duration: 0,
		}
		b = time.Date(0, time.January, 1, 0, 0, 0, 0, time.Local)
		e = b
	)

	for _, fix := range intervals {
		days, duration := fix.GetSize()
		e = e.Add(duration).AddDate(0, 0, days)
	}

	if b.After(e) {
		fix.days, fix.duration = size(e, b)
		fix.days = -1 * fix.days
		fix.duration = -1 * fix.duration
	} else {
		fix.days, fix.duration = size(b, e)
	}

	return fix
}
