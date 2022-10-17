package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.forEach method should call callback for each element
func TestForEach(t *testing.T) {
	array := New[int](1, 2, 3, 4, 5)
	sum := 0

	array.ForEach(func(value int, index int, self *Array[int]) {
		sum += value
	})

	assert.Equal(t, sum, 15, "The sum of the array should be 15")
}

// test array.forEach method should have access to the current index
func TestForEachIndex(t *testing.T) {
	array := New[int](1, 2, 3, 4, 5)
	sum := 0

	array.ForEach(func(value int, index int, self *Array[int]) {
		sum += index
	})

	assert.Equal(t, sum, 10, "The sum of the array indexes should be 10")
}

// test array.forEach method should have access to the array itself
func TestForEachSelf(t *testing.T) {
	array := New[int](1, 2, 3, 4, 5)
	sum := 0

	array.ForEach(func(value int, index int, self *Array[int]) {
		sum += self.Length()
	})

	assert.Equal(t, sum, 25, "The sum of the array length for each element should be 25")
}

// test array.forEach method should mutate the array itself
func TestForEachMutate(t *testing.T) {
	array := New[int](1, 2, 3, 4, 5)

	array.ForEach(func(value int, index int, self *Array[int]) {
		self.Values[index] = value * 2
	})

	assert.Equal(t, array.Values, []int{2, 4, 6, 8, 10}, "The array should be doubled")
}

// test array.forEach method shouldn't break loop by calling return
func TestForEachReturn(t *testing.T) {
	array := New[int](1, 2, 3, 4, 5)
	sum := 0

	array.ForEach(func(value int, index int, self *Array[int]) {
		sum += value

		if value == 2 {
			return
		}
	})

	assert.Equal(t, sum, 15, "array.forEach should not break loop by calling return")
}
