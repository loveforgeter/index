package index_test

import (
	"math/rand"
	"testing"
)

var m = make(map[int]int)
var s = make([]int, 10)

func init() {
	for i := 0; i < 10; i++ {
		m[i] = i
		s[i] = i
	}
}

func Benchmark_Array(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find(s, rand.Intn(100))
	}
}

func Benchmark_Map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := m[i]
		if 0 == x {
		}
	}
}

func find(s []int, v int) {
	for _, value := range s {
		if value == v {
			return
		}
	}
}
