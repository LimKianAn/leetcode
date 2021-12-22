package main

type LRUCache struct {
	cache      map[int]*node
	cap        int
	head, tail *node
}

// (-) tail head (+)
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache: make(map[int]*node),
		cap:   capacity,
	}

	return lru
}

func (this *LRUCache) Get(key int) int {
	if nd, ok := this.cache[key]; ok {
		this.rm(nd)
		this.put(nd)
		return nd.v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if nd, ok := this.cache[key]; ok {
		nd.v = value // updates the value
		this.rm(nd)
		this.put(nd)
		return
	}

	nd := &node{k: key, v: value}
	this.cache[key] = nd
	this.put(nd) // updates the doubly linked list

	if len(this.cache) > this.cap { // too many entries
		delete(this.cache, this.tail.k) // map
		this.rm(this.tail) // doubly linked list
	}
}

func (this *LRUCache) put(nd *node) { // growing from left to right
	nd.left = this.head // doubly linked list
	if this.head != nil {
		this.head.right = nd // doubly linked list
	}
	this.head = nd

	if this.tail == nil { // at the very beginning
		this.tail = nd
	}
}

func (this *LRUCache) rm(nd *node) {
	if nd == this.head {
		this.head = nd.left
		return
	}

	if nd == this.tail {
		this.tail = nd.right
		nd.right.left = nil
		return
	}

	nd.right.left = nd.left
	nd.left.right = nd.right
}

type node struct {
	left, right *node
	k, v        int
}
