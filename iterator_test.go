package timerange

import (
	"fmt"
	"sort"
	"testing"
	"time"
)


func valueCompare(t *testing.T,x[]TimeRange,y []Block) {

	if len(x) != len(y) {
		t.Fatalf("result length src<%d>, cmp<%d> not equal",len(x),len(y))
	}

	for index,value := range x {
		if !value.Equal(y[index].P) {
			t.Errorf("range src<%s> not equal <%s>", value,y[index])
		}

		if !value.Equal(Range(y[index].B,y[index].E)) {
			t.Errorf("range src<%s> not equal <%s>", value,y[index])
		}
	}
}

func printRanges(x []TimeRange) {
	for i, r := range x {
		fmt.Println(i,r)
	}
}


func TestSplitSeconds(t *testing.T) {

	var times = make([]TimeRange,0)

	for i:=time.Duration(0);i<120;i++ {
		times = append(times, RangeTo(tDate.Add(i*time.Second),Second))
	}

	valueCompare(
		t,
		times,
		Split(RangeTo(tDate, 2*Minute),Second))

}

func TestSplitMinute(t *testing.T) {

	var times = make([]TimeRange,0)

	for i:=0;i<120;i++ {
		times = append(times, RangeTo(Add(tDate, Minutes(i)),Minute))
	}

	valueCompare(
		t,
		times,
		Split(RangeTo(tDate, 2*Hour),Minute))

}

func TestSplitHour(t *testing.T) {

	var times = make([]TimeRange,0)

	for i:=0;i<48;i++ {
		times = append(times, RangeTo(Add(tDate, Hours(i)),Hour))
	}

	valueCompare(
		t,
		times,
		Split(RangeTo(tDate, 2*Day),Hour))

}


func TestSplitMonth(t *testing.T) {

	base := Day.Date(2018,01,28)
	begin := base

	var times = make([]TimeRange,0)

	for i:=0;i<120;i++ {
		times = append(times, Range(begin,begin.AddDate(0,1,0)))
		begin = begin.AddDate(0,1,0)
	}

	valueCompare(
		t,
		times,
		Split(RangeTo(base, 10*Year),Month))

}

func TestSplitYear(t *testing.T) {

	base := Day.Date(2018,01,28)
	begin := base

	var times = make([]TimeRange,0)

	for i:=0;i<100;i++{
		times = append(times, Range(begin,begin.AddDate(1,0,0)))
		begin = begin.AddDate(1,0,0)
	}

	valueCompare(
		t,
		times,
		Split(RangeTo(base, 100*Year),Year))

}

func TestSplitYearRevert(t *testing.T) {



	base := Day.Date(2018,01,01)
	begin := base

	var times = make([]TimeRange,0)

	for i:=0;i<100;i++{
		times = append(times, Range(begin,begin.AddDate(1,0,0)))
		begin = begin.AddDate(1,0,0)
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].b.After(times[j].b)
	})

	valueCompare(
		t,
		times,
		Split(RangeTo(base, 100*Year),-1*Year))

}

func TestSplitMonthRevert(t *testing.T) {

	base := Day.Date(2018,01,28)
	begin := base

	var times = make([]TimeRange,0)

	for i:=0;i<120;i++ {
		times = append(times, Range(begin,begin.AddDate(0,1,0)))
		begin = begin.AddDate(0,1,0)
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].b.After(times[j].b)
	})

	valueCompare(
		t,
		times,
		Split(RangeTo(base, 10*Year),-1*Month))

}

func TestSplitHourRevert(t *testing.T) {

	var times = make([]TimeRange,0)

	for i:=0;i<48;i++ {
		times = append(times, RangeTo(Add(tDate, Hours(i)),Hour))
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].b.After(times[j].b)
	})

	valueCompare(
		t,
		times,
		Split(RangeTo(tDate, 2*Day),-1*Hour))

}

func TestSplitMinuteRevert(t *testing.T) {

	var times = make([]TimeRange,0)

	for i:=0;i<120;i++ {
		times = append(times, RangeTo(Add(tDate, Minutes(i)),Minute))
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].b.After(times[j].b)
	})

	valueCompare(
		t,
		times,
		Split(RangeTo(tDate, 2*Hour),-1*Minute))

}

func TestSplitSecondsRevert(t *testing.T) {

	var times = make([]TimeRange,0)

	for i:=time.Duration(0);i<120;i++ {
		times = append(times, RangeTo(tDate.Add(i*time.Second),Second))
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].b.After(times[j].b)
	})

	valueCompare(
		t,
		times,
		Split(RangeTo(tDate, 2*Minute),-1*Second))

}

