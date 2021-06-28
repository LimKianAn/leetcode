package main

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{sum: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sum[right]
	}
	return this.sum[right] - this.sum[left-1]
}
