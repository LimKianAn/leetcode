package main

import "log"

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	// buy, sell
	b, s := make([]int, n), make([]int, n)
	b[0], b[1] = -prices[0], max(-prices[0], -prices[1]) // b[0] isn't initialized while evaluating b[1]
	s[1] = max(s[0], b[0]+prices[1])
	for i := 2; i < n; i++ {
		b[i] = max(b[i-1], s[i-2]-prices[i]) // s[i-2] because of cooldown
		s[i] = max(s[i-1], b[i-1]+prices[i])
	}
	log.Println(b, s)
	return s[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	log.Println(maxProfit([]int{1, 2, 4}))
}
