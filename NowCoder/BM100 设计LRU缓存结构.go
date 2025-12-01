package nowcoder

import "container/list"

// https://www.nowcoder.com/practice/5dfded165916435d9defb053c63f1e84

type pair struct {
	key   int
	value int
}

type Solution struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

func NewLRU(capacity int) Solution {
	return Solution{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *Solution) get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(pair).value
	}
	return -1
}

func (this *Solution) set(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value = pair{key, value}
	} else {
		if this.list.Len() >= this.capacity {
			last := this.list.Back()
			if last != nil {
				evict := last.Value.(pair)
				delete(this.cache, evict.key)
				this.list.Remove(last)
			}
		}
		newElem := this.list.PushFront(pair{key, value})
		this.cache[key] = newElem
	}
}
