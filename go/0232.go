package main

type MyQueue struct {
	ii []int
}

func Constructor0232() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.ii = append(this.ii, x)
}

func (this *MyQueue) Pop() int {
	i := this.ii[0]
	this.ii = this.ii[1:len(this.ii)]
	return i
}

func (this *MyQueue) Peek() int {
	return this.ii[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.ii) == 0
}
