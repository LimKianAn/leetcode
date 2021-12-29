package main

import "log"

type Heap []*Element

func NewHeap(ee []*Element) *Heap { // log(n)
	h := Heap(ee)
	for i := h.Len()/2 - 1; 0 <= i; i-- { // starts from the second last row
		h.heapifyDown(i)
	}
	return &h
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Peek() *Element {
	return (*h)[0]
}

func (h *Heap) Pop() *Element { // removes the tip and re-heapifies
	tip := h.Peek()

	h.swap(h.Len()-1, 0)
	(*h) = (*h)[:h.Len()-1]
	h.heapifyDown(0) // re-heapifies

	return tip
}

func (h *Heap) Push(e *Element) {
	(*h) = append((*h), e)
	h.heapifyUp(h.Len() - 1)
}

func (h *Heap) heapifyDown(i int) { // theta(log n)
	leftChild := 2*i + 1
	rightChild := 2*i + 2
	if leftChild >= h.Len() { // doesn't exists
		return
	}

	child := leftChild
	if rightChild < h.Len() && h.predicate(rightChild, leftChild) {
		child = rightChild
	}

	if h.predicate(i, child) {
		return
	}
	h.swap(i, child)
	h.heapifyDown(child)
}

func (h *Heap) heapifyUp(i int) { // theta(log n)
	if i == 0 {
		return
	}

	parent := (i - 1) / 2
	if h.predicate(parent, i) {
		return
	}
	h.swap(parent, i)
	h.heapifyUp(parent)
}

// > for max heap
// < for min heap
func (h *Heap) predicate(i, j int) bool {
	return (*h)[i].count >= (*h)[j].count
}

func (h *Heap) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

type Element struct {
	num, count int
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	ee := make([]*Element, 0, len(m))
	for k, v := range m {
		ee = append(ee, &Element{k, v})
	}

	h := NewHeap(ee)
	ii := []int{}
	for i := 0; i < k; i++ {
		ii = append(ii, h.Pop().num)
	}
	return ii
}

func main() {
	log.Print(topKFrequent([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}, 3))
}
