package main

// time: n:=len(s), O(n*2^n)
// space: O(n)
func partition(s string) (a [][]string) {
	var dfs func(string, int, []string)
	dfs = func(s string, i int, now []string) {
		if i == len(s) {
			a = append(a, copied(now))
			return
		}

		for j := i; j < len(s); j++ {
			sub := s[i : j+1]
			if isPalindrome(sub) {
				dfs(s, j+1, append(now, sub))
			}
		}
	}

	dfs(s, 0, []string{})
	return
}

func copied(s []string) []string {
	tmp := make([]string, len(s))
	copy(tmp, s)
	return tmp
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s); i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
