package main

type MyStack struct {
	a []int
}

func Constructor() MyStack {
	return MyStack{[]int{}}
}

func (this *MyStack) Push(x int) {
	this.a = append(this.a, x)
}

func (this *MyStack) Pop() (i int) {
	i = this.Top()
	this.a = this.a[:len(this.a)-1]
	return
}

func (this *MyStack) Top() int {
	return this.a[len(this.a)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.a) == 0
}
