package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test array.toString method with empty array
func TestToStringEmptyArray(t *testing.T) {
	array := New[int]()

	assert.Equal(t, "", array.ToString(), "array.toString() should return an empty string")
}

// test array.toString method with array of strings
func TestToStringArrayOfStrings(t *testing.T) {
	array := New[string]()

	array.Push("foo", "bar", "baz")

	assert.Equal(t, "foo,bar,baz", array.ToString(), "array.toString() should return foo,bar,baz")
}
