package sparse_test

import (
	"testing"

	. "."
)

var array = NewArray()

func TestArraySet(t *testing.T) {
	for i := 0; i < 100; i++ {
		array.Set('a', 1)
	}
}
