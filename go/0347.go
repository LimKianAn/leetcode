package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	counts := newCounts(nums)
	heap.Init(counts)
	ii := []int{}
	for len(ii) < k {
		count := heap.Pop(counts).(count)
		ii = append(ii, count.key)
	}
	return ii
}

type (
	count struct {
		key, count int
	}
	counts []count
)

func newCounts(nums []int) *counts {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	cs := make([]count, len(m))
	for key := range m {
		cs = append(cs, count{key: key, count: m[key]})
	}
	counts := counts(cs)
	return &counts
}

func (c *counts) Len() int {
	return len(*c)
}

func (c *counts) Less(i, j int) bool {
	return (*c)[i].count > (*c)[j].count
}

func (c *counts) Pop() interface{} {
	tmp := (*c)[c.Len()-1]
	*c = (*c)[:c.Len()-1]
	return tmp
}

func (c *counts) Push(x interface{}) {
	*c = append(*c, x.(count))
}

func (c *counts) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}
