package sharding

import (
	"fmt"
	"time"
)

type timeShard struct {
	Separator string
	Layout    string
	Time      time.Time
}

func (f *timeShard) Update() {
	f.Time = time.Now()
}

func (f timeShard) Format(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	return fmt.Sprintf("%s%s%s", s[0], f.Separator,
		f.Time.Format(f.Layout))
}

func NewTimeShard(s, l string) Shard {
	ret := &timeShard{
		Separator: s,
		Layout:    l,
		Time:      time.Now(),
	}
	return ret
}
