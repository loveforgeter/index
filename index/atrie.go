package index

import "gw.com.cn/dzhyun/app.key.git/index/sparse"

// ATrieNode 字典树节点
type ATrieNode struct {
	key      string        // 节点所表示的字符串
	value    []interface{} // 节点的值
	children *sparse.Array // 子节点
}

// NewATrieNode 新建字典树节点
func NewATrieNode() *ATrieNode {
	return &ATrieNode{
		children: sparse.NewArray(),
		value:    make([]interface{}, 0),
	}
}

// NewATrie 新建字典树
func NewATrie() IIndex {
	return NewATrieNode()
}

// Insert 插入数据
func (self *ATrieNode) Insert(key string, value interface{}) {
	if len(key) == 0 {
		return
	}

	node := self
	for _, ch := range key {
		if child := node.children.Get(ch); nil == child {
			child := NewATrieNode()
			child.key = node.key + string([]rune{ch})
			node.children.Set(ch, child)
			node = child
		} else {
			node = child.(*ATrieNode)
		}
	}
	node.value = append(node.value, value)
}

// HasKey 查看是否存在名称为key的节点
func (self *ATrieNode) HasKey(key string) bool {
	if len(key) == 0 {
		return false
	}

	node := self
	for _, ch := range key {
		if child := node.children.Get(ch); nil == child {
			return false
		} else {
			node = child.(*ATrieNode)
		}
	}
	if len(node.value) == 0 {
		return false
	}
	return true
}

// HasPrefix 查看是否存在名称前缀为prefix的节点
func (self *ATrieNode) HasPrefix(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	node := self
	for _, ch := range prefix {
		if child := node.children.Get(ch); nil == child {
			return false
		} else {
			node = child.(*ATrieNode)
		}
	}
	return true
}

// HasSubstr 查看是否存在名称包含substr的节点
func (self *ATrieNode) HasSubstr(substr string) bool {
	NotImplement("ATrieNode.HasSubstr")
	return false
}

// ValueForKey 获取节点名称为key的值
func (self *ATrieNode) ValueForKey(key string) []interface{} {
	if len(key) == 0 {
		return nil
	}

	node := self
	for _, ch := range key {
		if child := node.children.Get(ch); nil == child {
			return nil
		} else {
			node = child.(*ATrieNode)
		}
	}
	if len(node.value) == 0 {
		return nil
	}
	return node.value
}

// ValueForPrefix 获取节点名称前缀为prefix的值
func (self *ATrieNode) ValueForPrefix(prefix string) map[string][]interface{} {
	if len(prefix) == 0 {
		return nil
	}

	node := self
	ret := make(map[string][]interface{})
	for _, ch := range prefix {
		if child := node.children.Get(ch); nil == child {
			return nil
		} else {
			node = child.(*ATrieNode)
		}
	}

	// 遍历节点并加入结果
	node.walk(func(n *ATrieNode) {
		if len(n.value) != 0 {
			ret[n.key] = n.value
		}
	})

	if len(ret) != 0 {
		return ret
	}
	return nil
}

// ValueForSubstr 获取节点名称包含substr的值
func (self *ATrieNode) ValueForSubstr(substr string) map[string][]interface{} {
	NotImplement("ATrieNode.ValueForSubstr")
	return nil
}

// Destroy 销毁字典树
func (self *ATrieNode) Destroy() {
	NotImplement("ATrieNode.Destroy")
}

// walk 遍历节点
func (self *ATrieNode) walk(f func(node *ATrieNode)) {
	// fmt.Println(self.key, self.value)
	f(self)
	self.children.ForEach(func(r rune, v interface{}) {
		v.(*ATrieNode).walk(f)
	})
}
