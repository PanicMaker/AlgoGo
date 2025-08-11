package nowcoder

// https://www.nowcoder.com/practice/c3a6afee325e472386a1c4eb1ef987f3

func solveBM91(str string) string {
	n := len(str)

	if n <= 1 {
		return str
	}

	arr := []byte(str)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return string(arr)
}
