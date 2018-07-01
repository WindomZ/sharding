package sharding

type Shard interface {
	Update()
	Format(s ...string) string
}
