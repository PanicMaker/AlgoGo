package nowcoder

// https://www.nowcoder.com/practice/e297fdd8e9f543059b0b5f05f3a7f3b2

func judge(str string) bool {
	// write code here
	if len(str) == 0 {
		return false
	}

	if len(str) == 1 {
		return true
	}

	l, r := 0, len(str)-1

	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}

	return true
}
