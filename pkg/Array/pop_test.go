package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.pop should remove the last item from the array
func TestPop(t *testing.T) {
	array := New[int]()

	array.Push(1, 2, 3)

	_, err := array.Pop()

	assert.Nil(t, err, "array.pop() should not return error")
	assert.Equal(t, 2, array.Length(), "the length after pop should be 2")
	assert.Equal(t, 2, array.Values[array.Length()-1], "the last item should be 2")
}

// test array.pop should return the that was removed
func TestPopReturn(t *testing.T) {
	array := New[int]()

	array.Push(1, 2, 3)

	item, _ := array.Pop()

	assert.Equal(t, 3, item, "array.pop() should return the item that was removed")
}

// test array.pop should return error if array is empty
func TestPopEmpty(t *testing.T) {
	array := New[int]()

	_, err := array.Pop()

	assert.Error(t, err, "array.pop() should return error if array is empty")
}
