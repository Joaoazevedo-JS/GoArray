package Array

func (array *Array[T]) FindIndex(callback func(currentValue T, index int, self *Array[T]) bool) int {
	for index, value := range array.Values {
		found := callback(value, index, array)

		if found {
			return index
		}
	}

	return -1
}
