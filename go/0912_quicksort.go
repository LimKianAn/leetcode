package main

// func sortArray(nums []int) []int { // quicksort from huahua
// 	var quicksort func(l, r int)
// 	quicksort = func(l, r int) {
// 		if l >= r {
// 			return
// 		}
// 		i, j := l, r
// 		pivot := nums[l+rand.Intn(r-l+1)]
// 		for i <= j {
// 			for nums[i] < pivot {
// 				i++
// 			}

// 			for nums[j] > pivot {
// 				j--
// 			}

// 			if i <= j {
// 				nums[i], nums[j] = nums[j], nums[i]
// 				i++
// 				j--
// 			}
// 		}
// 		quicksort(l, j)
// 		quicksort(i, r)
// 	}
// 	quicksort(0, len(nums)-1)
// 	return nums
// }

func sortArray_q(nums []int) []int {
	var mergeSort func(l, r int)
	mergeSort = func(l, r int) {
		if l >= r {
			return
		}

		j := l                   // index of the element smaller than the value at the pivot
		for i := l; i < r; i++ { // r pivot
			if nums[i] < nums[r] {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			}
		}
		nums[j], nums[r] = nums[r], nums[j]
		mergeSort(l, j-1)
		mergeSort(j+1, r)
	}
	mergeSort(0, len(nums)-1)
	return nums
}
