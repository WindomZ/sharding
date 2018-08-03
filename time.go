package sharding

import (
	"fmt"
	"time"
)

type timeShard struct {
	formatter
	Layout string
	Time   time.Time
}

func (t *timeShard) Update() {
	t.Time = time.Now()
}

func (t timeShard) Format(s ...string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) >= 2 {
		return t.formatter.Format(s...)
	}
	return fmt.Sprintf("%s%s%s", s[0], t.Separator,
		t.Time.Format(t.Layout))
}

func (t timeShard) ShiftFormat(i int64, s ...string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) >= 2 {
		return t.formatter.Format(s...)
	}
	return fmt.Sprintf("%s%s%s", s[0], t.Separator,
		t.Time.Add(time.Duration(i)).Format(t.Layout))
}

func NewTimeShard(sep, layout string) Shard {
	ret := &timeShard{
		formatter: formatter{
			Separator: sep,
		},
		Layout: layout,
		Time:   time.Now(),
	}
	return ret
}
