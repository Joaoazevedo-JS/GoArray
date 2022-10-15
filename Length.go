package GoArray

func (array *Array[T]) Length() int {
	return len(array.Values)
}
