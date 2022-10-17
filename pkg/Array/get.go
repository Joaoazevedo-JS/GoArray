package Array

import "errors"

func (array *Array[T]) Get(index int) (T, error) {
	if index < 0 || index >= array.Length() {
		return nil, errors.New("index out of bounds")
	}

	return array.Values[index], nil
}
