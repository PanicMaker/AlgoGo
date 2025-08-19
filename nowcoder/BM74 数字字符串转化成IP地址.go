package nowcoder

import (
	"fmt"
	"strconv"
)

// https://www.nowcoder.com/practice/ce73540d47374dbe85b3125f57727e1e

func restoreIpAddresses(s string) []string {
	var result []string
	var backtrack func(start int, parts []string)
	backtrack = func(start int, parts []string) {
		if len(parts) == 4 && start == len(s) {
			result = append(result, fmt.Sprintf("%s.%s.%s.%s",
				parts[0], parts[1], parts[2], parts[3]))
			return
		}

		if len(parts) >= 4 || start >= len(s) {
			return // 剪枝
		}

		for i := start; i < len(s); i++ {
			if i-start+1 > 3 {
				break // 单段最大三位数
			}
			subStr := s[start : i+1]
			// 非法字段判断
			if subStr[0] == '0' && len(subStr) > 1 {
				continue // 前导零不合法
			}
			val, err := strconv.Atoi(subStr)
			if err != nil || val > 255 {
				continue // 数字转换失败或超出范围
			}
			// 递归尝试下一个部分
			parts = append(parts, subStr)
			backtrack(i+1, parts)
			parts = parts[:len(parts)-1] // 回溯
		}
	}

	backtrack(0, []string{})
	return result
}
