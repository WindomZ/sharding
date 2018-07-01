package sharding

import (
	"fmt"
	"hash/crc32"
)

type hashShard struct {
	Separator string
	Offset    uint32
	Range     uint32
}

func (f *hashShard) Update() {
}

func (f hashShard) Format(s ...string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return s[0]
	}
	return fmt.Sprintf("%s%s%d", s[0], f.Separator,
		f.Offset+crc32.ChecksumIEEE([]byte(s[1]))%f.Range)
}

func NewHashShard(s string, o, r uint32) Shard {
	ret := &hashShard{
		Separator: s,
	}
	if o < 0 {
		ret.Offset = 0
	} else {
		ret.Offset = uint32(o)
	}
	if r < 0 {
		ret.Range = 0
	} else {
		ret.Range = uint32(r)
	}
	return ret
}
