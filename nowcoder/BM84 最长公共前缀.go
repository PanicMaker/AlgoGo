package nowcoder

// https://www.nowcoder.com/practice/28eb3175488f4434a4a6207f6f484f47

func longestCommonPrefix(strs []string) string {
	n := len(strs)

	if n == 0 {
		return ""
	}

	if n == 1 {
		return strs[0]
	}

	for i := 1; i < len(strs[0]); i++ {
		char := strs[0][i]

		for j := 0; j < n; j++ {
			if i >= len(strs[j]) || char != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
