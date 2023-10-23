package go_time

import (
	"fmt"
	"github.com/pefish/go-test"
	"testing"
)

func TestTimeClass_GetUtcTimeFromOffsetStr(t *testing.T) {
	str := TimeInstance.MustOffsetStrToLocalTime(`0000-00-00 00:00:00`, `2019-11-11 00:00:00`, 8).Format(`2006-01-02 15:04:05`)
	if str != `2019-11-11 00:00:00` {
		t.Error()
	}
}

func TestTimeUtil_CurrentTimestamp(t *testing.T) {
	a := TimeInstance.CurrentTimestamp(TimeUnit_MILLISECOND)
	test.Equal(t, true, a > 1694078358311)
}

func TestTimeType_TimestampToTime(t *testing.T) {
	result := TimeInstance.TimestampToTime(TimeInstance.CurrentTimestamp(TimeUnit_SECOND), true)
	fmt.Println(TimeInstance.TimeToStr(result, "0000-00-00 00:00:00"))
}

func TestTimeType_LocalStrToLocalTime(t *testing.T) {
	ti, err := TimeInstance.LocalStrToLocalTime("2023-10-23 02:57:13", "0000-00-00 00:00:00")
	if err != nil {
		t.Error(err)
	}
	str := TimeInstance.TimeToStr(ti, "0000-00-00 00:00:00")
	fmt.Println(str)
}

func TestTimeType_UtcStrToLocalTime(t *testing.T) {
	ti, err := TimeInstance.UtcStrToLocalTime("2023-10-23 02:57:13", "0000-00-00 00:00:00")
	if err != nil {
		t.Error(err)
	}
	str := TimeInstance.TimeToStr(ti, "0000-00-00 00:00:00")
	fmt.Println(str)
}

func TestTimeType_NowToUtcStr(t *testing.T) {
	fmt.Println(TimeInstance.NowToUtcStr())
}

func TestTimeType_StrToTimestamp(t *testing.T) {
	timestamp, err := TimeInstance.StrToTimestamp("2023-10-23 11:50:46", "0000-00-00 00:00:00", false)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(timestamp)
	fmt.Println(TimeInstance.TimestampToStr(timestamp, "0000-00-00 00:00:00", false))
}
