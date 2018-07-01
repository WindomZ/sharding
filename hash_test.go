package sharding

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

var testHash Shard

func TestNewHashShard(t *testing.T) {
	testHash = NewHashShard("_", 0, 64)
}

func TestHashShard_Update(t *testing.T) {
	testHash.Update()
}

func TestHashShard_Format(t *testing.T) {
	assert.Equal(t, "hello_3", testHash.Format("hello", "world"))
	assert.Equal(t, "hello_35", testHash.Format("hello", "golang"))
}
