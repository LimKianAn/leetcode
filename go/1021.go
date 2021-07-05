package main

func removeOuterParentheses(s string) (out string) {
	tmp := ""
	delta := 0
	for _, ch := range s {
		tmp += string(ch)
		if ch == '(' {
			delta++
		} else {
			delta--
		}
		if delta == 0 {
			out += tmp[1 : len(tmp)-1]
			tmp = ""
		}
	}
	return
}
