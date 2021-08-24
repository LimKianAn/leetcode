package main

import "log"

type NumArray struct {
	segmentTree, nums []int
}

func Constructor(nums []int) NumArray {
	// why 4? https://www.quora.com/Why-does-4-*-N-space-have-to-be-allocated-for-a-segment-tree-where-N-is-the-size-of-the-original-array/answer/Brian-Bi?srid=5ygX
	a := NumArray{segmentTree: make([]int, 4*len(nums)), nums: nums}
	a.init(0, &rng{0, len(nums) - 1})
	return a
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.get(&rng{left, right}, &rng{left: 0, right: len(this.nums) - 1}, 0)
}

func (this *NumArray) Update(index int, val int) {
	this.set(index, val, 0, &rng{0, len(this.nums) - 1})
}

// O(log n)
func (this *NumArray) get(rngToGet, fromRng *rng, node int) int {
	if isEqual(rngToGet, fromRng) {
		return this.segmentTree[node]
	}
	mid := fromRng.mid()
	lChild, rChild := 2*node+1, 2*node+2 // lChild, rChild := left, right child index
	if rngToGet.left > mid {
		return this.get(rngToGet, &rng{left: mid + 1, right: fromRng.right}, rChild)
	}
	if rngToGet.right <= mid {
		return this.get(rngToGet, &rng{left: fromRng.left, right: mid}, lChild)
	}
	return this.get(&rng{left: rngToGet.left, right: mid}, &rng{left: fromRng.left, right: mid}, lChild) +
		this.get(&rng{left: mid + 1, right: rngToGet.right}, &rng{left: mid + 1, right: fromRng.right}, rChild)
}

// O(n log n)
func (this *NumArray) init(node int, r *rng) { // node := node index
	if r.left == r.right {
		this.segmentTree[node] = this.nums[r.left]
		return
	}
	mid := r.mid()
	lChild, rChild := 2*node+1, 2*node+2 // lChild, rChild := left, right child index
	this.init(lChild, &rng{r.left, mid})
	this.init(rChild, &rng{mid + 1, r.right})
	this.segmentTree[node] = this.segmentTree[lChild] + this.segmentTree[rChild]
}

// O(log n)
func (this *NumArray) set(index, val, node int, r *rng) { // node := node index
	if r.left == r.right {
		this.segmentTree[node] = val
		return
	}
	mid := r.mid()
	lChild, rChild := 2*node+1, 2*node+2 // lChild, rChild := left, right child index
	if index > mid {
		this.set(index, val, rChild, &rng{mid + 1, r.right})
	} else {
		this.set(index, val, lChild, &rng{r.left, mid})
	}
	this.segmentTree[node] = this.segmentTree[lChild] + this.segmentTree[rChild]
}

// rng := range
type rng struct {
	left, right int
}

func (r *rng) mid() int {
	return r.left + (r.right-r.left)>>1
}

func isEqual(a, b *rng) bool {
	if a.left == b.left && a.right == b.right {
		return true
	}
	return false
}

func main() {
	a := Constructor([]int{0, 9, 5, 7, 3})
	log.Print(a.SumRange(4, 4))
	log.Print(a.SumRange(2, 4))
	log.Print(a.SumRange(3, 3))
	a.Update(4, 5)
	a.Update(1, 7)
	a.Update(0, 8)
	log.Print(a.SumRange(1, 2))
	a.Update(1, 9)
	log.Print(a.SumRange(4, 4))
	a.Update(3, 4)
}
