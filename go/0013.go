package main

var toInt = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	last := len(s) - 1
	sum := toInt[s[last]]
	for i := last; i > 0; i-- {
		left := toInt[s[i-1]]
		if toInt[s[i]] > left {
			sum -= left
		} else {
			sum += left
		}
	}
	return sum
}
