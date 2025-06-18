package nowcoder

import (
	"strconv"
	"strings"
)

// 序列化二叉树
func Serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}

	var result []string
	var preTraverse func(node *TreeNode)
	preTraverse = func(node *TreeNode) {
		if node == nil {
			result = append(result, "null")
			return
		}
		result = append(result, strconv.Itoa(node.Val))
		preTraverse(node.Left)
		preTraverse(node.Right)
	}

	preTraverse(root)
	return strings.Join(result, ",")
}

// 反序列化二叉树
func Deserialize(data string) *TreeNode {
	if data == "null" || len(data) == 0 {
		return nil
	}

	nodes := strings.Split(data, ",")
	var build func() *TreeNode
	index := 0

	build = func() *TreeNode {
		if nodes[index] == "null" {
			index++
			return nil
		}

		val, _ := strconv.Atoi(nodes[index])
		index++
		node := &TreeNode{Val: val}
		node.Left = build()
		node.Right = build()
		return node
	}

	return build()
}
