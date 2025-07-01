package nowcoder

// https://www.nowcoder.com/practice/7298353c24cc42e3bd5f0e0bd3d1d759

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// write code here
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
