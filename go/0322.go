package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // for the first comparison
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount { // when the above `amount+1` comes through
		return -1
	}

	return dp[amount]
}

// func main() {
// 	log.Println(coinChange([]int{2}, 3))
// }
