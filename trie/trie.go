package trie

// ITrie 字典树
type ITire interface {
	Insert(str string, value interface{})
	Has(str string) bool
	Value(str string) interface{}
	Destroy()
}

// TrieNode 字典树节点
type TrieNode struct {
	value    interface{}
	children map[rune]*TrieNode
}

// NewTrie 新建字典树
func NewTrie() ITire {
	return NewTrieNode()
}

// NewTrieNode 新建字典树节点
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode, 0),
	}
}

// Insert 插入数据
func (self *TrieNode) Insert(str string, value interface{}) {
	node := self
	for _, ch := range str {
		if child, ok := node.children[ch]; !ok {
			child = NewTrieNode()
			node.children[ch] = child
			node = child
		} else {
			node = child
		}
	}
	node.value = value
}

// Has 判断字典树是否存在当前值
func (self *TrieNode) Has(str string) bool {
	node := self
	var ok bool
	for _, ch := range str {
		if node, ok = node.children[ch]; !ok {
			return false
		}
	}
	return true
}

// Value 获取数据
func (self *TrieNode) Value(str string) interface{} {
	node := self
	var ok bool
	for _, ch := range str {
		if node, ok = node.children[ch]; !ok {
			return nil
		}
	}
	return node.value
}

// Destroy 销毁字典树
func (self *TrieNode) Destroy() {
	for key, child := range self.children {
		child.Destroy()
		delete(self.children, key)
	}
}
