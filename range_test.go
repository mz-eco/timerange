package timerange

import (
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
	"time"
)

func tParseTime(value string) time.Time {

	tm, err := time.Parse("2006-01-02 15:04:05", value)

	if err != nil {
		panic(err)
	}

	return tm
}

var (
	tDate      = time.Date(2019, time.October, 10, 11, 58, 58, 0, time.Local)
	tIntervals = []Interval{
		Second,
		Minute,
		Hour,
		Day,
		Month,
		Year,
		Week,
	}
	tAddTwoDays = []Interval{
		2 * Second,
		2 * Minute,
		2 * Hour,
		2 * Day,
		2 * Month,
		2 * Year,
		2 * Week,
	}
	tSubTowDays = []Interval{
		-2 * Second,
		-2 * Minute,
		-2 * Hour,
		-2 * Day,
		-2 * Month,
		-2 * Year,
		-2 * Week,
	}
	tWholes = []Whole{
		Second,
		Minute,
		Hour,
		Day,
		Month,
		Year,
		Week,
	}
	tTimes = []time.Time{
		time.Date(2019, time.October, 10, 11, 58, 58, 0, time.Local),
		time.Date(2019, time.October, 10, 11, 58, 00, 0, time.Local),
		time.Date(2019, time.October, 10, 11, 00, 00, 0, time.Local),
		time.Date(2019, time.October, 10, 00, 00, 00, 0, time.Local),
		time.Date(2019, time.October, 01, 00, 00, 00, 0, time.Local),
		time.Date(2019, time.January, 01, 00, 00, 00, 0, time.Local),
		time.Date(2019, time.October, 07, 00, 00, 00, 0, time.Local),
	}
	tNext = []time.Time{
		time.Date(2019, time.October, 10, 11, 58, 59, 0, time.Local),
		time.Date(2019, time.October, 10, 11, 59, 00, 0, time.Local),
		time.Date(2019, time.October, 10, 12, 00, 00, 0, time.Local),
		time.Date(2019, time.October, 11, 00, 00, 00, 0, time.Local),
		time.Date(2019, time.November, 01, 00, 00, 00, 0, time.Local),
		time.Date(2020, time.January, 01, 00, 00, 00, 0, time.Local),
		time.Date(2019, time.October, 14, 00, 00, 00, 0, time.Local),
	}
	tPreview = []time.Time{
		time.Date(2019, time.October, 10, 11, 58, 57, 0, time.Local),
		time.Date(2019, time.October, 10, 11, 57, 00, 0, time.Local),
		time.Date(2019, time.October, 10, 10, 00, 00, 0, time.Local),
		time.Date(2019, time.October, 9, 00, 00, 00, 0, time.Local),
		time.Date(2019, time.September, 01, 00, 00, 00, 0, time.Local),
		time.Date(2018, time.January, 01, 00, 00, 00, 0, time.Local),
		time.Date(2019, time.September, 30, 00, 00, 00, 0, time.Local),
	}
	tAddTwoDay = []time.Time{
		tDate.Add(2 * time.Second),
		tDate.Add(2 * time.Minute),
		tDate.Add(2 * time.Hour),
		tDate.AddDate(0, 0, 2),
		tDate.AddDate(0, 2, 0),
		tDate.AddDate(2, 0, 0),
		time.Date(2019, time.October, 21, 11, 58, 58, 0, time.Local),
	}
	tAddWeek   = time.Date(2019, time.October, 14, 11, 58, 58, 0, time.Local)
	tSubTwoDay = []time.Time{
		tDate.Add(-2 * time.Second),
		tDate.Add(-2 * time.Minute),
		tDate.Add(-2 * time.Hour),
		tDate.AddDate(0, 0, -2),
		tDate.AddDate(0, -2, 0),
		tDate.AddDate(-2, 0, 0),
		time.Date(2019, time.September, 23, 11, 58, 58, 0, time.Local),
	}
	tSubWeek = time.Date(2019, time.September, 30, 11, 58, 58, 0, time.Local)
)

func tTimeEqual(t *testing.T, iv interface{}, x, y time.Time) {

	if !x.Equal(y) {
		t.Errorf("time<%s> value<%s> not equal raw<%s>", reflect.TypeOf(iv), x, y)
	}
}

