package sharding

import (
	"fmt"
	"hash/crc32"
)

type hashShard struct {
	formatter
	Offset uint32
	Range  uint32
}

func (h *hashShard) Update() {
}

func (h hashShard) Format(s ...string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return s[0]
	} else if len(s) >= 3 {
		return h.formatter.Format(s...)
	}
	return fmt.Sprintf("%s%s%d", s[0], h.Separator,
		h.Offset+crc32.ChecksumIEEE([]byte(s[1]))%h.Range)
}

func NewHashShard(sep string, offset, limit uint32) Shard {
	ret := &hashShard{
		formatter: formatter{
			Separator: sep,
		},
	}
	ret.Offset = uint32(offset)
	if limit <= 0 {
		ret.Range = 1
	} else {
		ret.Range = uint32(limit)
	}
	return ret
}
