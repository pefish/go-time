package go_time

import (
	"errors"
	"time"
)

type TimeUtil struct {
}

var TimeUtilInstance = &TimeUtil{}

type TimeUnit int

const (
	TimeUnit_SECOND      TimeUnit = 1
	TimeUnit_MILLISECOND TimeUnit = 2
)

func (tu *TimeUtil) CurrentTimestamp(unit TimeUnit) int64 {
	if unit == TimeUnit_SECOND {
		return time.Now().Unix()
	} else if unit == TimeUnit_MILLISECOND {
		return int64(time.Now().Nanosecond() / 1000)
	} else {
		panic(errors.New(`unit error`))
	}
}

func (tu *TimeUtil) TimeToStr(time time.Time, format string) string {
	layout := tu.getLayoutFromFormat(format)
	return time.Format(layout)
}

func (tu *TimeUtil) TimestampToStr(timestamp int64, format string, utc bool) string {
	tm := time.Unix(timestamp, 0)
	if utc {
		tm = tm.UTC()
	}
	return tu.TimeToStr(tm, format)
}

func (tu *TimeUtil) getLayoutFromFormat(format string) string {
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

func (tu *TimeUtil) OffsetStrToLocal(format string, str string, offsetHours int) (time.Time, error) {
	t, err := time.ParseInLocation(tu.getLayoutFromFormat(format), str, time.FixedZone("CST", offsetHours*3600))
	if err != nil {
		return time.Time{}, err
	}
	return t.Local(), nil
}

func (tu *TimeUtil) MustOffsetStrToLocal(format string, str string, offsetHours int) time.Time {
	t, err := tu.OffsetStrToLocal(format, str, offsetHours)
	if err != nil {
		panic(err)
	}
	return t
}

func (tu *TimeUtil) StrToLocal(format string, str string, loc *time.Location) (time.Time, error) {
	t, err := time.ParseInLocation(tu.getLayoutFromFormat(format), str, loc)
	if err != nil {
		return time.Time{}, err
	}
	return t.Local(), nil
}

func (tu *TimeUtil) MustStrToLocal(format string, str string, loc *time.Location) time.Time {
	t, err := tu.StrToLocal(format, str, loc)
	if err != nil {
		panic(err)
	}
	return t
}

func (tu *TimeUtil) UtcStrToLocal(format string, str string) (time.Time, error) {
	t, err := time.Parse(tu.getLayoutFromFormat(format), str)
	if err != nil {
		return time.Time{}, err
	}
	return t.Local(), nil
}

func (tu *TimeUtil) MustUtcStrToLocal(format string, str string) time.Time {
	t, err := tu.UtcStrToLocal(format, str)
	if err != nil {
		panic(err)
	}
	return t
}

func (tu *TimeUtil) LocalBeginTimeOfToday() time.Time {
	t := time.Now()
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func (tu *TimeUtil) LocalEndTimeOfToday() time.Time {
	t := time.Now()
	year, month, day := t.Date()
	return time.Date(year, month, day+1, 0, 0, 0, 0, t.Location())
}
