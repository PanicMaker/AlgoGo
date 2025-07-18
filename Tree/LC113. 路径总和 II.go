package Tree

// https://leetcode.cn/problems/path-sum-ii/description/

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(node *TreeNode, target int)
	backtrack = func(node *TreeNode, target int) {
		if node == nil {
			return
		}

		newTarget := target - node.Val
		if node.Left == nil && node.Right == nil && newTarget == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
		}

		path = append(path, node.Val)
		backtrack(node.Left, newTarget)
		backtrack(node.Right, newTarget)
		path = path[:len(path)-1]
	}

	backtrack(root, targetSum)

	return ans
}
