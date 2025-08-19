package nowcoder

// https://www.nowcoder.com/practice/91b69814117f4e8097390d107d2efbe0

func Print(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{pRoot}
	direction := 1 // 1 表示正序，0 表示逆序
	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			idx := 0
			if direction%2 == 1 {
				idx = i
			} else {
				idx = levelSize - 1 - i
			}
			level[idx] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)
		direction++
	}
	return res
}
