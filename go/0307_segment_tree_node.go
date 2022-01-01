package main

type NumArray struct {
	root *SegmentTreeNode
}

func ConstructorSegmentTreeNode(nums []int) NumArray {
	var newTree func(startIndex, endIndex int) *SegmentTreeNode
	newTree = func(startIndex, endIndex int) *SegmentTreeNode {
		if startIndex == endIndex {
			return &SegmentTreeNode{
				StartIndex: startIndex,
				EndIndex:   endIndex,
				Sum:        nums[startIndex],
			}
		}

		m := startIndex + (endIndex-startIndex)>>1
		l := newTree(startIndex, m)
		r := newTree(m+1, endIndex)
		return &SegmentTreeNode{
			StartIndex: startIndex,
			EndIndex:   endIndex,
			Sum:        l.Sum + r.Sum,
			Left:       l,
			Right:      r,
		}
	}

	if len(nums) == 0 {
		return NumArray{}
	}
	return NumArray{newTree(0, len(nums)-1)}
}

func (this *NumArray) Update(index int, val int) {
	var update func(root *SegmentTreeNode, i, val int)
	update = func(root *SegmentTreeNode, i, val int) {
		if root.StartIndex == i && root.EndIndex == i {
			root.Sum = val
			return // done!
		}

		m := root.StartIndex + (root.EndIndex-root.StartIndex)>>1
		if i <= m {
			update(root.Left, i, val)
		} else {
			update(root.Right, i, val)
		}
		root.Sum = root.Left.Sum + root.Right.Sum
	}
	update(this.root, index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	var sumRange func(root *SegmentTreeNode, left, right int) int
	sumRange = func(root *SegmentTreeNode, left, right int) int {
		if left == root.StartIndex && right == root.EndIndex {
			return root.Sum
		}

		m := root.StartIndex + (root.EndIndex-root.StartIndex)>>1
		if right <= m {
			return sumRange(root.Left, left, right)
		} else if m < left {
			return sumRange(root.Right, left, right)
		} else {
			return sumRange(root.Left, left, m) + sumRange(root.Right, m+1, right)
		}
	}
	return sumRange(this.root, left, right)
}

type SegmentTreeNode struct {
	StartIndex, EndIndex, Sum int
	Left, Right               *SegmentTreeNode
}
