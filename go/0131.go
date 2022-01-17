// 2022.01.17

package main

func partition(s string) [][]string {
	isPalindrome := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	ans := [][]string{}
	var dfs func(int, []string)
	dfs = func(start int, ss []string) {
		if start == len(s) { // exceeds the upper bound
			cp := make([]string, len(ss))
			copy(cp, ss)
			ans = append(ans, cp)
			return
		}

		for i := start; i < len(s); i++ {
			sub := s[start : i+1]
			if isPalindrome(sub) {
				dfs(i+1, append(ss, sub)) // further
			}
		}
	}
	dfs(0, []string{})
	return ans
}
