package index_test

import (
	"testing"

	. "."
)

var TrieIndex = NewTrie()
var ttests []string = []string{
	"a",
	"abc",
	"b",
	"cdefg",
}

func TestTrieInsert(t *testing.T) {
	for i, str := range ttests {
		TrieIndex.Insert(str, i)
	}
	t.Log("Test Insert OK")
}

func TestTrieHasKey(t *testing.T) {
	for _, str := range ttests {
		if !TrieIndex.HasKey(str) {
			t.Error("Key:", str, " not found!")
			return
		}
	}
	t.Log("Test Has OK")
}

func TestTrieHasPrefix(t *testing.T) {
	for _, str := range ttests {
		if !TrieIndex.HasPrefix(str) {
			t.Error("Prefix", str, " not found!")
			return
		}
	}

	if !TrieIndex.HasPrefix("a") ||
		!TrieIndex.HasPrefix("b") ||
		!TrieIndex.HasPrefix("c") ||
		!TrieIndex.HasPrefix("ab") ||
		!TrieIndex.HasPrefix("cd") ||
		!TrieIndex.HasPrefix("cde") ||
		!TrieIndex.HasPrefix("cdef") {
		t.Error("Item missing")
		return
	}
	t.Log("Test HasPrefix OK")
}

func TestTrieValueForKey(t *testing.T) {
	for _, str := range ttests {
		v := TrieIndex.ValueForKey(str)
		if nil == v {
			t.Error("Key:", str, " not found!")
			return
		}
	}

	t.Log("Test Value OK")
}

func TestTrieValueForPrefix(t *testing.T) {
	for _, str := range ttests {
		v := TrieIndex.ValueForPrefix(str)
		if nil == v {
			t.Error("Prefix:", str, " not found!")
			return
		}
	}

	var result map[string][]interface{}
	result = TrieIndex.ValueForPrefix("a")
	if len(result) != 2 {
		t.Error("Wrong number results:", len(result))
		return
	}

	result = TrieIndex.ValueForPrefix("cd")
	if len(result) != 1 {
		t.Error("Wrong number results:", len(result))
		return
	}

	result = TrieIndex.ValueForPrefix("cde")
	if len(result) != 1 {
		t.Error("Wrong number results:", len(result))
		return
	}
}
