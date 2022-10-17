package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.unshift should add elements to the beginning of an array
func TestUnshift(t *testing.T) {
	array := New[int]()

	array.Push(1, 2, 3)

	array.Unshift(4, 5, 6)

	assert.Equal(t, 6, array.Length(), "the length after unshift should be 6")
	assert.Equal(t, 4, array.Values[0], "the first item should be 4")
	assert.Equal(t, 3, array.Values[array.Length()-1], "the last item should be 3")
}

// test array.unshift should not add elements if no arguments are passed
func TestUnshiftWithoutArguments(t *testing.T) {
	array := New[int]()

	array.Push(1, 2, 3)

	array.Unshift()

	assert.Equal(t, 3, array.Length(), "the length after unshift should be 3")
	assert.Equal(t, 1, array.Values[0], "the first item should be 1")
	assert.Equal(t, 3, array.Values[array.Length()-1], "the last item should be 3")
}

// test array.unshift should return the new length of the array
func TestUnshiftReturn(t *testing.T) {
	array := New[int]()

	array.Push(1, 2, 3)

	length := array.Unshift(4, 5, 6)

	assert.Equal(t, 6, length, "array.unshift() should return the new length of the array")
}
