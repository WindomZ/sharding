package sharding

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

var testHash Shard

func TestNewHashShard(t *testing.T) {
	testHash = NewHashShard("_", 0, 0)
	assert.Equal(t, "a_0", testHash.Format("a", "b"))
	assert.Equal(t, "a_0", testHash.Format("a", "c"))

	testHash = NewHashShard("_", 0, 64)
}

func TestHashShard_Update(t *testing.T) {
	testHash.Update()
}

func TestHashShard_Format(t *testing.T) {
	assert.Equal(t, "", testHash.Format())
	assert.Equal(t, "hello", testHash.Format("hello"))

	assert.Equal(t, "hello_3", testHash.Format("hello", "world"))
	assert.Equal(t, "hello_35", testHash.Format("hello", "golang"))

	assert.Equal(t, "hello_golang_world", testHash.Format("hello", "golang", "world"))
}

func TestHashShard_ShiftFormat(t *testing.T) {
	assert.Equal(t, "", testHash.ShiftFormat(0))
	assert.Equal(t, "hello", testHash.ShiftFormat(0, "hello"))

	assert.Equal(t, "hello_2", testHash.ShiftFormat(-1, "hello", "world"))
	assert.Equal(t, "hello_36", testHash.ShiftFormat(1, "hello", "golang"))

	assert.Equal(t, "hello_golang_world", testHash.ShiftFormat(100, "hello", "golang", "world"))
}
