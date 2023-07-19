package go_time

import (
	"testing"
)

func TestTimeClass_GetUtcTimeFromOffsetStr(t *testing.T) {
	str := TimeUtilInstance.MustOffsetStrToLocal(`0000-00-00 00:00:00`, `2019-11-11 00:00:00`, 8).Format(`2006-01-02 15:04:05`)
	if str != `2019-11-11 00:00:00` {
		t.Error()
	}
}
