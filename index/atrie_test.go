package index_test

import (
	"testing"

	. "."
)

var ATrieIndex = NewATrie()
var attests []string = []string{
	"a",
	"abc",
	"b",
	"cdefg",
}

func TestATrieInsert(t *testing.T) {
	for i, str := range attests {
		ATrieIndex.Insert(str, i)
	}
	t.Log("Test Insert OK")
}

func TestATrieHasKey(t *testing.T) {
	for _, str := range attests {
		if !ATrieIndex.HasKey(str) {
			t.Error("Key:", str, " not found!")
			return
		}
	}
	t.Log("Test Has OK")
}

func TestATrieHasPrefix(t *testing.T) {
	for _, str := range attests {
		if !ATrieIndex.HasPrefix(str) {
			t.Error("Prefix", str, " not found!")
			return
		}
	}

	if !ATrieIndex.HasPrefix("a") ||
		!ATrieIndex.HasPrefix("b") ||
		!ATrieIndex.HasPrefix("c") ||
		!ATrieIndex.HasPrefix("ab") ||
		!ATrieIndex.HasPrefix("cd") ||
		!ATrieIndex.HasPrefix("cde") ||
		!ATrieIndex.HasPrefix("cdef") {
		t.Error("Item missing")
		return
	}
	t.Log("Test HasPrefix OK")
}

func TestATrieValueForKey(t *testing.T) {
	for _, str := range attests {
		v := ATrieIndex.ValueForKey(str)
		if nil == v {
			t.Error("Key:", str, " not found!")
			return
		}
	}

	t.Log("Test Value OK")
}

func TestATrieValueForPrefix(t *testing.T) {
	for _, str := range attests {
		v := ATrieIndex.ValueForPrefix(str)
		if nil == v {
			t.Error("Prefix:", str, " not found!")
			return
		}
	}

	var result map[string][]interface{}
	result = ATrieIndex.ValueForPrefix("a")
	if len(result) != 2 {
		t.Error("Wrong number results:", len(result))
		return
	}

	result = ATrieIndex.ValueForPrefix("cd")
	if len(result) != 1 {
		t.Error("Wrong number results:", len(result))
		return
	}

	result = ATrieIndex.ValueForPrefix("cde")
	if len(result) != 1 {
		t.Error("Wrong number results:", len(result))
		return
	}
}
