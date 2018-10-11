package timerange

import "time"

type Duration time.Duration

func (m Duration) AddTo(o time.Time) time.Time {
	return o.Add(time.Duration(m))
}

func (m Duration) Allow() Allow {
	switch {
	case m > 0:
		return AllowForward
	case m < 0:
		return AllowRevert
	default:
		return AllowStop
	}
}
