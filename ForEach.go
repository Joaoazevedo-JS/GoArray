package Array

func (array *Array[T]) ForEach(callback func(currentValue T, index int, self *Array[T])) {
	for index, value := range array.Values {
		callback(value, index, array)
	}
}
