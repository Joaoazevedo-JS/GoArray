package GoArray

func (array *Array[T]) Unshift(values ...T) int {
	list := New[T]()

	list.Push(values...)
	list.Push(array.Values...)

	array.Values = list.Values

	return array.Length()
}
