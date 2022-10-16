package Array

import (
	"fmt"
	"strings"
)

func (array *Array[T]) Join(separators ...string) string {
	separator := ","
	texts := New[string]()

	if separators != nil {
		separator = separators[0]
	}

	array.ForEach(func(currentValue T, index int, self *Array[T]) {
		texts.Push(fmt.Sprint(currentValue))
	})

	return strings.Join(texts.Values, separator)
}
