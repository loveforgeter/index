package trie

// IndexTree 索引树
type IndexTree map[string]IndexNode

// IndexTree 索引树节点
type IndexNode map[string]interface{}

const (
	PREFIX_NUM = 3
)

func NewIndexTree() IndexTree {
	return make(map[string]IndexNode)
}

func NewIndexNode() IndexNode {
	return IndexNode(make(map[string]interface{}))
}

func (self IndexTree) Insert(key string, value interface{}) {
	keyLen := len(key)
	storeKey := key
	if keyLen > PREFIX_NUM {
		storeKey = key[:PREFIX_NUM]
	}
	var child IndexNode
	var ok bool
	if child, ok = self[storeKey]; !ok {
		child = NewIndexNode()
		self[storeKey] = child
	}
	child[key] = value
}

func (self IndexTree) Has(key string) bool {
	keyLen := len(key)
	if keyLen < PREFIX_NUM {
		_, ok := self[key]
		return ok
	} else {
		storeKey := key[:PREFIX_NUM]
		if child, ok := self[storeKey]; ok {
			_, ok := child[key]
			return ok
		}
	}
	return false
}

func (self IndexTree) Value(key string) interface{} {
	keyLen := len(key)
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
