package index_test

import (
	"testing"

	. "."
)

var MapIndex = NewIndexMap()
var mtests []string = []string{
	"a",
	"abc",
	"b",
	"cdefg",
}

func TestMapInsert(t *testing.T) {
	for i, str := range mtests {
		MapIndex.Insert(str, i)
	}
	t.Log("Test Insert OK")
}

func TestMapHasKey(t *testing.T) {
	for _, str := range mtests {
		if !MapIndex.HasKey(str) {
			t.Error("Key:", str, " not found!")
			return
		}
	}
	t.Log("Test Has OK")
}

func TestMapHasPrefix(t *testing.T) {
	for _, str := range mtests {
		if !MapIndex.HasPrefix(str) {
			t.Error("Prefix", str, " not found!")
			return
		}
	}
	if !MapIndex.HasPrefix("a") ||
		!MapIndex.HasPrefix("b") ||
		!MapIndex.HasPrefix("c") ||
		!MapIndex.HasPrefix("ab") ||
		!MapIndex.HasPrefix("cd") ||
		!MapIndex.HasPrefix("cde") ||
		!MapIndex.HasPrefix("cdef") {
		t.Error("Item missing")
		return
	}

	t.Log("Test HasPrefix OK")
}

func TestMapValueForKey(t *testing.T) {
	for _, str := range mtests {
		v := MapIndex.ValueForKey(str)
		if nil == v {
			t.Error("Key:", str, " not found!")
			return
		}
	}

	t.Log("Test Value OK")
}

func TestMapValueForPrefix(t *testing.T) {
	for _, str := range mtests {
		v := MapIndex.ValueForPrefix(str)
		if nil == v {
			t.Error("Prefix:", str, " not found!")
			return
		}
	}

	var result map[string][]interface{}
	result = MapIndex.ValueForPrefix("a")
	if len(result) != 2 {
		t.Error("Wrong number results:", len(result))
		return
	}

	result = MapIndex.ValueForPrefix("cd")
	if len(result) != 1 {
		t.Error("Wrong number results:", len(result))
		return
	}

	result = MapIndex.ValueForPrefix("cde")
	if len(result) != 1 {
		t.Error("Wrong number results:", len(result))
		return
	}

}
