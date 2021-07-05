package main

func lengthOfLongestSubstring(s string) (n int) {
	index := map[rune]int{}
	start := 0
	for end, r := range s {
		if i, ok := index[r]; ok {
			if i >= start {
				start = i + 1 // starts from the rune right after the one which has already occurred
			}
		}
		index[r] = end
		n = max(n, end-start+1) // length = difference+1
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
