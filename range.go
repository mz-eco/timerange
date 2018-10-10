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

func (m TimeRange) Head(iv Whole) (head, body TimeRange) {

	if iv.IsWhole(m.b) {
		return At(m.b, m.b), m
	}

	e := iv.Next(m.b)

	if e.After(m.e) {
		return m, At(m.e, m.e)
	}

	return At(m.b, e), At(e, m.e)
}

func (m TimeRange) Tail(iv Whole) (tail, body TimeRange) {

	if iv.IsWhole(m.e) {
		return At(m.e, m.e), m
	}

	e := iv.Current(m.e)

	if e.Before(m.b) {
		return m, At(m.b, m.b)
	}

	return At(e, m.e), At(m.b, e)
}

func (m TimeRange) Trim(iv Whole) TimeRange {
	return m.TrimLeft(iv).TrimRight(iv)
}

func (m TimeRange) TrimRight(iv Whole) TimeRange {

	if iv.IsWhole(m.e) {
		return At(m.e, m.e)
	}

	e := iv.Current(m.e)

	if e.Before(m.b) {
		return At(m.b, m.b)
	}

	return At(m.b, e)
}

func (m TimeRange) TrimLeft(iv Whole) TimeRange {

	if iv.IsWhole(m.b) {
		return At(m.b, m.b)
	}

	b := iv.Next(m.b)

	if b.After(m.e) {
		return At(m.e, m.e)
	}

	return At(b, m.e)
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

func (m TimeRange) Add(interval Interval) TimeRange {

	return TimeRange{
		b: interval.Add(m.b),
		e: interval.Add(m.e),
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

	return md >= od && mt > ot
}

func (m TimeRange) Smaller(o TimeRange) bool {

	var (
		md, mt = m.Size()
		od, ot = o.Size()
	)

	return md <= od && mt < ot
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

func At(b, e time.Time) TimeRange {

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


func New(b time.Time, iv Interval) TimeRange {
	return At(
		b,
		iv.Add(b))
}

func Now(w Whole) TimeRange {
	var (
		now = time.Now()
	)

	return At(
		w.Current(now),
		w.Next(now))
}


func NowTo(iv Interval) TimeRange {

	var (
		now = time.Now()
	)
	return At(
		now,
		iv.Add(now),
	)
}

func Today() TimeRange {
	return Now(Day)
}


func Split(p TimeRange, iv Interval) []Block {

	var (
		blocks = make([]Block,0)
		iter = NewIterator(p,iv)
	)

	for iter.Next() {
		blocks = append(blocks, iter.Current)
	}

	return blocks
}
