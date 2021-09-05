package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}
	for len(t) > 0 {
		if s[0] == t[0] {
			s = s[1:]
			if len(s) == 0 {
				return true
			}
		}
		t = t[1:]
	}
	return false
}
