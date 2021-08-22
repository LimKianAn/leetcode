package main

type Trie struct {
	next   map[rune]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{next: map[rune]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	now := this
	for _, r := range word {
		if _, ok := now.next[r]; !ok {
			now.next[r] = &Trie{next: map[rune]*Trie{}}
		}
		now = now.next[r]
	}
	now.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	now := this
	for _, r := range word {
		if _, ok := now.next[r]; !ok {
			return false
		}
		now = now.next[r]
	}
	return now.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	now := this
	for _, r := range prefix {
		if _, ok := now.next[r]; !ok {
			return false
		}
		now = now.next[r]
	}
	return true
}
