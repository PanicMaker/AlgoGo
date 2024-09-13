package BackTrack

// https://leetcode.cn/problems/remove-invalid-parentheses/

// 选或不选
func removeInvalidParenthesesI(s string) []string {
	n := len(s)
	ans := make([]string, 0)
	path := make([]byte, 0)
	visited := make(map[string]bool)

	lCount, rCount := 0, 0

	for _, v := range s {
		if v == '(' {
			lCount++
		} else if v == ')' {
			if lCount > 0 {
				lCount--
			} else {
				rCount++
			}
		}
	}

	var backtrack func(i int, left int, right int, lCount int, rCount int)
	backtrack = func(i int, left int, right int, lCount int, rCount int) {
		// 如果删除的左右括号数量超过所需的数量，提前返回
		if lCount < 0 || rCount < 0 {
			return
		}

		// 如果到达字符串末尾并且删除的括号数量刚好匹配
		if i == n {
			if left-right == 0 && lCount == 0 && rCount == 0 {
				res := string(path)
				if !visited[res] {
					visited[res] = true
					ans = append(ans, res)
				}
			}
			return
		}

		if s[i] == '(' {
			// 删除左括号的选择
			backtrack(i+1, left, right, lCount-1, rCount)

			path = append(path, '(')
			backtrack(i+1, left+1, right, lCount, rCount)
			path = path[:len(path)-1]
		} else if s[i] == ')' {
			// 删除右括号的选择
			backtrack(i+1, left, right, lCount, rCount-1)

			if left > right {
				path = append(path, ')')
				backtrack(i+1, left, right+1, lCount, rCount)
				path = path[:len(path)-1]
			}
		} else {
			path = append(path, s[i])
			backtrack(i+1, left, right, lCount, rCount)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0, 0, lCount, rCount)

	return ans
}

// 枚举答案
func removeInvalidParentheses(s string) []string {
	// 初始化结果集
	ans := make([]string, 0)

	// 计算需要删除的左括号和右括号的数量
	lCount, rCount := 0, 0
	for _, v := range s {
		if v == '(' {
			lCount++ // 遇到左括号，计数加1
		} else if v == ')' {
			if lCount > 0 {
				lCount-- // 若存在多余的左括号，抵消一个
			} else {
				rCount++ // 需要删除的右括号数量加1
			}
		}
	}

	// 辅助函数，判断一个字符串的括号是否有效
	isValid := func(s string) bool {
		cnt := 0 // 记录左括号的数量
		for _, v := range s {
			if v == '(' {
				cnt++ // 遇到左括号，计数加1
			} else if v == ')' {
				cnt--        // 遇到右括号，计数减1
				if cnt < 0 { // 若右括号数量大于左括号，返回无效
					return false
				}
			}
		}
		return cnt == 0 // 若计数最终为0，字符串有效
	}

	// 回溯函数，用于生成所有可能的字符串组合
	var backtrack func(i int, s string, lCount int, rCount int)
	backtrack = func(i int, s string, lCount int, rCount int) {
		// 若需要删除的括号数量小于0，停止回溯
		if lCount < 0 || rCount < 0 {
			return
		}

		// 当左右括号删除数量均为0时，检查字符串的有效性
		if lCount == 0 && rCount == 0 {
			if isValid(s) {
				ans = append(ans, s) // 若有效，加入结果集
			}
			return
		}

		// 遍历字符串中的字符
		for j := i; j < len(s); j++ {
			// 若当前字符与上一个相同，跳过以避免重复
			if j > i && s[j] == s[j-1] {
				continue
			}

			// 尝试删除左括号
			if s[j] == '(' && lCount > 0 {
				backtrack(j, s[:j]+s[j+1:], lCount-1, rCount)
			}
			// 尝试删除右括号
			if s[j] == ')' && rCount > 0 {
				backtrack(j, s[:j]+s[j+1:], lCount, rCount-1)
			}
		}
	}

	// 开始回溯
	backtrack(0, s, lCount, rCount)

	return ans
}
