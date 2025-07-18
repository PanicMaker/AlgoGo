package nowcoder

// https://www.nowcoder.com/practice/8b3b95850edb4115918ecebdf1b4d222

func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here

	if pRoot == nil {
		return true
	}

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		return max(depth(node.Left), depth(node.Right)) + 1
	}

	left := depth(pRoot.Left)
	right := depth(pRoot.Right)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	if abs(left-right) > 1 {
		return false
	}

	return true
}
