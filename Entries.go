package GoArray

type Entry []interface{}

func (array *Array[T]) Entries() Array[Entry] {
	entries := New[Entry]()

	array.ForEach(func(currentValue T, index int, _ *Array[T]) {
		entries.Push(Entry{
			index,
			currentValue,
		})
	})

	return entries
}
