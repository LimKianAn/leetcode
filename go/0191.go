package main

// func hammingWeight(num uint32) int {
// 	return bits.OnesCount(uint(num))
// }

func hammingWeight(num uint32) (count int) {
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return
}
