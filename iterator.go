package timerange

import (
	"time"
)

type Block struct {
	B time.Time
	E time.Time
	P TimeRange
}

func (m Block) String() string {
	return m.P.String()
}

func NewBlock(b, e time.Time) Block {
	return Block{
		B: b,
		E: e,
		P: At(b, e),
	}
}

type Iterator struct {
	p       TimeRange
	c       time.Time
	Current Block
	iv      Interval
	a       Allow
	i int
}

func (m *Iterator) Index() int {
	return m.i
}

func NewIterator(p TimeRange, iv Interval) *Iterator {

	var (
		c time.Time
		a = iv.Allow()
	)

	switch a {
	case AllowForward:
		c = p.b
	case AllowRevert:
		c = p.e

	}
	return &Iterator{
		p:  p,
		c:  c,
		iv: iv,
		a:  a,
	}
}

func (m *Iterator) Next() bool {

	switch m.a {
	case AllowForward:

		if m.c.After(m.p.e) || m.c.Equal(m.p.e) {
			return false
		}

		next := m.iv.Add(m.c)

		if !next.Before(m.p.e) {
			m.Current = NewBlock(m.c, m.p.e)
		} else {
			m.Current = NewBlock(m.c, next)
		}

		m.c = next
		m.i ++

		return true
	case AllowRevert:

		if m.c.Before(m.p.b) || m.c.Equal(m.p.b) {
			return false
		}

		next := m.iv.Add(m.c)

		if !next.After(m.p.b) {
			m.Current = NewBlock(m.p.b, m.c)
		} else {
			m.Current = NewBlock(next, m.c)
		}

		m.c = next
		m.i++
		return true
	default:
		return false
	}
}
