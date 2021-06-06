package main

func twoSum(nums []int, target int) []int {
	reverse := map[int]int{}
	for i, v := range nums {
		if j, ok := reverse[target-v]; ok {
			return []int{j, i}
		}
		reverse[v] = i
	}
	return []int{}
}
