package p_time

import (
	"testing"
)

func TestTimeClass_GetUtcTimeFromOffsetStr(t *testing.T) {
	str := Time.GetUtcTimeFromOffsetStr(`2019-11-11 00:00:00`, 8, `0000-00-00 00:00:00`).Format(`2006-01-02 15:04:05`)
	if str != `2019-11-10 16:00:00` {
		t.Error()
	}

	str1 := Time.GetUtcTimeFromOffsetStr(`2019-11-11 00:00:00`, 0, `0000-00-00 00:00:00`).Format(`2006-01-02 15:04:05`)
	if str1 != `2019-11-11 00:00:00` {
		t.Error()
	}
}
