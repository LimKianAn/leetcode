package main

type NumArray0303 struct {
	sum []int
}

func Constructor0303(nums []int) NumArray0303 {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray0303{sum: nums}
}

func (this *NumArray0303) SumRange(left int, right int) int {
	if left == 0 {
		return this.sum[right]
	}
	return this.sum[right] - this.sum[left-1]
}
