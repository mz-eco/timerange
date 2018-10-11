package timerange

import "time"

type Allow int

const (

	//时间点向未来方向移动
	AllowForward Allow = 0

	//时间点向历史方向移动
	AllowRevert = 1

	//保持原状态，不移动时间点
	AllowStop = 2
)

//带移动方向时间间隔，时间点按照间隔步长在时间轴上进行移动
type Interval interface {

	//将时间点移动一个步长
	//移动方向由NextAllow()函数决定
	//AllowForward:
	//+-----------+-----------+-----------+
	//                  *    -->    ^
	//
	//AllowRevert:
	//+-----------+-----------+-----------+
	//      ^    <--    *
	//
	//
	//AllowStop:
	//+-----------+-----------+-----------+
	//                  *
	//                  ^
	AddTo(now time.Time) time.Time

	//时间点在数轴上的移动方向
	Allow() Allow
}

//固定长度的时间间隔
type FixedInterval interface {
	Interval
	GetSize() (days int, duration time.Duration)
}

//可取整的时间间隔
type Whole interface {

	//将时间点移动到下一个整点
	//12:00       12:01       12:02       12:03
	//  +-----------+-----------+-----------+
	//                    * --> ^
	Next(now time.Time) time.Time

	//将时间点移动到当前个整点
	//12:00       12:01       12:02       12:03
	//  +-----------+-----------+-----------+
	//              ^ <-- *
	Current(now time.Time) time.Time


	//将时间点移动到上一个整点
	//12:00       12:01       12:02       12:03
	//  +-----------+-----------+-----------+
	//  ^      <--        *
	Preview(now time.Time) time.Time

	//判断时间点是否是整点时间
	IsWhole(now time.Time) bool
}

type WholeInterval interface {
	Whole
	Interval
}



type durationUnitWhole interface {
	GetUnit() time.Duration
}
