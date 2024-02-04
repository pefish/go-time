package go_time

import (
	"fmt"
	"github.com/pefish/go-test"
	"testing"
	"time"
)

func TestTimeUtil_CurrentTimestamp(t *testing.T) {
	a := TimeInstance.CurrentTimestamp()
	go_test_.Equal(t, true, a > 1694078358311)
}

func TestTimeType_TimestampToTime(t *testing.T) {
	result := TimeInstance.TimestampToTime(1698036579000, true)
	go_test_.Equal(t, "2023-10-23 04:49:39", TimeInstance.TimeToStr(result, "0000-00-00 00:00:00"))

	result1 := TimeInstance.TimestampToTime(1698036579000, false)
	go_test_.Equal(t, "2023-10-23 12:49:39", TimeInstance.TimeToStr(result1, "0000-00-00 00:00:00"))
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

	timestamp2, err := TimeInstance.StrToTimestamp("2023-12-09 08:30:00 +0000 UTC", true)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "2023-12-09 16:30:00", TimeInstance.TimestampToStr(timestamp2, "0000-00-00 00:00:00", false))

	timestamp3, err := TimeInstance.StrToTimestamp("2024-01-05 09:55:01.631840873 +0000 UTC", true)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "2024-01-05 17:55:01", TimeInstance.TimestampToStr(timestamp3, "0000-00-00 00:00:00", false))

}

func TestTimeType_StrToTime(t *testing.T) {

}

func TestTimeType_BeginOfTime(t *testing.T) {
	time_ := TimeInstance.BeginOfTime(time.Date(2023, 5, 12, 11, 11, 11, 0, time.Local), true)
	str := TimeInstance.TimeToStr(time_, "0000-00-00 00:00:00")
	go_test_.Equal(t, "2023-05-11 16:00:00", str)
}
