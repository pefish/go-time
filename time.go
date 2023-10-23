package go_time

import (
	"errors"
	"time"
)

type TimeType struct {
}

var TimeInstance = &TimeType{}

type TimeUnit int

const (
	TimeUnit_SECOND      TimeUnit = 1
	TimeUnit_MILLISECOND TimeUnit = 2
)

func (tu *TimeType) CurrentTimestamp(unit TimeUnit) int64 {
	if unit == TimeUnit_SECOND {
		return time.Now().Unix()
	} else if unit == TimeUnit_MILLISECOND {
		return time.Now().UnixMilli()
	} else {
		panic(errors.New(`unit error`))
	}
}

func (tu *TimeType) TimestampToTime(timestamp int64, isToUtc bool) time.Time {
	tm := time.Unix(timestamp, 0)
	if isToUtc {
		tm = tm.UTC()
	}
	return tm
}

func (tu *TimeType) TimeToStr(time time.Time, format string) string {
	layout := tu.getLayoutFromFormat(format)
	return time.Format(layout)
}

func (tu *TimeType) TimestampToStr(timestamp int64, format string, isToUtc bool) string {
	tm := time.Unix(timestamp, 0)
	if isToUtc {
		tm = tm.UTC()
	}
	return tu.TimeToStr(tm, format)
}

func (tu *TimeType) NowToUtcStr() string {
	return tu.TimeToStr(time.Now().UTC(), "0000-00-00 00:00:00")
}

func (tu *TimeType) getLayoutFromFormat(format string) string {
	if format == `0000` {
		return `2006`
	} else if format == `000000000000` {
		return `200601021504`
	} else if format == `0000-00` {
		return `2006-01`
	} else if format == `0000-00-00` {
		return `2006-01-02`
	} else if format == `0000-00-00 00` {
		return `2006-01-02 15`
	} else if format == `0000-00-00 00:00` {
		return `2006-01-02 15:04`
	} else if format == `0000-00-00 00:00:00` {
		return `2006-01-02 15:04:05`
	} else if format == `0000-00-00 00:00:00.000` {
		return `2006-01-02 15:04:05.000`
	} else if format == `0000-00-00T00:00:00Z` {
		return `2006-01-02T15:04:05Z` // UTC时间
	} else if format == `0000-00-00T00:00:00-00:00` {
		return `2006-01-02T15:04:05+08:00` // local时间
	} else {
		panic(errors.New(`format not supported`))
	}
}

func (tu *TimeType) OffsetStrToLocalTime(str string, format string, offsetHours int) (time.Time, error) {
	t, err := time.ParseInLocation(tu.getLayoutFromFormat(format), str, time.FixedZone("CST", offsetHours*3600))
	if err != nil {
		return time.Time{}, err
	}
	return t.Local(), nil
}

func (tu *TimeType) MustOffsetStrToLocalTime(str string, format string, offsetHours int) time.Time {
	t, err := tu.OffsetStrToLocalTime(str, format, offsetHours)
	if err != nil {
		panic(err)
	}
	return t
}

func (tu *TimeType) LocalStrToLocalTime(str string, format string) (time.Time, error) {
	return time.ParseInLocation(tu.getLayoutFromFormat(format), str, time.Local)
}

func (tu *TimeType) MustLocalStrToLocalTime(str string, format string) time.Time {
	t, err := tu.LocalStrToLocalTime(str, format)
	if err != nil {
		panic(err)
	}
	return t
}

func (tu *TimeType) UtcStrToLocalTime(str string, format string) (time.Time, error) {
	t, err := time.ParseInLocation(tu.getLayoutFromFormat(format), str, time.UTC)
	if err != nil {
		return time.Time{}, err
	}
	return t.Local(), nil
}

func (tu *TimeType) MustUtcStrToLocalTime(str string, format string) time.Time {
	t, err := tu.UtcStrToLocalTime(str, format)
	if err != nil {
		panic(err)
	}
	return t
}

func (tu *TimeType) LocalBeginTimeOfToday() time.Time {
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func (tu *TimeType) LocalEndTimeOfToday() time.Time {
	year, month, day := time.Now().Date()
	return time.Date(year, month, day+1, 0, 0, 0, 0, time.Local)
}
