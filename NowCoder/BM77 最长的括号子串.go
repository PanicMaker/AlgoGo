package nowcoder

// https://www.nowcoder.com/practice/45fd68024a4c4e97a8d6c45fc61dc6ad

func longestValidParentheses(s string) int {
	stack := make([]int, 0)

	res := 0
	length := 0
	start := -1

	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				start = i
			} else {
				index := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					length = i - index
				} else {
					length = i - start
				}
			}
		}

		res = max(res, length)
	}

	return res
}
