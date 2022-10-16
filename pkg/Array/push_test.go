package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.push method without arguments
func TestPushWithoutArguments(t *testing.T) {
	array := New[int]()

	array.Push()

	assert.Equal(t, 0, array.Length(), "array should be empty")
}

// test array.push method with many int arguments
func TestPushWithManyIntArguments(t *testing.T) {
	array := New[int]()

	array.Push(1, 2, 3, 4, 5)

	assert.Equal(t, 5, array.Length(), "array should have 5 items")
	assert.Equal(t, 1, array.Values[0], "first item should be 1")
	assert.Equal(t, 5, array.Values[4], "last item should be 5")
}

// test array.push method with many string arguments
func TestPushWithManyStringArguments(t *testing.T) {
	array := New[string]()

	array.Push("one", "two", "three", "four", "five")

	assert.Equal(t, 5, array.Length(), "array should have 5 items")
	assert.Equal(t, "one", array.Values[0], "first item should be one")
	assert.Equal(t, "five", array.Values[4], "last item should be five")
}

type PushedItems struct {
	Id    int
	Value string
}

// test array.push method with many struct arguments
func TestPushWithManyStructArguments(t *testing.T) {
	array := New[PushedItems]()

	array.Push(
		PushedItems{1, "one"},
		PushedItems{2, "two"},
		PushedItems{3, "three"},
		PushedItems{4, "four"},
		PushedItems{5, "five"},
	)

	assert.Equal(t, 5, array.Length(), "array should have 5 items")
	assert.Equal(t, 1, array.Values[0].Id, "first item should have id equal 1")
	assert.Equal(t, "five", array.Values[4].Value, "last item should have value equal five")
}

// test array.push method should return new length of array
func TestPushShouldReturnNewLengthOfArray(t *testing.T) {
	array := New[int]()

	length := array.Push(1, 2, 3, 4, 5)

	assert.Equal(t, 5, length, "array should have 5 items")
}
