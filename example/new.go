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

	/*
	Second.At: [2018-10-11T15:46:28+08:00 - 2018-10-11T15:46:29+08:00)
	Minute.At: [2018-10-11T15:46:00+08:00 - 2018-10-11T15:47:00+08:00)
	Hour.At  : [2018-10-11T15:00:00+08:00 - 2018-10-11T16:00:00+08:00)
	Day.At   : [2018-10-11T00:00:00+08:00 - 2018-10-12T00:00:00+08:00)
	Month.At : [2018-10-01T00:00:00+08:00 - 2018-11-01T00:00:00+08:00)
	Year.At  : [2018-01-01T00:00:00+08:00 - 2019-01-01T00:00:00+08:00)
	Week.At  : [2018-10-08T00:00:00+08:00 - 2018-10-15T00:00:00+08:00)
	 */

	T("RangeAt(now,Second)", RangeAt(now, Second))
	T("RangeAt(now,Minute)", RangeAt(now, Minute))
	T("RangeAt(now,Hour)  ", RangeAt(now, Hour))
	T("RangeAt(now,Day)   ", RangeAt(now, Day))
	T("RangeAt(now,Month) ", RangeAt(now, Month))
	T("RangeAt(now,Year)  ", RangeAt(now, Year))
	T("RangeAt(now,Week)  ", RangeAt(now, Week))

	/*
	RangeAt(now,Second): [2018-10-11T00:00:00+08:00 - 2018-10-12T00:00:00+08:00)
    RangeAt(now,Minute): [2018-10-11T00:00:00+08:00 - 2018-10-12T00:00:00+08:00)
    RangeAt(now,Hour)  : [2018-10-11T00:00:00+08:00 - 2018-10-12T00:00:00+08:00)
    RangeAt(now,Day)   : [2018-10-11T00:00:00+08:00 - 2018-10-12T00:00:00+08:00)
    RangeAt(now,Month) : [2018-10-11T00:00:00+08:00 - 2018-10-12T00:00:00+08:00)
    RangeAt(now,Year)  : [2018-10-11T00:00:00+08:00 - 2018-10-12T00:00:00+08:00)
    RangeAt(now,Week)  : [2018-10-11T00:00:00+08:00 - 2018-10-12T00:00:00+08:00)
	*/

	//当天时间
}
