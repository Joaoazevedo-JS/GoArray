package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.join method with empty array
func TestJoinWithEmptyArray(t *testing.T) {
	array := New[int]()

	assert.Equal(t, "", array.Join(), "array.join() should return an empty string")
}

// test array.join method should return string
func TestJoinType(t *testing.T) {
	array := New[int]()

	array.Push(1, 2, 3)

	assert.IsType(t, "", array.Join(), "array.join() should return string")
}

// test array.join method without arguments
func TestJoinWithoutSeparator(t *testing.T) {
	array := New[string]()

	array.Push("foo", "bar", "baz")

	assert.Equal(t, "foo,bar,baz", array.Join(), "array.join() should return foo,bar,baz")
}

// test array.join method with separator *
func TestJoinWithSeparator(t *testing.T) {
	array := New[int]()

	array.Push(1, 2, 3)

	assert.Equal(t, "1 * 2 * 3", array.Join(" * "), "array.join() should return 1 * 2 * 3")
}

// test array.join method with many separators
func TestJoinWithManySeparators(t *testing.T) {
	array := New[int]()

	array.Push(1, 2, 3)

	assert.Equal(t, "1 * 2 * 3", array.Join(" * ", " / ", " ' "), "array.join() should return 1 * 2 * 3")
}
