package parser

import (
	"fmt"
	"github.com/mz-eco/timerange"
	"strconv"
	"strings"
	"text/scanner"
	"time"
)

func makeTime(tm DateTime, now time.Time) time.Time {

	var (
		y, m, d = now.Date()
	)

	if tm.Date.Year == -1 {
		tm.Date.Year = y
	}

	if tm.Date.Month == 0 {
		tm.Date.Month = int(m)
	}

	if tm.Date.Day == 0 {
		tm.Date.Day = d
	}

	return time.Date(
		tm.Date.Year,
		time.Month(tm.Date.Month),
		tm.Date.Day,
		tm.Time.Hour,
		tm.Time.Minute,
		tm.Time.Second,
		0,
		time.Local)

}

func makeRange(b, e DateTime) timerange.TimeRange {

	var (
		now = time.Now()
	)
	return timerange.Range(
		makeTime(b, now),
		makeTime(e, now),
	)
}

type DateTime struct {
	Date Date
	Time Time
}

func NewDateTime(d Date, t Time) DateTime {
	return DateTime{
		Date: d,
		Time: t,
	}
}

type Date struct {
	Year  int
	Month int
	Day   int
}

func NewDate(y, m, d int) Date {
	return Date{
		Year:  y,
		Month: m,
		Day:   d,
	}
}

type Time struct {
	Hour   int
	Minute int
	Second int
}

func NewTime(h, m, s int) Time {
	return Time{
		Hour:   h,
		Minute: m,
		Second: s,
	}
}

type Lex struct {
	tr      timerange.TimeRange
	scanner scanner.Scanner
}

func NewLex(src string) *Lex {

	lex := new(Lex)
	lex.scanner.Init(strings.NewReader(src))
	lex.scanner.Whitespace = 0
	lex.scanner.Mode = scanner.ScanInts | scanner.ScanIdents

	return lex
}

func (m *Lex) Lex(lval *yySymType) int {

	tok := m.scanner.Scan()

	if tok == scanner.EOF {
		return 0
	}

	if m.scanner.Peek() == 'T' {
		return int(m.scanner.Next())
	}

	text := m.scanner.TokenText()

	switch tok {
	case scanner.Int:

		v, err := strconv.ParseInt(text, 10, 32)

		if err != nil {
			m.Error("parse error")
			return 0
		}

		lval.stmt = int(v)

		switch len(text) {
		case 1, 2:
			return num
		case 4:
			return stmt
		default:
			return big
		}

	case scanner.Ident:

		switch text {
		case "seconds", "s":
			lval.cb = timerange.Seconds
		case "minutes", "m":
			lval.cb = timerange.Minutes
		case "hours", "h":
			lval.cb = timerange.Hours
		case "days", "D":
			lval.cb = timerange.Days
		case "months", "M":
			lval.cb = timerange.Months
		case "years", "Y":
			lval.cb = timerange.Years
		default:
			m.Error("err")
			return 0
		}

		return interval
	default:
		return int(text[0])

	}

	return int(tok)
}

func (m *Lex) Error(s string) {
	fmt.Println(s, m.scanner.Pos())
}
