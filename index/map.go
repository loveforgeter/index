package index

import "strings"

// IndexMap 索引树
type IndexMap map[string]IndexNode

// IndexMap 索引树节点
type IndexNode map[string][]interface{}

const (
	PREFIX_NUM = 3
)

// NewIndexMap 新建map索引
func NewIndexMap() IIndex {
	return IndexMap(make(map[string]IndexNode))
}

// NewIndexNode 新建索引节点
func NewIndexNode() IndexNode {
	return IndexNode(make(map[string][]interface{}))
}

// Insert 插入节点
func (self IndexMap) Insert(key string, value interface{}) {
	if len(key) == 0 {
		return
	}

	// 按前缀插入
	for i := 1; i <= PREFIX_NUM; i++ {
		self.insert(key, 0, i, value)
	}

	/*
		// 分组插入
		start := 1
		end := start + PREFIX_NUM
		for end <= len(key) {
			self.insert(key, start, end, value)
			start++
			end++
		}
	*/
}

// insert 以key[:length]为存储键插入节点
func (self IndexMap) insert(key string, start, end int, value interface{}) {
	// 范围检测
	if start < 0 || end < 0 ||
		start > len(key) || end > len(key) {
		return
	}

	// 插入值
	storeKey := key[start:end]
	var child IndexNode
	var ok bool
	if child, ok = self[storeKey]; !ok {
		child = NewIndexNode()
		self[storeKey] = child
	}

	if _, ok = child[key]; !ok {
		child[key] = make([]interface{}, 0)
	}
	child[key] = append(child[key], value)
}

// Has 查看是否有名称为key的节点
func (self IndexMap) HasKey(key string) bool {
	keyLen := len(key)
	if 0 == keyLen {
		return false
	}

	if keyLen <= PREFIX_NUM { // 长度小于等于PREFIX_NUM
		if child, ok := self[key]; ok {
			_, ok := child[key]
			return ok
		} else {
			return false
		}
	} else {
		storeKey := key[:PREFIX_NUM]
		if child, ok := self[storeKey]; ok {
			_, ok := child[key]
			return ok
		}
	}
	return false
}

// HasPrefix 查看是否有前缀为prefix的节点
func (self IndexMap) HasPrefix(prefix string) bool {
	keyLen := len(prefix)
	if 0 == keyLen {
		return false
	}

	if keyLen <= PREFIX_NUM {
		_, ok := self[prefix]
		return ok
	} else {
		storeKey := prefix[:PREFIX_NUM]
		if child, ok := self[storeKey]; ok {
			for key, _ := range child {
				if strings.HasPrefix(key, prefix) {
					return true
				}
			}
		}
	}
	return false
}

// HasSubstr 查看是否有名称包含substr的值
func (self IndexMap) HasSubstr(substr string) bool {
	NotImplement("IndexMap.HasSubstr")
	return false
}

// ValueForKey 获取节点名称为key的值
func (self IndexMap) ValueForKey(key string) []interface{} {
	keyLen := len(key)
	if 0 == keyLen {
		return nil
	}

	if keyLen < PREFIX_NUM {
		if child, ok := self[key]; ok {
			return child[key]
		}
	} else {
		storeKey := key[:PREFIX_NUM]
		if child, ok := self[storeKey]; ok {
			if value, ok := child[key]; ok {
				return value
			}
		}
	}
	return nil
}

// ValueForPrefix 获取节点名称前缀为prefix的值
func (self IndexMap) ValueForPrefix(prefix string) map[string][]interface{} {
	keyLen := len(prefix)
	if 0 == keyLen {
		return nil
	}

	if keyLen < PREFIX_NUM {
		child, _ := self[prefix]
		return child
	} else {
		ret := make(map[string][]interface{})
		storeKey := prefix[:PREFIX_NUM]
		if child, ok := self[storeKey]; ok {
			for key, value := range child {
				if strings.HasPrefix(key, prefix) {
					ret[key] = value
				}
			}
		}
		if len(ret) != 0 {
			return ret
		}
	}
	return nil
}

func (self IndexMap) ValueForSubstr(substr string) map[string][]interface{} {
	NotImplement("IndexMap.ValueForSubstr")
	return nil
}

// Destroy 销毁
func (self IndexMap) Destroy() {
	NotImplement("IndexMap.Destroy")
}
