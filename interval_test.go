package timerange

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

func Test_Allow(t *testing.T) {

	var (
		preview = []Interval{
			5 * Second,
			5 * Minute,
			5 * Hour,
			5 * Day,
			5 * Month,
			5 * Year,
			Duration(4 * time.Hour),
		}

		next = []Interval{
			-5 * Second,
			-5 * Minute,
			-5 * Hour,
			-5 * Day,
			-5 * Month,
			-5 * Year,
			Duration(-4 * time.Hour),
		}

		zero = []Interval{
			0 * Second,
			0 * Minute,
			0 * Hour,
			0 * Day,
			0 * Month,
			0 * Year,
			Duration(0),
		}

		equal = func(ivs []Interval, allow Allow) {

			for _, iv := range ivs {
				assert.Equal(t, iv.Allow(), allow)
			}
		}
	)

	equal(preview, AllowForward)
	equal(next, AllowRevert)
	equal(zero, AllowStop)

}

func Test_Now(t *testing.T) {

	assert.Equal(
		t,
		Year.Now(),
		RangeAt(time.Now(), Year))

	assert.Equal(
		t,
		Month.Now(),
		RangeAt(time.Now(), Month))

	assert.Equal(
		t,
		Day.Now(),
		RangeAt(time.Now(), Day))

	assert.Equal(
		t,
		Hour.Now(),
		RangeAt(time.Now(), Hour))

	assert.Equal(
		t,
		Minute.Now(),
		RangeAt(time.Now(), Minute))

	assert.Equal(
		t,
		Second.Now(),
		RangeAt(time.Now(), Second))

}

func Test_Date(t *testing.T) {

	now := time.Date(2018, 1, 1, 12, 12, 12, 0, time.Local)

	assert.Equal(
		t,
		Year.Date(2018),
		Truncate(now, Year))

	assert.Equal(
		t,
		Month.Date(2018, 1),
		Truncate(now, Month))

	assert.Equal(
		t,
		Day.Date(2018, 1, 1),
		Truncate(now, Day))

	assert.Equal(
		t,
		Hour.Date(2018, 1, 1, 12),
		Truncate(now, Hour))

	assert.Equal(
		t,
		Minute.Date(2018, 1, 1, 12, 12),
		Truncate(now, Minute))

	assert.Equal(
		t,
		Second.Date(2018, 1, 1, 12, 12, 12),
		Truncate(now, Second))

	assert.Equal(
		t,
		Today.Range(),
		Day.Today())
}
