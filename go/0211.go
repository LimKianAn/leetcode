package main

type WordDictionary struct {
	next  map[rune]*WordDictionary
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor0211() WordDictionary {
	return WordDictionary{next: map[rune]*WordDictionary{}}
}

func (this *WordDictionary) AddWord(word string) {
	now := this
	for _, r := range word {
		if _, ok := now.next[r]; !ok {
			now.next[r] = &WordDictionary{next: map[rune]*WordDictionary{}}
		}
		now = now.next[r]
	}
	now.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	now := this
	for i, r := range word {
		if r == '.' {
			for _, v := range now.next {
				if v.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}
		if _, ok := now.next[r]; !ok {
			return false
		}
		now = now.next[r]
	}
	return now.isEnd
}
