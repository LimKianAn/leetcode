package main

func generateParenthesis(n int) []string {
	ss := []string{}

	var dfs func(string, int, int)
	dfs = func(s string, l, r int) { // l := numbers of left parenthesis, r := right
		if l > r {
			return
		}
		if l == 0 && r == 0 {
			ss = append(ss, s)
		}
		if l > 0 {
			dfs(s+"(", l-1, r)
		}
		if r > 0 {
			dfs(s+")", l, r-1)
		}
	}

	dfs("", n, n)

	return ss
}
