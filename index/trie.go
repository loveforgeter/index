package index

// TrieNode 字典树节点
type TrieNode struct {
	key      string             // 节点所表示的字符串
	value    []interface{}      // 节点的值
	children map[rune]*TrieNode // 子节点
}

// NewTrieNode 新建字典树节点
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		value:    make([]interface{}, 0),
	}
}

// NewTrie 新建字典树
func NewTrie() IIndex {
	return NewTrieNode()
}

// Insert 插入数据
func (self *TrieNode) Insert(key string, value interface{}) {
	if len(key) == 0 {
		return
	}

	node := self
	for _, ch := range key {
		if child, ok := node.children[ch]; !ok {
			child = NewTrieNode()
			child.key = node.key + string([]rune{ch})
			node.children[ch] = child
			node = child
		} else {
			node = child
		}
	}
	node.value = append(node.value, value)
}

// HasKey 查看是否存在名称为key的节点
func (self *TrieNode) HasKey(key string) bool {
	if len(key) == 0 {
		return false
	}

	node := self
	var ok bool
	for _, ch := range key {
		if node, ok = node.children[ch]; !ok || nil == node {
			return false
		}
	}
	if len(node.value) == 0 {
		return false
	}
	return true
}

// HasPrefix 查看是否存在名称前缀为prefix的节点
func (self *TrieNode) HasPrefix(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	node := self
	var ok bool
	for _, ch := range prefix {
		if node, ok = node.children[ch]; !ok {
			return false
		}
	}
	return true
}

// HasSubstr 查看是否存在名称包含substr的节点
func (self *TrieNode) HasSubstr(substr string) bool {
	NotImplement("TrieNode.HasSubstr")
	return false
}

// ValueForKey 获取节点名称为key的值
func (self *TrieNode) ValueForKey(key string) []interface{} {
	if len(key) == 0 {
		return nil
	}

	node := self
	var ok bool
	for _, ch := range key {
		if node, ok = node.children[ch]; !ok {
			return nil
		}
	}
	if len(node.value) == 0 {
		return nil
	}
	return node.value
}

// ValueForPrefix 获取节点名称前缀为prefix的值
func (self *TrieNode) ValueForPrefix(prefix string) map[string][]interface{} {
	if len(prefix) == 0 {
		return nil
	}

	node := self
	var ok bool
	ret := make(map[string][]interface{})
	for _, ch := range prefix {
		if node, ok = node.children[ch]; !ok {
			return nil
		}
	}

	// 遍历节点并加入结果
	node.walk(func(n *TrieNode) {
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
func (self *TrieNode) ValueForSubstr(substr string) map[string][]interface{} {
	NotImplement("TrieNode.ValueForSubstr")
	return nil
}

// Destroy 销毁字典树
func (self *TrieNode) Destroy() {
	for key, child := range self.children {
		child.Destroy()
		delete(self.children, key)
	}
}

// walk 遍历节点
func (self *TrieNode) walk(f func(node *TrieNode)) {
	// fmt.Println(self.key, self.value)
	f(self)
	for _, node := range self.children {
		node.walk(f)
	}
}
