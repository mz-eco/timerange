# TimeRange - 时间区间
描述一个 [begin,end) 的开闭型时间区间、提供对时间区间的操作函数

## Install
```shell
go get -u github.com/mz-eco/timerange
```
## Usage
```go
now, _ := time.Parse(
    time.RFC3339Nano,
    "2018-10-11T15:46:28.132318+08:00")

//创建区间
Range(now,now.Add(24*time.Hour)) //[2018-10-11T15:46:28.132318+08:00 - 2018-10-12T15:46:28.132318+08:00)
RangeAt(now,Day)                 //[2018-10-11T00:00:00+08:00 - 2018-10-12T00:00:00+08:00)
RangeAt(now,Month)				 //[2018-10-01T00:00:00+08:00 - 2018-11-01T00:00:00+08:00)
RangeTo(now,48*Hour)             //[2018-10-11T15:46:28.132318+08:00 - 2018-10-13T15:46:28.132318+08:00)

//时间处理
Truncate(now,Day) //2018-10-11 00:00:00 +0800 CST
Preview(now,Day)  //2018-10-10 00:00:00 +0800 CST
Next(now,Day)     //2018-10-12 00:00:00 +0800 CST
Add(now,3*Day)    //2018-10-14 15:46:28.132318 +0800 CST
Add(now,-3*Day)   //2018-10-08 15:46:28.132318 +0800 CST
Add(now,Duration(24*time.Hour)) //2018-10-12 15:46:28.132318 +0800 CST
Begin(now,Day)    //2018-10-11 00:00:00 +0800 CST
End(now,Day)      //2018-10-11 23:59:59.999999999 +0800 CST

//时间区间
x := RangeAt(now,Day)
x.Begin() //2018-10-11 00:00:00 +0800 CST
x.End()   //2018-10-11 23:59:59.999999999 +0800 CST

/*
[2018-10-11T00:00:00+08:00 - 2018-10-11T08:00:00+08:00)
[2018-10-11T08:00:00+08:00 - 2018-10-11T16:00:00+08:00)
[2018-10-11T16:00:00+08:00 - 2018-10-12T00:00:00+08:00)
 */
x.Split(8*Hour)

x.Contains(now) //true
x.Before(now)   //false
x.After(now)    //false

x.Move(2*Day)   //[2018-10-13T00:00:00+08:00 - 2018-10-14T00:00:00+08:00)
x.Add(2*Day)    //[2018-10-11T00:00:00+08:00 - 2018-10-14T00:00:00+08:00)
x.Add(-2*Minute)//[2018-10-10T23:58:00+08:00 - 2018-10-12T00:00:00+08:00)
x.Sub(1*Hour)   //[2018-10-11T01:00:00+08:00 - 2018-10-12T00:00:00+08:00)
x.Sub(Duration(-2*time.Hour)) //[2018-10-11T00:00:00+08:00 - 2018-10-11T22:00:00+08:00)

x.In(RangeAt(now,Month)) //true
x.Size() //1day 0s
x.Larger(RangeAt(now,Month)) //false
x.Smaller(RangeAt(now,Month)) //true

y := RangeTo(now,2*Day) //[2018-10-11T15:46:28.132318+08:00 - 2018-10-13T15:46:28.132318+08:00)
y.Head(Day)     //[2018-10-11T15:46:28.132318+08:00 - 2018-10-12T00:00:00+08:00)
y.Tail(Day)     //[2018-10-13T00:00:00+08:00 - 2018-10-13T15:46:28.132318+08:00)
y.Truncate(Day) //[2018-10-12T00:00:00+08:00 - 2018-10-13T00:00:00+08:00)

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