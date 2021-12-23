package main

func combine(n int, k int) (result [][]int) {
	var f func(head int, memo []int)
	f = func(head int, memo []int) {
		if len(memo) == k {
			result = append(result, copiedInts(memo))
			return
		}

		for i := head; i <= n; i++ {
			f(i+1, append(memo, i))
		}
	}


	return
}

// func main() {
// 	combine(5, 2)
// }
