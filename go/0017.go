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

func letterCombinations(digits string) (results []string) {
	if digits == "" {
		return
	}

	var f func(sub string)
	f = func(sub string) {
		if len(sub) == len(digits) { // out of range
			results = append(results, sub)
			return
		}

		byte := digits[len(sub)] // sub contains the info of the current index
		s := ss[byte-'2']
		for _, r := range s { // r := rune
			f(sub + string(r))
		}
		return
	}
	f("")

	return
}
