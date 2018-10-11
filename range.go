package timerange

import (
	"fmt"
	"time"
)

type TimeRange struct {
	b time.Time
	e time.Time
}

func (m TimeRange) Truncate(iv Whole) (head, tail, body TimeRange) {

	head, body = m.Head(iv)
	tail, body = body.Tail(iv)

	return
}

func (m TimeRange) Day() TimeRange {
	return Day.At(m.b)
}

func (m TimeRange) Month() TimeRange {
	return Month.At(m.b)
}

func (m TimeRange) Year() TimeRange {
	return Year.At(m.b)
}

func (m TimeRange) Begin() time.Time {
	return m.b
}

func (m TimeRange) Split(iv Interval) Blocks {
	return Split(m, iv)
}

func (m TimeRange) End() time.Time {
	return m.e.Add(-1 * time.Nanosecond)
}

func (m TimeRange) Head(iv Whole) (head, body TimeRange) {

	if iv.IsWhole(m.b) {
		return Range(m.b, m.b), m
	}

	e := iv.Next(m.b)

	if e.After(m.e) {
		return m, Range(m.e, m.e)
	}

	return Range(m.b, e), Range(e, m.e)
}

func (m TimeRange) Tail(iv Whole) (tail, body TimeRange) {

	if iv.IsWhole(m.e) {
		return Range(m.e, m.e), m
	}

	e := iv.Current(m.e)

	if e.Before(m.b) {
		return m, Range(m.b, m.b)
	}

	return Range(e, m.e), Range(m.b, e)
}

func (m TimeRange) Trim(iv Whole) TimeRange {
	return m.TrimLeft(iv).TrimRight(iv)
}

func (m TimeRange) TrimRight(iv Whole) TimeRange {

	if iv.IsWhole(m.e) {
		return Range(m.e, m.e)
	}

	e := iv.Current(m.e)

	if e.Before(m.b) {
		return Range(m.b, m.b)
	}

	return Range(m.b, e)
}

func (m TimeRange) TrimLeft(iv Whole) TimeRange {

	if iv.IsWhole(m.b) {
		return Range(m.b, m.b)
	}

	b := iv.Next(m.b)

	if b.After(m.e) {
		return Range(m.e, m.e)
	}

	return Range(b, m.e)
}

func (m TimeRange) Empty() bool {
	return m.b.Equal(m.e)
}

func (m TimeRange) Contains(time time.Time) bool {

	if !(time.Before(m.b)) && time.Before(m.e) {
		return true
	}

	return false

}
func (m TimeRange) Format(format string) string {

	return fmt.Sprintf(
		"[%s,%s)",
		m.b.Format(format),
		m.e.Format(format))
}

func (m TimeRange) Move(interval Interval) TimeRange {

	return TimeRange{
		b: interval.AddTo(m.b),
		e: interval.AddTo(m.e),
	}
}

func (m TimeRange) Add(ivs Interval) TimeRange {

	switch ivs.Allow() {
	case AllowForward:
		return Range(m.b, Add(m.e, ivs))
	case AllowRevert:
		return Range(Add(m.b, ivs), m.e)
	default:
		return m
	}
}

func (m TimeRange) Sub(ivs Interval) TimeRange {

	switch ivs.Allow() {
	case AllowForward:
		var (
			next = Add(m.b, ivs)
		)

		if next.After(m.e) {
			return Range(m.e, m.e)
		}

		return Range(next, m.e)

	case AllowRevert:
		var (
			next = Add(m.e, ivs)
		)

		if next.Before(m.b) {
			return Range(m.b, m.b)
		}

		return Range(m.b, next)
	default:
		return m
	}
}

func (m TimeRange) Duration() time.Duration {
	return m.e.Sub(m.b)
}

func (m TimeRange) Size() (days int, duration time.Duration) {

	head, tail, body := m.Truncate(Day)

	duration = head.Duration() + tail.Duration()

	if !body.IsZero() {
		days = dayAt(body.e.Date()) - dayAt(body.b.Date())
	}

	if duration >= 24*time.Hour {
		days = days + 1
		duration -= 24 * time.Hour
	}

	return

}

func (m TimeRange) Larger(o TimeRange) bool {

	var (
		md, mt = m.Size()
		od, ot = o.Size()
	)

	switch {
	case md == od:
		return mt > ot
	default:
		return md > od
	}
}

func (m TimeRange) Smaller(o TimeRange) bool {

	var (
		md, mt = m.Size()
		od, ot = o.Size()
	)

	switch {
	case md == od:
		return mt < ot
	default:
		return md < od
	}
}

func (m TimeRange) In(o TimeRange) bool {
	return !(m.b.Before(o.b)) && !(m.e.After(o.e))
}

func (m TimeRange) Time() (b, e time.Time) {
	return m.b, m.e
}

func (m TimeRange) Before(time time.Time) bool {
	return !m.e.After(time)
}

func (m TimeRange) After(time time.Time) bool {
	return m.b.After(time)
}

func (m TimeRange) Equal(other TimeRange) bool {
	return m.b.Equal(other.b) && m.e.Equal(other.e)
}

func (m TimeRange) String() string {
	return fmt.Sprintf(
		"[%s - %s)",
		m.b.Format(time.RFC3339Nano),
		m.e.Format(time.RFC3339Nano),
	)
}

func (m TimeRange) DayIndex() int {
	return m.b.Day() - 1
}

type IteratorType int

func (m TimeRange) IsZero() bool {
	return m.b.Equal(m.e)
}

func Range(b, e time.Time) TimeRange {

	if b.After(e) {
		return TimeRange{
			b: e,
			e: b,
		}
	}

	return TimeRange{
		b: b,
		e: e,
	}
}

func RangeAt(now time.Time, w Whole) TimeRange {

	return Range(
		w.Current(now),
		w.Next(now),
	)
}

func RangeTo(b time.Time, iv Interval) TimeRange {
	return Range(
		b,
		iv.AddTo(b))
}

func NowTo(iv Interval) TimeRange {

	var (
		now = time.Now()
	)
	return Range(
		now,
		iv.AddTo(now),
	)
}

func Split(p TimeRange, iv Interval) Blocks {

	var (
		blocks   = make(Blocks, 0)
		iterator = NewIterator(p, iv)
	)

	for iterator.Next() {
		blocks = append(blocks, iterator.Current)
	}

	return blocks
}
