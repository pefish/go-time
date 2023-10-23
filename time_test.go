package go_time

import (
	"fmt"
	"github.com/pefish/go-test"
	"testing"
)

func TestTimeClass_GetUtcTimeFromOffsetStr(t *testing.T) {
	str := TimeInstance.MustOffsetStrToLocalTime(`2019-11-11 00:00:00`, 8).Format(`2006-01-02 15:04:05`)

	go_test_.Equal(t, `2019-11-11 00:00:00`, str)
}

func TestTimeUtil_CurrentTimestamp(t *testing.T) {
	a := TimeInstance.CurrentTimestamp(TimeUnit_MILLISECOND)
	go_test_.Equal(t, true, a > 1694078358311)
}

func TestTimeType_TimestampToTime(t *testing.T) {
	result := TimeInstance.TimestampToTime(1698036579, true)
	go_test_.Equal(t, "2023-10-23 04:49:39", TimeInstance.TimeToStr(result, "0000-00-00 00:00:00"))

	result1 := TimeInstance.TimestampToTime(1698036579, false)
	go_test_.Equal(t, "2023-10-23 12:49:39", TimeInstance.TimeToStr(result1, "0000-00-00 00:00:00"))
}

func TestTimeType_LocalStrToLocalTime(t *testing.T) {
	ti, err := TimeInstance.LocalStrToLocalTime("2023-10-23 02:57:13")
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "2023-10-23 02:57:13", TimeInstance.TimeToStr(ti, "0000-00-00 00:00:00"))
}

func TestTimeType_UtcStrToLocalTime(t *testing.T) {
	ti, err := TimeInstance.UtcStrToLocalTime("2023-10-23 02:57:13")
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "2023-10-23 10:57:13", TimeInstance.TimeToStr(ti, "0000-00-00 00:00:00"))
}

func TestTimeType_NowToUtcStr(t *testing.T) {
	fmt.Println(TimeInstance.NowToUtcStr())
}

func TestTimeType_StrToTimestamp(t *testing.T) {
	timestamp, err := TimeInstance.StrToTimestamp("2023-10-23 11:50:46", false)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "2023-10-23 11:50:46", TimeInstance.TimestampToStr(timestamp, "0000-00-00 00:00:00", false))

	timestamp1, err := TimeInstance.StrToTimestamp("2023-10-23T11:50:46Z", false)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "2023-10-23 11:50:46", TimeInstance.TimestampToStr(timestamp1, "0000-00-00 00:00:00", false))
	go_test_.Equal(t, "2023-10-23T11:50:46Z", TimeInstance.TimestampToStr(timestamp1, "0000-00-00T00:00:00Z", false))
}
