package nowcoder

// https://www.nowcoder.com/practice/b4525d1d84934cf280439aeecc36f4af

func getLongestPalindrome(A string) int {
	n := len(A)

	result := 0
	for index := 0; index < n; index++ {
		length := 1
		for i, j := index-1, index+1; i >= 0 && j < n; {
			if A[i] == A[j] {
				length += 2
				i, j = i-1, j+1
			} else {
				break
			}
		}
		result = max(result, length)

		length = 0
		for i, j := index, index+1; i >= 0 && j < n; {
			if A[i] == A[j] {
				length += 2
				i, j = i-1, j+1
			} else {
				break
			}
		}
		result = max(result, length)
	}

	return result
}
