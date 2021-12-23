package main

var open = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

func isValid0020(s string) bool {
	opens := []rune{}
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			opens = append(opens, v)
		} else {
			if len(opens) == 0 {
				return false
			}
			last := len(opens) - 1
			if opens[last] != open[v] {
				return false
			}
			opens = opens[:last] // drops the last open parentheses
		}
	}
	return len(opens) == 0
}
