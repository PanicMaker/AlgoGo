package nowcoder

// https://www.nowcoder.com/practice/20ef0972485e41019e39543e8e895b7f

func twoSum(numbers []int, target int) []int {
	maps := make(map[int]int)

	for i, v := range numbers {
		maps[v] = i
	}

	for i, v := range numbers {
		if j, ok := maps[target-v]; ok {
			if i == j {
				continue
			}
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}
