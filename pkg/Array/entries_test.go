package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.entries should return an array of index and value pairs
func TestEntries(t *testing.T) {
	array := New[string]("lorem", "ipsum", "dolor")

	entries := array.Entries()

	entry01, _ := entries.Get(0)
	entry02, _ := entries.Get(1)
	entry03, _ := entries.Get(2)

	assert.Equal(t, 3, entries.Length(), "should return an array of index and value pairs")
	assert.EqualValues(t, []interface{}{0, "lorem"}, entry01, "should return an array of index and value pairs")
	assert.EqualValues(t, []interface{}{1, "ipsum"}, entry02, "should return an array of index and value pairs")
	assert.EqualValues(t, []interface{}{2, "dolor"}, entry03, "should return an array of index and value pairs")
}

// test array.entries should return an empty array if the array is empty
func TestEntriesEmpty(t *testing.T) {
	array := New[string]()

	entries := array.Entries()

	assert.Equal(t, 0, entries.Length(), "should return an empty array if the array is empty")
}
