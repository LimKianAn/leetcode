package main

type MinStack struct {
	elem, min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(val int) {
	this.elem = append(this.elem, val)

	if len(this.min) == 0 {
		this.min = append(this.min, val)
		return
	}
	min := this.GetMin()
	if val < min {
		this.min = append(this.min, val)
		return
	}
	this.min = append(this.min, min) // same minimum
}

func (this *MinStack) Pop() {
	len := len(this.elem)
	this.elem = this.elem[:len-1]
	this.min = this.min[:len-1]
}

func (this *MinStack) Top() int {
	return this.elem[len(this.elem)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
