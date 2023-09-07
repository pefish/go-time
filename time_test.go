package go_time

import (
	"github.com/pefish/go-test"
	"testing"
)

func TestTimeClass_GetUtcTimeFromOffsetStr(t *testing.T) {
	str := TimeInstance.MustOffsetStrToLocal(`0000-00-00 00:00:00`, `2019-11-11 00:00:00`, 8).Format(`2006-01-02 15:04:05`)
	if str != `2019-11-11 00:00:00` {
		t.Error()
	}
}

func TestTimeUtil_CurrentTimestamp(t *testing.T) {
	a := TimeInstance.CurrentTimestamp(TimeUnit_MILLISECOND)
	test.Equal(t, true, a > 1694078358311)
}
