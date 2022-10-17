package Array

import "errors"

func (array *Array[T]) Get(index int) (T, error) {
	var undefined T

	if index < 0 || index >= array.Length() {
		return undefined, errors.New("index out of range")
	}

	return array.Values[index], nil
}
