package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test new array without arguments
func TestNewArrayWithoutArguments(t *testing.T) {
	array := New[int]()

	assert.Equal(t, 0, array.Length(), "array length should be 0")
}

// test new array with many arguments
func TestNewArrayWithManyArguments(t *testing.T) {
	array := New[int](1, 2, 3)

	assert.Equal(t, 3, array.Length(), "array length should be 3")
	assert.Equal(t, 1, array.Values[0], "array first value should be 1")
	assert.IsType(t, []int{0}, array.Values, "array values should be []int")
}
