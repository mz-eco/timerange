package timerange

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

func TestToday_Now(t *testing.T) {

	var (
		now = time.Now()
	)

	assert.Equal(
		t,
		Today.Day(),
		Truncate(now, Day))

	assert.Equal(
		t,
		Today.Month(),
		Truncate(now, Month))

	assert.Equal(
		t,
		Today.Year(),
		Truncate(now, Year))

	assert.Equal(
		t,
		Today.Begin(),
		Truncate(now, Day))

	assert.Equal(
		t,
		Today.End(),
		End(now, Day))

	Today.Now()

	y, m, d := now.Date()

	nt := Today.Time(12, 25, 12)

	assert.Equal(
		t,
		nt,
		time.Date(y, m, d, 12, 25, 12, 0, time.Local))
}
