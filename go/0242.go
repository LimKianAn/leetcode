package main

func isAnagram(s string, t string) bool {
	counts := map[rune]int{}
	for _, rn := range s {
		counts[rn]++
	}

	for _, rn := range t {
		counts[rn]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}
