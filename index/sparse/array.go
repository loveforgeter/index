package sparse

type Array struct {
	keys   []rune
	values []interface{}
}

func NewArray() *Array {
	return &Array{
		keys:   make([]rune, 0),
		values: make([]interface{}, 0),
	}
}

func (self *Array) Len() int {
	return len(self.keys)
}

func (self *Array) Set(r rune, v interface{}) {
	if 0 == self.Len() {
		self.keys = append(self.keys, r)
		self.values = append(self.values, v)
		return
	}
	low := 0
	high := self.Len() - 1
	for low <= high {
		mid := (low + high) / 2
		if self.keys[mid] == r {
			self.keys[mid] = r
			self.values[mid] = v
			return
		} else if self.keys[mid] > r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	self.keys = append(self.keys[:low], append([]rune{r}, self.keys[low:]...)...)
	self.values = append(self.values[:low], append([]interface{}{v}, self.values[low:]...)...)
}

func (self *Array) Get(r rune) interface{} {
	i := self.search(r)
	if -1 == i {
		return nil
	}
	return self.values[i]
}

func (self *Array) Delete(r rune) {
	i := self.search(r)
	if -1 == i {
		return
	}
	self.keys = append(self.keys[:i], self.keys[i+1:]...)
	self.values = append(self.values[:i], self.values[i+1:]...)
}

func (self *Array) ForEach(f func(r rune, v interface{})) {
	for i := 0; i < self.Len(); i++ {
		f(self.keys[i], self.values[i])
	}
}

func (self *Array) search(r rune) int {
	low := 0
	high := self.Len() - 1
	for low <= high {
		mid := (low + high) / 2
		if self.keys[mid] == r {
			return mid
		} else if self.keys[mid] > r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
