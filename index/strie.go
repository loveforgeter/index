package index

// STrieNode 字典树节点
type STrieNode struct {
	key       rune
	value     []interface{}       // 节点的值
	children  []*STrieNode        // 子节点
	mchildren map[rune]*STrieNode // 子节点
}

// NewSTrieNode 新建字典树节点
func NewSTrieNode() *STrieNode {
	return &STrieNode{
		mchildren: make(map[rune]*STrieNode, 0),
		children:  make([]*STrieNode, 0),
		value:     make([]interface{}, 0),
	}
}

// NewTrie 新建字典树
func NewSTrie() IIndex {
	return NewSTrieNode()
}

// Insert 插入数据
func (self *STrieNode) Insert(key string, value interface{}) {
	if len(key) == 0 {
		return
	}

	node := self
	for _, ch := range key {
		if child := node.get(ch); nil == child {
			child = NewSTrieNode()
			child.key = ch
			node.set(ch, child)
			node = child
		} else {
			node = child
		}
	}
	return
	node.value = append(node.value, value)
}

// HasKey 查看是否存在名称为key的节点
func (self *STrieNode) HasKey(key string) bool {
	if len(key) == 0 {
		return false
	}

	node := self
	for _, ch := range key {
		if node = node.get(ch); nil == node {
			return false
		}
	}
	if len(node.value) == 0 {
		return false
	}
	return true
}

// HasPrefix 查看是否存在名称前缀为prefix的节点
func (self *STrieNode) HasPrefix(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	node := self
	for _, ch := range prefix {
		if node = node.get(ch); nil == node {
			return false
		}
	}
	return true
}

// HasSubstr 查看是否存在名称包含substr的节点
func (self *STrieNode) HasSubstr(substr string) bool {
	NotImplement("STrieNode.HasSubstr")
	return false
}

// ValueForKey 获取节点名称为key的值
func (self *STrieNode) ValueForKey(key string) []interface{} {
	if len(key) == 0 {
		return nil
	}

	node := self
	for _, ch := range key {
		if node = node.get(ch); nil == node {
			return nil
		}
	}
	if len(node.value) == 0 {
		return nil
	}
	return node.value
}

// ValueForPrefix 获取节点名称前缀为prefix的值
func (self *STrieNode) ValueForPrefix(prefix string) map[string][]interface{} {
	if len(prefix) == 0 {
		return nil
	}

	node := self
	ret := make(map[string][]interface{})
	for _, ch := range prefix {
		if node = node.get(ch); nil == node {
			return nil
		}
	}

	// 遍历节点并加入结果
	node.walk(func(n *STrieNode) {
		if len(n.value) != 0 {
		}
	})

	if len(ret) != 0 {
		return ret
	}
	return nil
}

// ValueForSubstr 获取节点名称包含substr的值
func (self *STrieNode) ValueForSubstr(substr string) map[string][]interface{} {
	NotImplement("STrieNode.ValueForSubstr")
	return nil
}

// Destroy 销毁字典树
func (self *STrieNode) Destroy() {
	NotImplement("STrieNode.Destroy")
}

// walk 遍历节点
func (self *STrieNode) walk(f func(node *STrieNode)) {
	// fmt.Println(self.key, self.value)
	f(self)
	if nil != self.children {
		for _, n := range self.children {
			n.walk(f)
		}
	} else {
		for _, node := range self.mchildren {
			node.walk(f)
		}
	}
}

func (self *STrieNode) get(ch rune) *STrieNode {
	if nil != self.children {
		for i := 0; i < len(self.children); i++ {
			if ch == self.children[i].key {
				return self.children[i]
			}
		}
	} else {
		return self.mchildren[ch]
	}
	return nil
}

func (self *STrieNode) set(ch rune, n *STrieNode) {
	if nil == n {
		return
	}
	if len(self.mchildren) != 0 {
		self.mchildren[ch] = n
	} else {
		length := len(self.children)
		if 10 == length {
			for i := 0; i < length; i++ {
				self.mchildren[ch] = self.children[i]
			}
			self.children = nil
			self.mchildren[ch] = n
		} else {
			self.children = append(self.children, n)
		}
	}
}
