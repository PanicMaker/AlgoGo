package Tree

type BSTIterator struct {
	arr []int
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{
		arr: make([]int, 0),
	}

	it.inorder(root)

	return it
}

func (this *BSTIterator) inorder(node *TreeNode) {
	if node == nil {
		return
	}

	this.inorder(node.Left)
	this.arr = append(this.arr, node.Val)
	this.inorder(node.Right)
}

func (this *BSTIterator) Next() int {
	val := this.arr[0]
	this.arr = this.arr[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.arr) > 0
}
