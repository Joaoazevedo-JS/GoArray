package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.length should return the length of the array
func TestLength(t *testing.T) {
	array := New[int]()

	array.Push(1, 2, 3)

	assert.Equal(t, 3, array.Length(), "length should be 3")
}

// test array.length should return int
func TestLengthType(t *testing.T) {
	array := New[int]()

	array.Push(1, 2, 3)

	assert.IsType(t, 0, array.Length(), "length should be int")
}
