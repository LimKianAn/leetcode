// 2022.02.09

package main

func maxScoreIndices(nums []int) []int {
	n := len(nums)
	ones := 0
	for _, num := range nums {
		ones += num
	}

	ans := []int{}
	curOnes := 0
	maxScore := 0
	for i, num := range nums {
		score := (i - curOnes) + (ones - curOnes) // left zeros + right ones
		if score > maxScore {
			maxScore = score
			ans = []int{}
		}
		if score == maxScore {
			ans = append(ans, i)
		}
		curOnes += num
	}

	zeros := n - ones
	if zeros > maxScore {
		maxScore = zeros
		ans = []int{}
	}
	if zeros == maxScore {
		ans = append(ans, n)
	}
	return ans
}
