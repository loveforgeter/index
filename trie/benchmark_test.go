package trie_test

import (
	"math/rand"
	"testing"

	. "."
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func BenchmarkTrie(b *testing.B) {
	t := NewTrie()
	for i := 0; i < b.N; i++ {
		t.Insert(randomKey(), 0)
	}
}

func BenchmarkIndexMap(b *testing.B) {
	t := NewIndexTree()
	for i := 0; i < b.N; i++ {
		t.Insert(randomKey(), 0)
	}
}

// randomKey 随机生成长度小于或等于10的字符串(0-9,a-z,A-Z)
func randomKey() string {
	strLen := rand.Intn(10) + 1
	runes := make([]rune, strLen)
	for i := 0; i < strLen; i++ {
		runes[i] = letters[rand.Intn(len(letters))]
	}
	return string(runes)
}
