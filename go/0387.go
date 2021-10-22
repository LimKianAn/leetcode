package main

func firstUniqChar(s string) int {
	abc := [26]int{}

	for _, letter := range s {
		abc[letter-'a']++
	}

	for i, letter := range s {
		if abc[letter-'a'] == 1 {
			return i
		}
	}
	return -1
}
