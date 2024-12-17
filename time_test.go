package go_time

import (
	"fmt"
	"testing"
	"time"

	go_test_ "github.com/pefish/go-test"
)

func TestTimeUtil_CurrentTimestamp(t *testing.T) {
	a := CurrentTimestamp()
	go_test_.Equal(t, true, a > 1694078358311)
}

func TestTimeType_TimestampToTime(t *testing.T) {
	result := TimestampToTime(1698036579000, true)
	go_test_.Equal(t, "2023-10-23 04:49:39", TimeToStr(result, "0000-00-00 00:00:00"))

	result1 := TimestampToTime(1698036579000, false)
	go_test_.Equal(t, "2023-10-23 12:49:39", TimeToStr(result1, "0000-00-00 00:00:00"))
}

func TestTimeType_NowToUtcStr(t *testing.T) {
	fmt.Println(NowToUtcStr())
}

func TestTimeType_StrToTimestamp(t *testing.T) {
	timestamp, err := StrToTimestamp("2023-10-23 11:50:46", false)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "2023-10-23 11:50:46", TimestampToStr(timestamp, "0000-00-00 00:00:00", false, 0))

	timestamp1, err := StrToTimestamp("2023-10-23T11:50:46Z", false)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "2023-10-23 11:50:46", TimestampToStr(timestamp1, "0000-00-00 00:00:00", false, 0))
	go_test_.Equal(t, "2023-10-23T11:50:46Z", TimestampToStr(timestamp1, "0000-00-00T00:00:00Z", false, 0))

	timestamp2, err := StrToTimestamp("2023-12-09 08:30:00 +0000 UTC", true)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "2023-12-09 16:30:00", TimestampToStr(timestamp2, "0000-00-00 00:00:00", false, 0))

	timestamp3, err := StrToTimestamp("2024-01-05 09:55:01.631840873 +0000 UTC", true)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "2024-01-05 17:55:01", TimestampToStr(timestamp3, "0000-00-00 00:00:00", false, 0))

}

func TestTimeType_TimestampToStr(t *testing.T) {
	str := TimestampToStr(1734412133000, "0000-00-00 00:00:00", true, 0)
	go_test_.Equal(t, "2024-12-17 05:08:53", str)

	str1 := TimestampToStr(1734412133000, "0000-00-00 00:00:00", true, 8)
	go_test_.Equal(t, "2024-12-17 13:08:53", str1)
}

func TestTimeType_StrToTime(t *testing.T) {

}

func TestTimeType_BeginOfTime(t *testing.T) {
	time_ := BeginOfTime(time.Date(2023, 5, 12, 11, 11, 11, 0, time.Local), true)
	str := TimeToStr(time_, "0000-00-00 00:00:00")
	go_test_.Equal(t, "2023-05-11 16:00:00", str)
}
