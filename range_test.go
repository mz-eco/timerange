package timerange

import (
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
	tDate = time.Date(2019,time.October,10,11,58,58,0,time.Local)
	tIntervals = []Interval {
		Second,
		Minute,
		Hour,
		Day,
		Month,
		Year,
		Week,
	}
	tWholes = []Whole {
		Second,
		Minute,
		Hour,
		Day,
		Month,
		Year,
		Week,
	}
	tTimes = []time.Time {
		time.Date(2019,time.October,10,11,58,58,0,time.Local),
		time.Date(2019,time.October,10,11,58,00,0,time.Local),
		time.Date(2019,time.October,10,11,00,00,0,time.Local),
		time.Date(2019,time.October,10,00,00,00,0,time.Local),
		time.Date(2019,time.October,01,00,00,00,0,time.Local),
		time.Date(2019,time.January,01,00,00,00,0,time.Local),
		time.Date(2019,time.October,07,00,00,00,0,time.Local),
	}
	tNext = []time.Time {
		time.Date(2019,time.October,10,11,58,59,0,time.Local),
		time.Date(2019,time.October,10,11,59,00,0,time.Local),
		time.Date(2019,time.October,10,12,00,00,0,time.Local),
		time.Date(2019,time.October,11,00,00,00,0,time.Local),
		time.Date(2019,time.November,01,00,00,00,0,time.Local),
		time.Date(2020,time.January,01,00,00,00,0,time.Local),
		time.Date(2019,time.October,14,00,00,00,0,time.Local),
	}
	tPreview = []time.Time {
		time.Date(2019,time.October,10,11,58,57,0,time.Local),
		time.Date(2019,time.October,10,11,57,00,0,time.Local),
		time.Date(2019,time.October,10,10,00,00,0,time.Local),
		time.Date(2019,time.October,9,00,00,00,0,time.Local),
		time.Date(2019,time.September,01,00,00,00,0,time.Local),
		time.Date(2018,time.January,01,00,00,00,0,time.Local),
		time.Date(2019,time.September,30,00,00,00,0,time.Local),
	}
)


func tTimeEqual(t *testing.T, iv interface{},x,y time.Time,) {

	if !x.Equal(y) {
		t.Errorf("time<%s> value<%s> not equal raw<%s>", reflect.TypeOf(iv),x,y)
	}
}

func TestTruncate(t *testing.T) {

	for index, iv := range tWholes {
		tTimeEqual(
			t,
			iv,
			Truncate(tDate,iv), tTimes[index])
	}
}

func TestNext(t *testing.T) {
	for index, iv := range tWholes {
		tTimeEqual(
			t,
			iv,
			Next(tDate,iv), tNext[index])
	}
}

func TestPreview(t *testing.T) {
	for index, iv := range tWholes {
		tTimeEqual(
			t,
			iv,
			Preview(tDate,iv), tPreview[index])
	}
}


