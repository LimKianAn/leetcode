package main

var ss = []string{
	"abc", // 2
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz", // 9
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	results := []string{}
	f(&results, "", digits, 0)
	return results
}

func f(results *[]string, sub string, digits string, index int) {
	if index == len(digits) { // out of range
		*results = append(*results, sub)
		return
	}

	b := digits[index] // b := byte
	s := ss[b-'2']
	for _, r := range s { // r := rune
		f(results, sub+string(r), digits, index+1)
	}
	return
}
