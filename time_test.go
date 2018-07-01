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
	assert.Equal(t, fmt.Sprintf("hello_%s",
		time.Now().Format("20060102")), testTime.Format("hello"))
}
