package nowcoder

// https://www.nowcoder.com/practice/50ec6a5b0e4e45348544348278cdcee5

func minNumberDisappeared(nums []int) int {
	dict := make(map[int]int, len(nums))

	for _, v := range nums {
		dict[v] = 1
	}

	res := 1
	for {
		_, ok := dict[res]
		if !ok {
			break
		}
		res++
	}
	return res
}
