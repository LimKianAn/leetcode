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
		nd.v = value
		this.rm(nd)
		this.put(nd)
		return
	}

	nd := &node{k: key, v: value}
	this.cache[key] = nd
	this.put(nd)

	if len(this.cache) > this.cap {
		delete(this.cache, this.tail.k)
		this.rm(this.tail)
	}
}

func (this *LRUCache) put(nd *node) {
	nd.left = this.head
	if this.head != nil {
		this.head.right = nd
	}
	this.head = nd

	if this.tail == nil {
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
	right, left *node
	k, v        int
}
