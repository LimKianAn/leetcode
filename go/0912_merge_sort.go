package main

func sortArray_m(nums []int) []int { // merge sort from huahua
	tmp := make([]int, len(nums))
	var mergeSort func(l, r int)
	mergeSort = func(l, r int) {
		if l >= r {
			return
		}
		m := l + (r-l)>>1
		mergeSort(l, m)
		mergeSort(m+1, r)

		// merge
		i := l
		j := m + 1
		k := 0
		for i <= m && j <= r {
			if nums[i] <= nums[j] {
				tmp[k] = nums[i]
				i++
				k++
			} else {
				tmp[k] = nums[j]
				j++
				k++
			}
		}

		// either for loop will be needed
		for i <= m {
			tmp[k] = nums[i]
			i++
			k++
		}
		for j <= r {
			tmp[k] = nums[j]
			j++
			k++
		}
		copy(nums[l:r+1], tmp[:r+1-l])
	}
	mergeSort(0, len(nums)-1) // [0,len(nums)-1] closed interval
	return nums
}
