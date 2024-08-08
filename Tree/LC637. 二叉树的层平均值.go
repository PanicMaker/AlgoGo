package Tree

// https://leetcode.cn/problems/average-of-levels-in-binary-tree/description/?envType=study-plan-v2&envId=top-interview-150

func averageOfLevels(root *TreeNode) []float64 {
	ans := make([]float64, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0

		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			sum += cur.Val

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		ans = append(ans, float64(sum)/float64(levelSize))
	}

	return ans
}