func TestTruncate(t *testing.T) {

	for index, iv := range tWholes {
		tTimeEqual(
			t,
			iv,
			Truncate(tDate, iv), tTimes[index])
	}
}

func TestNext(t *testing.T) {
	for index, iv := range tWholes {
		tTimeEqual(
			t,
			iv,
			Next(tDate, iv), tNext[index])
	}
}

func TestPreview(t *testing.T) {
	for index, iv := range tWholes {
		tTimeEqual(
			t,
			iv,
			Preview(tDate, iv), tPreview[index])
	}
}

func TestBegin(t *testing.T) {
	for index, iv := range tWholes {
		tTimeEqual(
			t,
			iv,
			Begin(tDate, iv), tTimes[index])
	}
}

func TestEnd(t *testing.T) {

	var (
		next  = Day.Date(2018, 1, 2)
		now   = next.Add(-12 * time.Hour)
		equal = next.Add(-1 * time.Nanosecond)
	)

	assert.Equal(
		t,
		End(now, Day),
		equal,
		"days not equal",
	)
}

func TestAddTwoDays(t *testing.T) {
	for index, iv := range tAddTwoDays {
		tTimeEqual(
			t,
			iv,
			Add(tDate, iv), tAddTwoDay[index])
	}

	tTimeEqual(
		t,
		1*Week,
		Add(tDate, 1*Week),
		tAddWeek,
	)
}

func TestTimeRange_BeginEnd(t *testing.T) {

	now := Month.Date(2018,9)
	tr := RangeAt(now,1*Month)

	assert.Equal(
		t,
		tr.Begin(),
		Begin(now,Month))

	assert.Equal(
		t,
		tr.End(),
		End(now,Month))
}

func TestTimeRange_Trim(t *testing.T) {

	now := Month.Date(2018,time.June)
	head := now.AddDate(0,0,-3)
	tail := Next(now,Month).AddDate(0,0,3)

	tr := Range(head,tail)

	assert.Equal(
		t,
		tr.Trim(Month),
		RangeTo(now,Month))

	assert.Equal(
		t,
		RangeAt(now,Month).Trim(Day),
		RangeTo(now,Month))

}

func TestSubTwoDays(t *testing.T) {
	for index, iv := range tSubTowDays {
		tTimeEqual(
			t,
			iv,
			Add(tDate, iv), tSubTwoDay[index])
	}

	tTimeEqual(
		t,
		-1*Week,
		Add(tDate, -1*Week),
		tSubWeek,
	)
}

func TestNewAdd(t *testing.T) {

	for index, iv := range tAddTwoDays {

		tr := RangeTo(tDate, iv)

		tTimeEqual(
			t,
			iv,
			tr.b, tDate)

		tTimeEqual(
			t,
			iv,
			tr.e,
			tAddTwoDay[index])
	}

	tr := RangeTo(tDate, 1*Week)

	tTimeEqual(
		t,
		1*Week,
		tr.b, tDate)

	tTimeEqual(
		t,
		1*Week,
		tr.e,
		tAddWeek)
}

func TestTimeRange_Equal(t *testing.T) {

	if !(RangeTo(tDate, Day).Equal(RangeTo(tDate, 24*Hour))) {
		t.Errorf("time range not equal")
	}
}

func TestTimeRange_Size(t *testing.T) {

	var (
		tr          = RangeTo(Year.Date(2018), 128*Day, 12*Hour, 12*Minute)
		days, hours = tr.Size()
	)

	assert.Equal(t, days, 128, "days is not 128")
	assert.Equal(t, hours, 12*time.Hour+12*time.Minute, "hours is not 12")

}

func TestNewSub(t *testing.T) {

	for index, iv := range tSubTowDays {

		tr := RangeTo(tDate, iv)

		tTimeEqual(
			t,
			iv,
			tr.e, tDate)

		tTimeEqual(
			t,
			iv,
			tr.b,
			tSubTwoDay[index])
	}

	tr := RangeTo(tDate, -1*Week)

	tTimeEqual(
		t,
		1*Week,
		tr.e, tDate)

	tTimeEqual(
		t,
		1*Week,
		tr.b,
		tSubWeek)

}
