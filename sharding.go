package sharding

import "strings"

type Shard interface {
	Update()
	Format(s ...string) string
	ShiftFormat(i int64, s ...string) string
}

type formatter struct {
	Separator string
}

func (f formatter) Format(a ...string) string {
	return strings.Join(a, f.Separator)
}
