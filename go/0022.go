// 2022.01.16

package main

func generateParenthesis(n int) []string {
	ss := []string{}
	var dfs func(s string, l, r int)
	dfs = func(s string, l, r int) { // l := numbers of parenthesis, r := right
		if l < r {
			return // things like ()), )  are not valid
		}

		if l == n && r == n {
			ss = append(ss, s)
		}

		if l < n {
			dfs(s+"(", l+1, r)
		}
		if r < n {
			dfs(s+")", l, r+1)
		}
	}
	dfs("", 0, 0)
	return ss
}
