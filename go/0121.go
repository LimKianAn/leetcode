package main

func maxProfit0121(prices []int) int {
	len := len(prices)
	if len <= 1 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len; i++ {
		profit := prices[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return maxProfit
}
