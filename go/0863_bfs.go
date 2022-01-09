package main

func distanceK_bfs(root *TreeNode, target *TreeNode, k int) []int { // from huahua
	graph := map[*TreeNode][]*TreeNode{}

	var build func(parent, child *TreeNode)
	build = func(parent, child *TreeNode) {
		if parent != nil {
			graph[parent] = append(graph[parent], child)
			graph[child] = append(graph[child], parent)
		}

		if child.Left != nil {
			build(child, child.Left)
		}
		if child.Right != nil {
			build(child, child.Right)
		}
	}
	build(nil, root)

	seen := map[*TreeNode]bool{target: true} // starts from the target!
	q := []*TreeNode{target}

	ans := []int{}
	distance := 0
	for len(q) > 0 && distance <= k {
		for qLen := len(q); qLen > 0; qLen-- {
			node := q[0]
			q = q[1:]

			if distance == k {
				ans = append(ans, node.Val)
			}

			for _, next := range graph[node] {
				if seen[next] {
					continue
				}
				seen[next] = true
				q = append(q, next)
			}
		}
		distance++
	}
	return ans
}
