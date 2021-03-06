package index_test

import (
	"flag"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	. "."
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var TIndex = NewTrie()

func init() {
	flag.Parse()
	bench := flag.Lookup("test.bench")
	if "" == bench.Value.String() {
		return
	}
	PrintMem("Initial Stat")
	t := time.Now()
	i := 1000000
	for i > 0 {
		TIndex.Insert(randomKey(), i)
		i--
	}
	fmt.Println("Init time: ", time.Since(t))
	PrintMem("Stat After Trie Init")
}

func BenchmarkTrieHasKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TIndex.HasKey(randomKey())
	}
}

func BenchmarkTrieHasPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TIndex.HasPrefix(randomKey())
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

func PrintMem(msg string) {
	separator := "-------------------------------------"
	fmt.Println(separator)
	defer fmt.Println(separator)
	fmt.Println(msg)
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)
	numField := t.NumField()
	for i := 0; i < numField; i++ {
		st := t.Field(i)
		sv := v.Field(i)
		switch st.Type.String() {
		case "uint64":
			fmt.Println(st.Name, "->", Readable(sv.Uint()))
		case "int64":
			fmt.Println(st.Name, "->", sv.Int())
		case "bool":
			fmt.Println(st.Name, "->", sv.Bool())
		case "float64", "float32":
			fmt.Println(st.Name, "->", sv.Float())
		}
	}
}

const (
	BYTE     = 1.0
	KILOBYTE = 1024 * BYTE
	MEGABYTE = 1024 * KILOBYTE
	GIGABYTE = 1024 * MEGABYTE
	TERABYTE = 1024 * GIGABYTE
)

func Readable(bytes uint64) string {
	unit := ""
	value := float32(bytes)
	switch {
	case bytes >= TERABYTE:
		unit = "T"
		value = value / TERABYTE
	case bytes >= GIGABYTE:
		unit = "G"
		value = value / GIGABYTE
	case bytes >= MEGABYTE:
		unit = "M"
		value = value / MEGABYTE
	case bytes >= KILOBYTE:
		unit = "K"
		value = value / KILOBYTE
	case bytes >= BYTE:
		unit = "B"
	case bytes == 0:
		return "0"
	}
	stringValue := fmt.Sprintf("%.1f", value)
	stringValue = strings.TrimSuffix(stringValue, ".0")
	return fmt.Sprintf("%s%s", stringValue, unit)
}
