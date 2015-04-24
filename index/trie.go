package index

const (
	MAX_ARRAY_SIZE = 10
)

type hash map[rune]*TrieNode
type array []*TrieNode

// TrieNode 字典树节点
type TrieNode struct {
	key    rune               // 节点代表的值
	values []interface{}      // 节点的值
	array  []*TrieNode        // 子节点
	hash   map[rune]*TrieNode // 子节点
}

// NewTrieNode 新建字典树节点
func NewTrieNode() *TrieNode {
	return &TrieNode{
		values: make([]interface{}, 0),
		array:  make([]*TrieNode, 0),
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
	for _, r := range key {
		if child := node.get(r); nil == child {
			child = NewTrieNode()
			child.key = r
			node.set(r, child)
			node = child
		} else {
			node = child
		}
	}
	node.values = append(node.values, value)
}

// HasKey 查看是否存在名称为key的节点
func (self *TrieNode) HasKey(key string) bool {
	if len(key) == 0 {
		return false
	}

	node := self
	for _, r := range key {
		if node = node.get(r); nil == node {
			return false
		}
	}

	// 没有值的节点为空节点
	if len(node.values) == 0 {
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
	for _, r := range prefix {
		if node = node.get(r); nil == node {
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
	for _, r := range key {
		if node = node.get(r); nil == node {
			return nil
		}
	}

	if len(node.values) == 0 {
		return nil
	}

	return node.values
}

// ValueForPrefix 获取节点名称前缀为prefix的值
func (self *TrieNode) ValueForPrefix(prefix string) map[string][]interface{} {
	if len(prefix) == 0 {
		return nil
	}

	node := self
	result := make(map[string][]interface{})
	for _, r := range prefix {
		if node = node.get(r); nil == node {
			return nil
		}
	}

	// 遍历节点并加入结果
	node.walk(func(key string, n *TrieNode) {
		if len(n.values) != 0 {
			result[key] = n.values
		}
	}, prefix)

	if len(result) == 0 {
		return nil
	}

	return result
}

// ValueForSubstr 获取节点名称包含substr的值
func (self *TrieNode) ValueForSubstr(substr string) map[string][]interface{} {
	NotImplement("TrieNode.ValueForSubstr")
	return nil
}

// Destroy 销毁字典树
func (self *TrieNode) Destroy() {
}

// walk 遍历节点，key代表从根到父节点所代表的字符串
func (self *TrieNode) walk(f func(key string, node *TrieNode), key string) {
	// 当前节点代表的字符串
	current := key + string([]rune{self.key})
	f(current, self)
	if nil != self.array {
		for _, child := range self.array {
			child.walk(f, current)
		}
	} else {
		for _, child := range self.hash {
			child.walk(f, current)
		}
	}
}

// get 获取子节点
func (self *TrieNode) get(r rune) *TrieNode {
	if nil != self.array {
		for _, node := range self.array {
			if node.key == r {
				return node
			}
		}
	} else {
		return self.hash[r]
	}
	return nil
}

// set 设置子节点
func (self *TrieNode) set(r rune, child *TrieNode) {
	// TODO:重复处理
	if nil == child {
		return
	}
	if nil != self.hash {
		self.hash[r] = child
	} else {
		if MAX_ARRAY_SIZE == len(self.array) {
			self.hash = make(map[rune]*TrieNode)
			for _, node := range self.array {
				self.hash[node.key] = node
			}
			self.hash[r] = child
			self.array = nil
		} else {
			self.array = append(self.array, child)
		}
	}
}
