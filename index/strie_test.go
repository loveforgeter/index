package index_test

import (
	"testing"

	. "."
)

var STrieIndex = NewSTrie()
var sttests []string = []string{
	"a",
	"abc",
	"b",
	"cdefg",
}

func TestSTrieInsert(t *testing.T) {
	for i, str := range sttests {
		STrieIndex.Insert(str, i)
	}
	t.Log("Test Insert OK")
}

func TestSTrieHasKey(t *testing.T) {
	for _, str := range sttests {
		if !STrieIndex.HasKey(str) {
			t.Error("Key:", str, " not found!")
			return
		}
	}
	t.Log("Test Has OK")
}

func TestSTrieHasPrefix(t *testing.T) {
	for _, str := range sttests {
		if !STrieIndex.HasPrefix(str) {
			t.Error("Prefix", str, " not found!")
			return
		}
	}

	if !STrieIndex.HasPrefix("a") ||
		!STrieIndex.HasPrefix("b") ||
		!STrieIndex.HasPrefix("c") ||
		!STrieIndex.HasPrefix("ab") ||
		!STrieIndex.HasPrefix("cd") ||
		!STrieIndex.HasPrefix("cde") ||
		!STrieIndex.HasPrefix("cdef") {
		t.Error("Item missing")
		return
	}
	t.Log("Test HasPrefix OK")
}

func TestSTrieValueForKey(t *testing.T) {
	for _, str := range sttests {
		v := STrieIndex.ValueForKey(str)
		if nil == v {
			t.Error("Key:", str, " not found!")
			return
		}
	}

	t.Log("Test Value OK")
}

func TestSTrieValueForPrefix(t *testing.T) {
	for _, str := range sttests {
		v := STrieIndex.ValueForPrefix(str)
		if nil == v {
			t.Error("Prefix:", str, " not found!")
			return
		}
	}

	var result map[string][]interface{}
	result = STrieIndex.ValueForPrefix("a")
	if len(result) != 2 {
		t.Error("Wrong number results:", len(result))
		return
	}

	result = STrieIndex.ValueForPrefix("cd")
	if len(result) != 1 {
		t.Error("Wrong number results:", len(result))
		return
	}

	result = STrieIndex.ValueForPrefix("cde")
	if len(result) != 1 {
		t.Error("Wrong number results:", len(result))
		return
	}
}
