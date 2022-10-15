package GoArray

func (array *Array[T]) Filter(callback func(currentValue T, index int, self *Array[T]) bool) Array[T] {
	list := New[T]()

	for index, value := range array.Values {
		add := callback(value, index, array)

		if add {
			list.Push(value)
		}
	}

	return list
}
