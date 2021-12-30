package main

func minEatingSpeed(piles []int, h int) int {
	max := func(nums []int) int {
		ans := 0
		for _, num := range nums {
			if ans < num {
				ans = num
			}
		}
		return ans
	}(piles)

	l, r := 1, max+1
	for l < r {
		m := l + (r-l)>>1
		localH := 0
		for _, pile := range piles {
			localH += (pile-1)/m + 1 // remainder == 0 or != 0
		}

		if localH <= h { // lower bound (how about upper bound, i.e. localH < h ?)
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// func main() {
// 	print(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
// }
