package cache

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	testChache := NewLRUCache(3)

	testChache.Put("key1", "val1")
	testChache.Put("key2", "val2")
	testChache.Put("key3", "val3")

	// All 3 vals should exist
	val1, _ := testChache.Get("key1")
	val2, _ := testChache.Get("key2")
	val3, _ := testChache.Get("key3")
	assert.Equal(t, "val1", val1)
	assert.Equal(t, "val2", val2)
	assert.Equal(t, "val3", val3)

	testChache.Put("key4", "val4")

	val4, _ := testChache.Get("key4")
	assert.Equal(t, "val4", val4)

	// val1 should be evicted now
	_, val1Exists := testChache.Get("val1")
	assert.False(t, val1Exists)
}