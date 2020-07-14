package localtime

import (
	"database/sql/driver"
	"time"
)

type LocalTime time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
)

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (l LocalTime) String() string {
	return time.Time(l).Format(timeFormat)
}

func (l LocalTime) Now() LocalTime {
	return LocalTime(time.Now())
}

func (l LocalTime) ParseTime(t time.Time) LocalTime {
	return LocalTime(t)
}

func (j LocalTime) format() string {
	return time.Time(j).Format(timeFormat)
}

func (j LocalTime) MarshalText() ([]byte, error) {
	return []byte(j.format()), nil
}

func (l *LocalTime) FromDB(b []byte) error {
	if nil == b || len(b) == 0 {
		l = nil
		return nil
	}
	var now time.Time
	var err error
	now, err = time.ParseInLocation(timeFormat, string(b), time.Local)
	if nil == err {
		*l = LocalTime(now)
		return nil
	}
	now, err = time.ParseInLocation("2006-01-02T15:04:05Z", string(b), time.Local)
	if nil == err {
		*l = LocalTime(now)
		return nil
	}
	// panic("自己定义个layout日期格式处理一下数据库里面的日期型数据解析!")
	return err
}

//func (t *LocalTime) Scan(v interface{}) error {
//	// Should be more strictly to check this type.
//	vt, err := time.Parse("2006-01-02 15:04:05", string(v.([]byte)))
//	if err != nil {
//		return err
//	}
//	*t = LocalTime(vt)
//	return nil
//}

func (l *LocalTime) ToDB() ([]byte, error) {
	if nil == l {
		return nil, nil
	}
	return []byte(time.Time(*l).Format(timeFormat)), nil
}

func (l *LocalTime) Value() (driver.Value, error) {
	if nil == l {
		return nil, nil
	}
	return time.Time(*l).Format(timeFormat), nil
}
