// 2022.01.21

package main

func wordCount(startWords []string, targetWords []string) int { // from huahua
	bitFields := func(s string) int {
		sum := 0
		for _, rn := range s {
			sum |= 1 << (rn - 'a')
		}
		return sum
	}

	existing := map[int]bool{}
	for _, w := range startWords {
		existing[bitFields(w)] = true
	}

	ans := 0
	for _, w := range targetWords {
		bf := bitFields(w) // bf := bit fields
		for _, rn := range w {
			if existing[bf^1<<(rn-'a')] {
				ans++
				break // any variant would do
			}
		}
	}
	return ans
}
