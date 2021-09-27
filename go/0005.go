package main

// Time: O(n^2)
func longestPalindrome(s string) string {
	g := "" // global
	for i := range s {
		g = longest(g, longestLocal(s, i, i), longestLocal(s, i, i+1))
	}
	return g
}

func longestLocal(s string, low, high int) string {
	l := "" // local
	for low >= 0 && high < len(s) && s[low] == s[high] {
		l = s[low : high+1]
		low--
		high++
	}
	return l
}

func longest(a, b, c string) string {
	tmp := longer(a, b)
	return longer(tmp, c)
}

func longer(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}
