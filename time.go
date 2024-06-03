package go_time

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type TimeType struct {
}

var TimeInstance = &TimeType{}

func (tu *TimeType) CurrentTimestamp() uint64 {
	return uint64(time.Now().UnixMilli())
}

func (tu *TimeType) TimestampToTime(timestamp uint64, isToUtc bool) time.Time {
	tm := time.UnixMilli(int64(timestamp))
	if isToUtc {
		tm = tm.UTC()
	}
	return tm
}

func (tu *TimeType) TimeToTimestamp(time_ time.Time) int64 {
	return time_.UnixMilli()
}

func (tu *TimeType) TimeToStr(time_ time.Time, toFormat string) string {
	layout := tu.getLayoutFromFormat(toFormat)
	return time_.Format(layout)
}

func (tu *TimeType) TimestampToStr(timestamp uint64, format string, isToUtc bool) string {
	tm := time.UnixMilli(int64(timestamp))
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

func (tu *TimeType) getLayout(str string) (string, error) {
	if len(str) == 4 {
		return `2006`, nil
	}

	if len(str) == 12 {
		return `200601021504`, nil
	}

	if len(str) == 7 && str[4] == '-' {
		return `2006-01`, nil
	}

	if len(str) == 10 && str[4] == '-' && str[7] == '-' {
		return `2006-01-02`, nil
	}

	if len(str) == 13 && str[4] == '-' && str[7] == '-' && str[10] == ' ' {
		return `2006-01-02 15`, nil
	}

	if len(str) == 16 && str[4] == '-' && str[7] == '-' && str[10] == ' ' && str[13] == ':' {
		return `2006-01-02 15:04`, nil
	}

	if len(str) == 19 && str[4] == '-' && str[7] == '-' && str[10] == ' ' && str[13] == ':' && str[16] == ':' {
		return `2006-01-02 15:04:05`, nil
	}

	if len(str) == 23 && str[4] == '-' && str[7] == '-' && str[10] == ' ' && str[13] == ':' && str[16] == ':' && str[19] == '.' {
		return `2006-01-02 15:04:05.000`, nil
	}

	if len(str) == 20 && str[4] == '-' && str[7] == '-' && str[10] == 'T' && str[13] == ':' && str[16] == ':' && str[19] == 'Z' {
		return `2006-01-02T15:04:05Z`, nil
	}

	if len(str) == 25 && str[4] == '-' && str[7] == '-' && str[10] == 'T' && str[13] == ':' && str[16] == ':' && (str[19] == '-' || str[19] == '+') {
		return `2006-01-02T15:04:05+08:00`, nil
	}

	if str[4] == '-' && str[7] == '-' && str[10] == ' ' && str[13] == ':' && str[16] == ':' && strings.HasSuffix(str, " +0000 UTC") {
		return "2006-01-02 15:04:05.999999999 +0000 UTC", nil
	}

	return "", errors.New(fmt.Sprintf("TimeStr <%s> format error.", str))
}

func (tu *TimeType) MustStrToTime(str string, isFromUtc bool, isToUtc bool) time.Time {
	t, err := tu.StrToTime(str, isFromUtc, isToUtc)
	if err != nil {
		panic(err)
	}
	return t
}

func (tu *TimeType) StrToTime(str string, isFromUtc bool, isToUtc bool) (time.Time, error) {
	var loc *time.Location
	if isFromUtc {
		loc = time.UTC
	} else {
		loc = time.Local
	}
	layout, err := tu.getLayout(str)
	if err != nil {
		return time.Time{}, err
	}
	t, err := time.ParseInLocation(layout, str, loc)
	if err != nil {
		return time.Time{}, err
	}
	if isToUtc {
		t = t.UTC()
	}
	return t, nil
}

func (tu *TimeType) StrToTimestamp(str string, isFromUtc bool) (int64, error) {
	var loc *time.Location
	if isFromUtc {
		loc = time.UTC
	} else {
		loc = time.Local
	}
	layout, err := tu.getLayout(str)
	if err != nil {
		return 0, err
	}
	t, err := time.ParseInLocation(layout, str, loc)
	if err != nil {
		return 0, err
	}
	return t.UnixMilli(), nil
}

func (tu *TimeType) BeginOfTime(time_ time.Time, isToUtc bool) time.Time {
	year, month, day := time_.Date()
	t := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	if isToUtc {
		t = t.UTC()
	}
	return t
}

func (tu *TimeType) EndOfTime(time_ time.Time, isToUtc bool) time.Time {
	year, month, day := time_.Date()
	t := time.Date(year, month, day+1, 0, 0, 0, 0, time.Local)
	if isToUtc {
		t = t.UTC()
	}
	return t
}
