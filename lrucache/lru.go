//go:build !solution

package lrucache

type LRUCache struct {
	head      *node // cannot convert to diff type. need map in internal funcs
	tail      *node
	keyToNode map[int]*node
	cap       int
}

type node struct {
	key, data   int
	left, right *node
}

func (lc *LRUCache) addLast(key, x int) {

	newNode := &node{key: key, data: x, left: nil, right: nil}
	switch {
	case lc.head == nil:
		// first node
		lc.head = newNode
		lc.tail = newNode
		lc.keyToNode[key] = lc.head
	case lc.tail == lc.head:
		// second node. head != nil
		lc.tail = newNode
		lc.tail.left = lc.head
		lc.head.right = lc.tail
		lc.keyToNode[key] = lc.tail
	default:
		lc.tail.right = newNode
		newNode.left = lc.tail
		lc.tail = newNode
		lc.keyToNode[key] = lc.tail
	}

}

func (lc *LRUCache) popFront() bool {
	switch {
	case lc.head == nil:
		return false
	case lc.head == lc.tail:
		key := lc.head.key
		delete(lc.keyToNode, key)
		lc.head = nil
		lc.tail = nil
	default:
		key := lc.head.key
		delete(lc.keyToNode, key)
		lc.head = lc.head.right
	}
	return true
}

func (lc *LRUCache) nodeToFront(nd *node) *int {
	// replace and return it's value
	if nd == nil {
		return nil
	}
	switch {
	case nd == lc.head:
		if nd != lc.tail {
			lc.head = lc.head.right
			lc.tail.right = nd
			nd.left = lc.tail
			lc.tail = nd
		}
	case nd == lc.tail:
		// tail != head
	default:
		nd.left.right = nd.right
		nd.right.left = nd.left
		lc.tail.right = nd
		nd.right = nil
		nd.left = lc.tail
		lc.tail = nd
	}
	return &nd.data
}

func (lc *LRUCache) Get(key int) (int, bool) {
	node, isPresent := lc.keyToNode[key]
	if !isPresent || node == nil {
		return 0, false
	}
	lc.nodeToFront(node)
	return node.data, true
}

func (lc *LRUCache) Set(key, value int) {
	if _, found := lc.Get(key); found {
		lc.keyToNode[key].data = value
		//lc.nodeToFront(lc.keyToNode[key])
	} else {
		if len(lc.keyToNode) >= lc.cap {
			if !lc.popFront() {
				return
			}
		}
		lc.addLast(key, value)
	}
}

func (lc *LRUCache) Range(f func(key, value int) bool) {
	for node := lc.head; node != nil; node = node.right {
		if cont := f(node.key, node.data); !cont {
			break
		}
	}
}

func (lc *LRUCache) Clear() {
	clear(lc.keyToNode)
	for lc.head != nil {
		lc.popFront()
	}
}

func New(cap int) Cache {
	cache := LRUCache{}
	cache.keyToNode = make(map[int]*node, cap)
	cache.cap = cap
	return &cache
}

////go:build !solution
//
//package lrucache
//
//type LRUCache struct {
//	head      *node // cannot convert to diff type. need map in internal funcs
//	tail      *node
//	keyToNode map[int]*node
//	cap       int
//}
//
//type node struct {
//	key, data   int
//	left, right *node
//}
//
//func (lc *LRUCache) addLast(key, x int) {
//
//	newNode := &node{key: key, data: x, left: nil, right: nil}
//	switch {
//	case lc.head == nil:
//		// first node
//		lc.head = newNode
//		lc.tail = newNode
//		lc.keyToNode[key] = lc.head
//	case lc.tail == lc.head:
//		// second node. head != nil
//		lc.tail = newNode
//		lc.tail.left = lc.head
//		lc.head.right = lc.tail
//		lc.keyToNode[key] = lc.tail
//	default:
//		lc.tail.right = newNode
//		newNode.left = lc.tail
//		lc.tail = newNode
//		lc.keyToNode[key] = lc.tail
//	}
//
//}
//
//func (lc *LRUCache) popFront() bool {
//	switch {
//	case lc.head == nil:
//		return false
//	case lc.head == lc.tail:
//		key := lc.head.key
//		delete(lc.keyToNode, key)
//		lc.head = nil
//		lc.tail = nil
//	default:
//		key := lc.head.key
//		delete(lc.keyToNode, key)
//		lc.head = lc.head.right
//	}
//	return true
//}
//
//func (lc *LRUCache) nodeToFront(nd *node) *int {
//	// replace and return it's value
//	if nd == nil {
//		return nil
//	}
//	switch {
//	case nd == lc.head:
//		if nd != lc.tail {
//			lc.head = lc.head.right
//			lc.tail.right = nd
//			nd.left = lc.tail
//			lc.tail = nd
//		}
//	case nd == lc.tail:
//		// tail != head
//	default:
//		nd.left.right = nd.right
//		nd.right.left = nd.left
//		lc.tail.right = nd
//		nd.right = nil
//		nd.left = lc.tail
//		lc.tail = nd
//	}
//	return &nd.data
//}
//
//func (lc *LRUCache) Get(key int) (int, bool) {
//	node, isPresent := lc.keyToNode[key]
//	if !isPresent || node == nil {
//		return 0, false
//	}
//	lc.nodeToFront(node)
//	return node.data, true
//}
//
//func (lc *LRUCache) Set(key, value int) {
//	if _, found := lc.Get(key); found {
//		lc.keyToNode[key].data = value
//		//lc.nodeToFront(lc.keyToNode[key])
//	} else {
//		if len(lc.keyToNode) >= lc.cap {
//			if !lc.popFront() {
//				return
//			}
//		}
//		lc.addLast(key, value)
//	}
//}
//
//func (lc *LRUCache) Range(f func(key, value int) bool) {
//	for node := lc.head; node != nil; node = node.right {
//		if cont := f(node.key, node.data); !cont {
//			break
//		}
//	}
//}
//
//func (lc *LRUCache) Clear() {
//	clear(lc.keyToNode)
//	for lc.head != nil {
//		lc.popFront()
//	}
//}
//
//func New(cap int) Cache {
//	cache := LRUCache{}
//	cache.keyToNode = make(map[int]*node, cap)
//	cache.cap = cap
//	return &cache
//}
