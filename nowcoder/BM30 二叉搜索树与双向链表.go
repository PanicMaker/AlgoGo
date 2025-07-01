package nowcoder

// https://www.nowcoder.com/practice/947f6eb80d944a84850b0538bf0ec3a5

func Convert(pRootOfTree *TreeNode) *TreeNode {
	// write code here
	if pRootOfTree == nil {
		return nil
	}

	left := Convert(pRootOfTree.Left)
	right := Convert(pRootOfTree.Right)

	if left != nil {
		p := left
		for p.Right != nil {
			p = p.Right
		}

		p.Right = pRootOfTree
		pRootOfTree.Left = p
	}

	if right != nil {
		right.Left = pRootOfTree
		pRootOfTree.Right = right
	}

	if left != nil {
		return left
	}
	return pRootOfTree
}
