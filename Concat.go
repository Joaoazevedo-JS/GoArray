package GoArray

func (array *Array[T]) Concat(arrays ...Array[T]) Array[T] {
	concat := New[T]()

	concat.Push(array.Values...)

	for _, value := range arrays {
		concat.Push(value.Values...)
	}

	return concat
}
