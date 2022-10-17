package Array

import "errors"

func (array *Array[T]) Pop() (T, error) {
	var item T

	if array.Length() == 0 {
		return item, errors.New("Array is empty")
	}

	item = array.Values[array.Length()-1]

	// this will fail if length is zero
	array.Values = array.Values[:array.Length()-1]

	return item, nil
}
