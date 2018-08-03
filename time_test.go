package sharding

import (
	"fmt"
	"testing"
	"time"

	"github.com/WindomZ/testify/assert"
)

var testTime Shard

func TestNewTimeShard(t *testing.T) {
	testTime = NewTimeShard("_", "20060102")
}

func TestTimeShard_Update(t *testing.T) {
	testTime.Update()
}

func TestTimeShard_Format(t *testing.T) {
	assert.Equal(t, "", testTime.Format())

	assert.Equal(t, fmt.Sprintf("hello_%s",
		time.Now().Format("20060102")),
		testTime.Format("hello"),
	)

	assert.Equal(t, "hello_world",
		testTime.Format("hello", "world"))
}

func TestTimeShard_ShiftFormat(t *testing.T) {
	assert.Equal(t, "", testTime.ShiftFormat(0))

	assert.Equal(t, fmt.Sprintf("hello_%s",
		time.Now().Add(-time.Hour*24).Format("20060102")),
		testTime.ShiftFormat(-int64(time.Hour*24), "hello"),
	)
	assert.Equal(t, fmt.Sprintf("hello_%s",
		time.Now().Add(time.Hour*48).Format("20060102")),
		testTime.ShiftFormat(int64(time.Hour*48), "hello"),
	)

	assert.Equal(t, "hello_world",
		testTime.ShiftFormat(100, "hello", "world"))
}
