package trie_test

import (
	"testing"

	. "."
)

var Trie = NewTrie()
var tests []string = []string{
	"a",
	"abc",
	"b",
	"cdefg",
}

func TestInsert(t *testing.T) {
	for i, str := range tests {
		Trie.Insert(str, i)
	}
	t.Log("Test Insert OK")
}

func TestHas(t *testing.T) {
	for _, str := range tests {
		if !Trie.Has(str) {
			t.Error("Key:", str, " not found!")
			return
		}
	}
	if Trie.Has("123") {
		t.Error("Contains other item")
		return
	}
	t.Log("Test Has OK")
}

func TestValue(t *testing.T) {
	for i, str := range tests {
		v := Trie.Value(str)
		if nil == v {
			t.Error("Key:", str, " not found!")
			return
		}
		if i != v.(int) {
			t.Error("Value not equal")
			return
		}
	}

	if Trie.Value("123") != nil {
		t.Error("Contains other item")
		return
	}
	t.Log("Test Value OK")
}
