package GoArray

type Array[T any] struct {
	Values []T
}

func New[T any](values ...T) Array[T] {
	array := Array[T]{
		Values: []T{},
	}

	array.Push(values...)

	return array
}
