package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.concat method should concat arrays
func TestConcat(t *testing.T) {
	array := New[int]()
	array2 := New[int]()

	array.Push(1, 2, 3, 4, 5)
	array2.Push(6, 7, 8, 9, 10)

	concat := array.Concat(array2)

	first, _ := concat.Get(0)

	assert.Equal(t, 10, concat.Length(), "the length of the concat array should be the sum of the lengths of the original arrays and concat array")
	assert.Equal(t, 1, first, "the first element of the concat array should be the first element of array that used concat method")
}

// test array.concat method shouldn't mutate the original array
func TestConcatNotMutate(t *testing.T) {
	array := New[int]()
	array2 := New[int]()

	array.Push(1, 2, 3, 4, 5)
	array2.Push(6, 7, 8, 9, 10)

	array.Concat(array2)

	assert.Equal(t, 5, array.Length(), "should not mutate the original array")
	assert.Equal(t, 5, array2.Length(), "should not mutate the original array")
}
