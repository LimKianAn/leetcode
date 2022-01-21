// 2022.01.21

package main

import "container/list"

func maxSlidingWindow_m(nums []int, k int) []int {
	q := list.New() // queue of indices
	ans := make([]int, 0, len(nums))
	for i, num := range nums {
		for q.Len() > 0 && nums[q.Back().Value.(int)] <= num {
			q.Remove(q.Back())
		}
		q.PushBack(i)

		if start := i - k + 1; start >= 0 {
			ans = append(ans, nums[q.Front().Value.(int)])
			if start == q.Front().Value.(int) {
				q.Remove(q.Front())
			}
		}
	}
	return ans
}
