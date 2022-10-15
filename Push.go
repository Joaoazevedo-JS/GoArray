package GoArray

func (array *Array[T]) Push(values ...T) int {
	array.Values = append(array.Values, values...)

	return array.Length()
}
