package index

import "fmt"

// IIndex 索引树接口
type IIndex interface {
	Insert(key string, value interface{})
	HasKey(key string) bool
	HasPrefix(prefix string) bool
	HasSubstr(substr string) bool
	ValueForKey(key string) []interface{}
	ValueForPrefix(prefix string) map[string][]interface{}
	ValueForSubstr(substr string) map[string][]interface{}
	Destroy()
}

func NotImplement(name string) {
	fmt.Println("\033[0;31m", name, " not implement!", "\033[0m")
}

func NotFinished(name string) {
	fmt.Println("\033[0;31m", name, " not finished!", "\033[0m")
}
