package main

import (
	"fmt"
	. "github.com/mz-eco/timerange"
	"time"
)

func T(name string, v interface{}) {
	fmt.Println(
		fmt.Sprintf(
			"%s: %s", name, v))
}

func main() {

	now, _ := time.Parse(
		time.RFC3339Nano,
		"2018-10-11T15:46:28.132318+08:00")

	//获取时间所在的时间区间
	T("Second.At", Second.At(now))
	T("Minute.At", Minute.At(now))
	T("Hour.At  ", Hour.At(now))
	T("Day.At   ", Day.At(now))
	T("Month.At ", Month.At(now))
	T("Year.At  ", Year.At(now))
	T("Week.At  ", Week.At(now))

	T("RangeAt(now,Second)", RangeAt(now, Second))
	T("RangeAt(now,Minute)", RangeAt(now, Minute))
	T("RangeAt(now,Hour)  ", RangeAt(now, Hour))
	T("RangeAt(now,Day)   ", RangeAt(now, Day))
	T("RangeAt(now,Month) ", RangeAt(now, Month))
	T("RangeAt(now,Year)  ", RangeAt(now, Year))
	T("RangeAt(now,Week)  ", RangeAt(now, Week))

}
