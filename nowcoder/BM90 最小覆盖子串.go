package nowcoder

// https://www.nowcoder.com/practice/c466d480d20c4c7c9d322d12ca7955ac

func minWindow(S string, T string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	valid := 0 // 已经满足的字符种类数
	// 初始化 need 表
	for i := range T {
		need[T[i]]++
	}
	left, right := 0, 0
	lenS := len(S)
	res := ""
	minLength := lenS + 1 // 初始设为比最大可能还长
	for right < lenS {
		c := S[right]
		right++
		// 如果是待找的字符，就添加进 window
		if _, exists := need[c]; exists {
			window[c]++
			if window[c] == need[c] {
				valid += 1
			}
		}
		// 开始向内收缩窗口
		for valid == len(need) {
			// 如果当前窗口比上次更优，则更新结果区间
			if right-left < minLength {
				minLength = right - left
				res = S[left:right]
			}
			d := S[left]
			left++
			if _, exists := need[d]; exists {
				if window[d] == need[d] {
					valid -= 1
				}
				window[d]--
			}
		}
	}
	return res
}
