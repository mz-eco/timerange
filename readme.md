# TimeRange - 时间区间
描述一个 [begin,end) 的开闭型时间区间、提供对时间区间的操作函数

## Install
```shell
go get -u github.com/mz-eco/timerange
```
## Usage
```go
package main

import (
	"fmt"
	. "github.com/mz-eco/timerange"
	"time"
)

func main() {

	days := RangeAt(time.Now(), 2*Day) //[2018-10-12T00:00:00+08:00 - 2018-10-14T00:00:00+08:00)

	for _,block := range days.Split(12*Hour) {
		fmt.Println(block)
	}
}
```
output
```
[2018-10-12T00:00:00+08:00 - 2018-10-12T12:00:00+08:00)
[2018-10-12T12:00:00+08:00 - 2018-10-13T00:00:00+08:00)
[2018-10-13T00:00:00+08:00 - 2018-10-13T12:00:00+08:00)
[2018-10-13T12:00:00+08:00 - 2018-10-14T00:00:00+08:00)
```

## API
### Interval
#### 时间点在时间轴上位移步长的抽象描述

- ```+``` 表示往 --> 方向移动
- ```-``` 表示往 <-- 方向移动

#### construct
```go
//+3 days
3*Day
Days(3)

//-3 hours
-3*Hour
Hours(-3)

//+10 seconds
Duration(10*time.Second)
```

#### builtin intervals

| Name     |construct|value        |
|----------|-----    |:-------------|
| Second   |Seconds(int)|now.AddDuration(n*time.Second) |
| Minute   |Minutes(int)| now.AddDuration(n*time.Minute) |
| Hour     |Hours(int)| now.AddDuration(n*time.Hour)  |
| Day      |Days(int)| now.AddDate(0,0,n)  |
| Month    |Month(int)| now.AddDate(0,n,0)  |
| Year     |Year(int) |now.AddDate(n,0,0)  |
| Week     |Week(int)| to next week's monday|

#### type
| Name     |value       |
|----------|:-------------|
| Duration   | now.AddDuration(n) |

### functions
#### Add
```go
Add(now time.Time, ivs Interval) time.Time
```
shift ```now``` via given ```Interval``` on time axis


### Whole Point
#### 取时间点的整点时间
#### usage
```go
now = time.Now()  //2018-10-12 10:58:43.964305 +0800 CST m=+0.000684927
Truncate(now,Day) //2018-10-12 00:00:00 +0800 CST
Next(now,Hour)    //2018-10-13 00:00:00 +0800 CST
```
#### Functions
##### Truncate
```go
func Truncate(now time.Time, w Whole) time.Time
```
取当前整点
```
12:00       12:01       12:02       12:03
  +-----------+-----------+-----------+
              ^ <-- *
```
##### Next
```go
func Next(now time.Time, w Whole) time.Time
```
取上一个整点
```
12:00       12:01       12:02       12:03
  +-----------+-----------+-----------+
                    * --> ^
```
##### Preview
```go
func Preview(now time.Time, w Whole) time.Time
```
取下一个整点
```
12:00       12:01       12:02       12:03
  +-----------+-----------+-----------+
  ^      <--        *
```
##### Begin
```go
func Begin(now time.Time, w Whole) time.Time
```
整点的开始时间
```go
now = time.Now()  //2018-10-12 10:58:43.964305 +0800 CST m=+0.000684927
Begin（now,Day)   //2018-10-12 00:00:00 +0800 CST
```
##### End
```go
func End(now time.Time, w Whole) time.Time
```
整点的结束时间
```go
now = time.Now()  //2018-10-12 10:58:43.964305 +0800 CST m=+0.000684927
End(now,Day)      //2018-10-12 23:59:59.999999999 +0800 CST
```
#### built in wholes
| Name     |range|
|----------|-----    |
| Second   |[00:00:01,00:00:02)|
| Minute   |[00:00:00,00:01:00)|
| Hour     |[00:00:00,01:00:00)|
| Day      |[2018-01-01,2018-01-02)|
| Month    |[2018-01-01,2018-02-01)|
| Year     |[2018-01-01,2019-01-01)|
| Week     |[Monday,Monday+7Days)|

### TimeRange
time range is a interval on time axis

it is a open,close interval like [begin,end)
```
  12:00                   14:00
 ---+---------+------------+---
    ^                      ^
  begin                   end
```
#### construct
##### Range
```go
func Range(b, e time.Time) TimeRange
```
create a time range by given time
##### RangeAt
```go
func RangeAt(now time.Time, w WholeInterval) TimeRange
```
create a time range from whole time point of ```now``` to given interval
```go
//[2018-10-12T00:00:00+08:00 - 2018-10-14T00:00:00+08:00)
days := RangeAt(time.Now(), 2*Day)
```
##### RangeTo
```go
func RangeTo(b time.Time, iv Interval) TimeRange
```
create a time range from begin time to given interval
```go
//[2018-10-12T13:57:07.665073+08:00 - 2018-10-14T13:57:07.665073+08:00)
RangeTo(time.Now(), 2*Day)
```
