package main

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse0189(nums)
	reverse0189(nums[:k])
	reverse0189(nums[k:])
}

func reverse0189(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
