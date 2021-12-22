package main

func combine(n int, k int) (result [][]int) {
	var f func(head int, memo []int)
	f = func(head int, memo []int) {
		if len(memo) == k {
			result = append(result, copied(memo))
			return
		}

		for i := head; i <= n; i++ {
			f(i+1, append(memo, i))
		}
	}
	

	return
}

func copied(ii []int) []int {
	alias := make([]int, len(ii))
	copy(alias, ii)
	return alias
}

func main() {
	combine(5, 2)
}
