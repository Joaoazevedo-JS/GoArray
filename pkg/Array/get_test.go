package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.get should return the value at the given index
func TestGet(t *testing.T) {
	array := New[int](1, 2, 3)

	value, err := array.Get(1)

	assert.Equal(t, 2, value, "should return the value at the given index")
	assert.Nil(t, err, "should not return an error")
}

// test array.get should return undefined if the index is out of bounds
func TestGetOutOfBounds(t *testing.T) {
	array := New[int](1, 2, 3)

	value, err := array.Get(4)

	assert.Equal(t, nil, value, "should return undefined if the index is out of bounds")
	assert.NotNil(t, err, "should return an error if the index is out of bounds")
}
