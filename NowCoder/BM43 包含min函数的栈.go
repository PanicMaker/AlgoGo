package nowcoder

// https://www.nowcoder.com/practice/4c776177d2c04c2494f2555c9fcc1e49

type minStack struct {
	arr    []int
	minNum int
}

var stack *minStack

func NewMinStack() {
	if stack == nil {
		stack = &minStack{
			arr: make([]int, 0),
		}
	}
}

func Push(node int) {
	NewMinStack()

	stack.arr = append(stack.arr, node)
	if node < stack.minNum {
		stack.minNum = node
	}
}
func Pop() {
	NewMinStack()
	num := stack.arr[len(stack.arr)-1]
	stack.arr = stack.arr[:len(stack.arr)-1]

	if num == stack.minNum {
		stack.minNum = stack.arr[0]
		for _, v := range stack.arr {
			if v < stack.minNum {
				stack.minNum = v
			}
		}
	}
}
func Top() int {
	NewMinStack()

	return stack.arr[len(stack.arr)-1]
}
func Min() int {
	NewMinStack()

	return stack.minNum
}
