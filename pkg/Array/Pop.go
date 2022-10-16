package Array

func (array *Array[T]) Pop() T {
	var item T

	if array.Length() == 0 {
		return item
	}

	item = array.Values[array.Length()-1]

	// this will fail if length is zero
	array.Values = array.Values[:array.Length()-1]

	return item
}
