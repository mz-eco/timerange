# TimeRange - 时间区间
描述一个 [begin,end) 的开闭型时间区间、提供对时间区间的操作函数

## Install
```shell
go get -u github.com/mz-eco/timerange
```
## Usage
```go
now, _ := time.Parse(time.RFC3339Nano,"2018-10-11T15:46:28.132318+08:00")
```

创建区间
```go
Range(now,now.Add(24*time.Hour)) //[2018-10-11T15:46:28.132318+08:00 - 2018-10-12T15:46:28.132318+08:00)
RangeAt(now,Day)                 //[2018-10-11T00:00:00+08:00 - 2018-10-12T00:00:00+08:00)
RangeAt(now,Month)				 //[2018-10-01T00:00:00+08:00 - 2018-11-01T00:00:00+08:00)
RangeTo(now,48*Hour)             //[2018-10-11T15:46:28.132318+08:00 - 2018-10-13T15:46:28.132318+08:00)
```
时间处理
```go
Truncate(now,Day) //2018-10-11 00:00:00 +0800 CST
Preview(now,Day)  //2018-10-10 00:00:00 +0800 CST
Next(now,Day)     //2018-10-12 00:00:00 +0800 CST
Add(now,3*Day)    //2018-10-14 15:46:28.132318 +0800 CST
Add(now,-3*Day)   //2018-10-08 15:46:28.132318 +0800 CST
Add(now,Duration(24*time.Hour)) //2018-10-12 15:46:28.132318 +0800 CST
Begin(now,Day)    //2018-10-11 00:00:00 +0800 CST
End(now,Day)      //2018-10-11 23:59:59.999999999 +0800 CST
```

区间-时间
```go
x := RangeAt(now,Day)
x.Begin() //2018-10-11 00:00:00 +0800 CST
x.End()   //2018-10-11 23:59:59.999999999 +0800 CST
x.Contains(now) //true
x.Before(now)   //false
x.After(now)    //false
```
区间位置
```go

x.Size() //1day 0s
x.In(RangeAt(now,Month)) //true
x.Larger(RangeAt(now,Month)) //false
x.Smaller(RangeAt(now,Month)) //true


x.Move(2*Day)   //[2018-10-13T00:00:00+08:00 - 2018-10-14T00:00:00+08:00)
x.Add(2*Day)    //[2018-10-11T00:00:00+08:00 - 2018-10-14T00:00:00+08:00)
x.Add(-2*Minute)//[2018-10-10T23:58:00+08:00 - 2018-10-12T00:00:00+08:00)
x.Sub(1*Hour)   //[2018-10-11T01:00:00+08:00 - 2018-10-12T00:00:00+08:00)
x.Sub(Duration(-2*time.Hour)) //[2018-10-11T00:00:00+08:00 - 2018-10-11T22:00:00+08:00)



```
区间切割
```go
y := RangeTo(now,2*Day) //[2018-10-11T15:46:28.132318+08:00 - 2018-10-13T15:46:28.132318+08:00)
y.Head(Day)     //[2018-10-11T15:46:28.132318+08:00 - 2018-10-12T00:00:00+08:00)
y.Tail(Day)     //[2018-10-13T00:00:00+08:00 - 2018-10-13T15:46:28.132318+08:00)
y.Truncate(Day) //[2018-10-12T00:00:00+08:00 - 2018-10-13T00:00:00+08:00)

x.Split(8*Hour)
/*
[2018-10-11T00:00:00+08:00 - 2018-10-11T08:00:00+08:00)
[2018-10-11T08:00:00+08:00 - 2018-10-11T16:00:00+08:00)
[2018-10-11T16:00:00+08:00 - 2018-10-12T00:00:00+08:00)
 */

//正向迭代器
i := NewIterator(y,4*Hour)

for i.Next() {
    fmt.Println(i.Current)
}

//反向迭代器
i = NewIterator(y,-4*Hour)

for i.Next() {
    fmt.Println(i.Current)
}
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
