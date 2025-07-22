package nowcoder

// https://www.nowcoder.com/practice/c9480213597e45f4807880c763ddd5f0

func solveBM41(preOrder []int, inOrder []int) []int {
	if len(preOrder) < 1 || len(inOrder) < 1 {
		return []int{}
	}

	var build func(pre []int, in []int) *TreeNode
	build = func(pre, in []int) *TreeNode {
		if len(pre) == 0 {
			return nil
		}
		rootVal := pre[0]
		root := &TreeNode{Val: rootVal}

		var inRootIndex int = -1
		for i, v := range in {
			if v == rootVal {
				inRootIndex = i
				break
			}
		}

		if inRootIndex == -1 {
			return nil // Should not happen in a valid case
		}

		leftSubtreeSize := inRootIndex
		root.Left = build(pre[1:1+leftSubtreeSize], in[:leftSubtreeSize])
		root.Right = build(pre[1+leftSubtreeSize:], in[leftSubtreeSize+1:])

		return root
	}

	ans := make([]int, 0)

	root := build(preOrder, inOrder)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelSize := len(queue)
		ans = append(ans, queue[levelSize-1].Val)

		newQueue := make([]*TreeNode, 0, levelSize*2)
		for i := 0; i < levelSize; i++ {
			node := queue[i]

			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}

		queue = newQueue
	}

	return ans
}
