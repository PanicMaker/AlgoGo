package Tree

// https://www.nowcoder.com/practice/8daa4dff9e36409abba2adbe413d6fae

func isCompleteTree(root *TreeNode) bool {
	// write code here
	if root == nil {
		return true
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	isComplete := false

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == nil {
			isComplete = true
			continue
		}
		if isComplete {
			return false
		}

		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
	}

	return true
}
