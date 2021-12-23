package main

func wiggleMaxLength(nums []int) int {
	if l := len(nums); l < 2 {
		return l
	}

	l := 1
	diffMemo := nums[1] - nums[0]
	if diffMemo != 0 {
		l = 2
	}

	for i := 2; i < len(nums); i++ {
		if diff := nums[i] - nums[i-1]; isAlternating(diffMemo, diff) {
			l++
			diffMemo = diff
		}
	}
	return l
}

func isAlternating(last, now int) bool {
	return last >= 0 && now < 0 || last <= 0 && now > 0 // "=" is for [3,3,3,2,5]
}

// func main() {
// 	log.Println(wiggleMaxLength([]int{3, 3, 3, 2, 5}))
// }
