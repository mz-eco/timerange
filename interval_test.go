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
			5*Week,
			Duration(4 * time.Hour),
		}

		next = []Interval{
			-5 * Second,
			-5 * Minute,
			-5 * Hour,
			-5 * Day,
			-5 * Month,
			-5 * Year,
			-5*Week,
			Duration(-4 * time.Hour),
		}

		zero = []Interval{
			0 * Second,
			0 * Minute,
			0 * Hour,
			0 * Day,
			0 * Month,
			0 * Year,
			0*Week,
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

func TestRangeTo(t *testing.T) {

	now := time.Date(2018, 1, 1, 12, 12, 12, 0, time.Local)


	assert.Equal(
		t,
		Year.To(now,4),
		RangeAt(now,4*Year),
		"Year",
	)
	assert.Equal(
		t,
		Month.To(now,4),
		RangeAt(now,4*Month),
		"Month",
	)
	assert.Equal(
		t,
		Day.To(now,4),
		RangeAt(now,4*Day),
		"Day",
	)
	assert.Equal(
		t,
		Hour.To(now,4),
		RangeAt(now,4*Hour),
		"Hour",
	)
	assert.Equal(
		t,
		Minute.To(now,4),
		RangeAt(now,4*Minute),
		"Minute",
	)
	assert.Equal(
		t,
		Second.To(now,4),
		RangeAt(now,4*Second),
		"Second",
	)
	assert.Equal(
		t,
		Week.To(now,4),
		RangeAt(now,4*Week),
		"Week",
	)

}

func TestConstruct(t *testing.T) {

	assert.Equal(
		t,
		Days(10),
		10*Day)

	assert.Equal(
		t,
		Years(10),
		10*Year)
	assert.Equal(
		t,
		Months(10),
		10*Month)
	assert.Equal(
		t,
		Hours(10),
		10*Hour)
	assert.Equal(
		t,
		Minutes(10),
		10*Minute)
	assert.Equal(
		t,
		Seconds(10),
		10*Second)
	assert.Equal(
		t,
		Weeks(10),
		10*Week)

}

func TestDay_Range(t *testing.T) {

	begin := Second.Date(2018,1,1,12,12,12)
	end := Second.Date(2018,1,20,12,12,12)

	assert.Equal(
		t,
		Day.Range(begin,end),
		Range(Truncate(begin,Day),Truncate(end,Day)))

	assert.Equal(
		t,
		Month.Range(begin,end),
		Range(Truncate(begin,Month),Truncate(end,Month)))
	assert.Equal(
		t,
		Year.Range(begin,end),
		Range(Truncate(begin,Year),Truncate(end,Year)))
	assert.Equal(
		t,
		Hour.Range(begin,end),
		Range(Truncate(begin,Hour),Truncate(end,Hour)))
	assert.Equal(
		t,
		Minute.Range(begin,end),
		Range(Truncate(begin,Minute),Truncate(end,Minute)))
	assert.Equal(
		t,
		Second.Range(begin,end),
		Range(Truncate(begin,Second),Truncate(end,Second)))
	assert.Equal(
		t,
		Week.Range(begin,end),
		Range(Truncate(begin,Week),Truncate(end,Week)))

}

func Test_At(t *testing.T) {

	now := time.Date(2018, 1, 1, 12, 12, 12, 0, time.Local)

	assert.Equal(
		t,
		Year.At(now),
		RangeAt(now,Year),
		"Year",
		)
	assert.Equal(
		t,
		Month.At(now),
		RangeAt(now,Month),
		"Month",
	)
	assert.Equal(
		t,
		Day.At(now),
		RangeAt(now,Day),
		"Day",
	)
	assert.Equal(
		t,
		Hour.At(now),
		RangeAt(now,Hour),
		"Hour",
	)
	assert.Equal(
		t,
		Minute.At(now),
		RangeAt(now,Minute),
		"Minute",
	)
	assert.Equal(
		t,
		Second.At(now),
		RangeAt(now,Second),
		"Second",
	)
	assert.Equal(
		t,
		Week.At(now),
		RangeAt(now,Week),
		"Week",
	)

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
