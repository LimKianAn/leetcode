package main

type MapSum struct {
	root *node
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{&node{next: make(map[rune]*node)}}
}

func (this *MapSum) Insert(key string, val int) {
	now := this.root
	for _, r := range key {
		if _, ok := now.next[r]; !ok {
			now.next[r] = &node{next: make(map[rune]*node)}
		}
		now = now.next[r]
	}
	now.val = val
}

func (this *MapSum) Sum(prefix string) int {
	now := this.root
	for _, r := range prefix {
		if _, ok := now.next[r]; !ok {
			return 0
		}
		now = now.next[r]
	}
	return now.sum()
}

type node struct {
	val  int
	next map[rune]*node
}

func (nd *node) sum() int {
	sum := nd.val
	for _, nextNd := range nd.next {
		sum += nextNd.sum()
	}
	return sum
}
