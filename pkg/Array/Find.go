package Array

import "errors"

func (array *Array[T]) Find(callback func(currentValue T, index int, self *Array[T]) bool) (T, error) {
	var undefined T

	for index, value := range array.Values {
		found := callback(value, index, array)

		if found {
			return value, nil
		}
	}

	return undefined, errors.New("not Found")
}
