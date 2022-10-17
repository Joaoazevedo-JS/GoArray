package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.filter should return a new array with all elements that pass the test implemented by the provided function
func TestFilter(t *testing.T) {
	array := New[int]()
	array.Push(1, 2, 3, 4, 5)

	result := array.Filter(func(currentValue int, index int, self *Array[int]) bool {
		return currentValue > 2
	})

	assert.Equal(t, 3, result.Length(), "should return a new array with all elements that pass the test implemented by the provided function")
	assert.Equal(t, []int{3, 4, 5}, result.Values, "should return a new array with all elements that pass the test implemented by the provided function")
}

// test array.filter shouldn't mutate the original array
func TestFilterNotMutate(t *testing.T) {
	array := New[int]()
	array.Push(1, 2, 3, 4, 5)

	array.Filter(func(currentValue int, index int, self *Array[int]) bool {
		return currentValue > 2
	})

	assert.Equal(t, 5, array.Length(), "should not mutate the original array")
}

// test array.filter should mutate itself inside callback but not the new array
func TestFilterMutate(t *testing.T) {
	array := New[int]()
	array.Push(1, 2, 3, 4, 5)

	result := array.Filter(func(currentValue int, index int, self *Array[int]) bool {
		self.Push(currentValue + 1)

		return currentValue > 2
	})

	assert.Equal(t, 3, result.Length(), "should return a new array with all elements that pass the test implemented by the provided function")
	assert.Equal(t, []int{3, 4, 5}, result.Values, "should return a new array with all elements that pass the test implemented by the provided function")
	assert.Equal(t, 10, array.Length(), "should mutate itself inside callback but not the new array")
	assert.Equal(t, []int{1, 2, 3, 4, 5, 2, 3, 4, 5, 6}, array.Values, "should mutate itself inside callback but not the new array")
}

// test array.filter should return an empty array if no elements pass the test implemented by the provided function
func TestFilterEmpty(t *testing.T) {
	array := New[int]()
	array.Push(1, 2, 3, 4, 5)

	result := array.Filter(func(currentValue int, index int, self *Array[int]) bool {
		return currentValue > 5
	})

	assert.Equal(t, 0, result.Length(), "should return an empty array if no elements pass the test implemented by the provided function")
	assert.Empty(t, result.Values, "should return an empty array if no elements pass the test implemented by the provided function")
}
