package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.findIndex method should return the index that satisfies the provided testing function
func TestFindIndex(t *testing.T) {
	array := New[int]()
	array.Push(1, 2, 3, 4, 5)

	index := array.FindIndex(func(currentValue int, index int, self *Array[int]) bool {
		return currentValue == 3
	})

	assert.Equal(t, 2, index, "should return the index that satisfies the provided testing function")
	assert.Equal(t, 5, array.Length(), "should not mutate the original array")
}

// test array.findIndex method should return -1 if no element satisfies the provided testing function
func TestFindIndexNotFound(t *testing.T) {
	array := New[int]()
	array.Push(1, 2, 3, 4, 5)

	index := array.FindIndex(func(currentValue int, index int, self *Array[int]) bool {
		return currentValue == 6
	})

	assert.Equal(t, -1, index, "should return -1 if no element satisfies the provided testing function")
	assert.Equal(t, 5, array.Length(), "should not mutate the original array")
}

// test array.findIndex method should return the index of the first element that satisfies the provided testing function
func TestFindIndexFirst(t *testing.T) {
	array := New[int]()
	array.Push(1, 2, 3, 4, 5)

	index := array.FindIndex(func(currentValue int, index int, self *Array[int]) bool {
		return currentValue > 2
	})

	assert.Equal(t, 2, index, "should return the index of the first element that satisfies the provided testing function")
	assert.Equal(t, 5, array.Length(), "should not mutate the original array")
}
