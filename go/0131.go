package main

import "log"

func partition(s string) (a [][]string) {
	var dfs func(int, []string)
	dfs = func(start int, strings []string) {
		if start == len(s) { // exceeding the upper bound
			log.Println(copied(strings))
			a = append(a, copied(strings))
			return
		}

		for j := start; j < len(s); j++ {
			sub := s[start : j+1]
			if isPalindrome(sub) {
				dfs(j+1, append(strings, sub)) // further evaluating
			}
		}
	}

	dfs(0, []string{})
	return
}

func copied(s []string) []string {
	tmp := make([]string, len(s))
	copy(tmp, s)
	return tmp
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	partition("aab")
}
