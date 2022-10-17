package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.find method should return the element that satisfies the provided testing function
func TestFind(t *testing.T) {
	array := New[string]("one", "two", "three", "four", "five")

	result, err := array.Find(func(currentValue string, index int, array *Array[string]) bool {
		return index == 3
	})

	assert.Equal(t, "four", result, "should return the element that satisfies the provided testing function")
	assert.Nil(t, err, "should not return an error")
}

// test array.find method should return only the first element that satisfies the provided testing function
func TestFindFirst(t *testing.T) {
	array := New[string]("one", "two", "three", "four", "five")

	result, err := array.Find(func(currentValue string, index int, array *Array[string]) bool {
		return index > 2
	})

	assert.Equal(t, "four", result, "should return the only the first element that satisfies the provided testing function")
	assert.Nil(t, err, "should not return an error")
}

// test array.find method should return an error if no element satisfies the provided testing function
func TestFindNotFound(t *testing.T) {
	array := New[string]("one", "two", "three", "four", "five")

	_, err := array.Find(func(currentValue string, index int, array *Array[string]) bool {
		return index == 6
	})

	assert.NotNil(t, err, "should return an error if no element satisfies the provided testing function")
}
