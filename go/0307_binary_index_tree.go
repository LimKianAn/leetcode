package main

type NumArrayBIT struct {
	tree, nums []int
}

// O(n)
// https://www.youtube.com/watch?v=uSFzHCZ4E-8
func ConstructorBinaryIndexTree(nums []int) NumArrayBIT {
	a := NumArrayBIT{tree: make([]int, len(nums)+1), nums: nums}
	copy(a.tree[1:], nums)
	for i := 1; i < len(a.tree); i++ {
		p := i + i&-i // p := parent
		if p < len(a.tree) {
			a.tree[p] += a.tree[i]
		}
	}
	return a
}

func (this *NumArrayBIT) SumRange(left int, right int) int {
	return this.sum(right+1) - this.sum(left)
}

func (this *NumArrayBIT) Update(index int, val int) {
	this.add(index+1, val-this.nums[index])
	this.nums[index] = val
}

// O(log n)
func (this *NumArrayBIT) add(i, v int) { // i, v := index, value
	for ; i < len(this.tree); i += i & -i {
		this.tree[i] += v
	}
}

// O(log n)
func (this *NumArrayBIT) sum(i int) int {
	sum := 0
	for ; i > 0; i -= i & -i {
		sum += this.tree[i]
	}
	return sum
}

// func main() {
// 	a := Constructor0307b([]int{5, 2, 9, -3, 5, 20, 10, -7, 2, 3, -4, 0, -2, 15, 5})
// 	log.Print(a.sum(7), a.sum(8))
// 	a.add(4, 10)
// 	log.Print(a.tree[4], a.tree[8])
// }
