package go_time

import (
	"errors"
	"time"
)

type TimeClass struct {
}

var Time = TimeClass{}

type TimeUnit int

const (
	TimeUnit_SECOND      TimeUnit = 1
	TimeUnit_MILLISECOND TimeUnit = 2
)

func (this *TimeClass) GetCurrentTimestamp(unit TimeUnit) int64 {
	if unit == TimeUnit_SECOND {
		return time.Now().Unix()
	} else if unit == TimeUnit_MILLISECOND {
		return int64(time.Now().Nanosecond() / 1000)
	} else {
		panic(errors.New(`unit error`))
	}
}

func (this *TimeClass) GetFormatTimeFromTimeObj(time time.Time, format string) string {
	layout := this.getLayoutFromFormat(format)
	return time.Format(layout)
}

func (this *TimeClass) GetFormatTimeFromTimestamp(timestamp int64, format string, utc bool) string {
	tm := time.Unix(timestamp, 0) // 默认是local时间
	if utc {
		tm = tm.UTC()
	}
	return this.GetFormatTimeFromTimeObj(tm, format)
}

func (this *TimeClass) getLayoutFromFormat(format string) string {
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

func (this *TimeClass) GetLocalTimeFromLocalStr(str string) time.Time {
	t, err := time.ParseInLocation(`2006-01-02T15:04:05+08:00`, str, time.Local)
	if err != nil {
		panic(err)
	}
	return t.Local()
}

func (this *TimeClass) GetUtcTimeFromLocalStr(str string) time.Time {
	t, err := time.ParseInLocation(`2006-01-02T15:04:05+08:00`, str, time.Local)
	if err != nil {
		panic(err)
	}
	return t.UTC()
}

func (this *TimeClass) GetUtcTimeFromOffsetStr(str string, offsetHours int, format string) time.Time {
	t, err := time.ParseInLocation(this.getLayoutFromFormat(format), str, time.FixedZone("CST", offsetHours * 3600))
	if err != nil {
		panic(err)
	}
	return t.UTC()
}

func (this *TimeClass) GetUtcTimeFromIsoDateStr(str string, loc *time.Location) time.Time {
	t, err := time.ParseInLocation(`2006-01-02T15:04:05Z`, str, loc)
	if err != nil {
		panic(err)
	}
	return t.UTC()
}

func (this *TimeClass) GetLocalTimeFromIsoDateStr(str string, loc *time.Location) time.Time {
	t, err := time.ParseInLocation(`2006-01-02T15:04:05Z`, str, loc)
	if err != nil {
		panic(err)
	}
	return t.Local()
}

func (this *TimeClass) GetLocalBeginTimeOfToday() time.Time {
	t := time.Now()
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func (this *TimeClass) GetLocalEndTimeOfToday() time.Time {
	t := time.Now()
	year, month, day := t.Date()
	return time.Date(year, month, day+1, 0, 0, 0, 0, t.Location())
}
