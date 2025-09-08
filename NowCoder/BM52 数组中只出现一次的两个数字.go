package nowcoder

// https://www.nowcoder.com/practice/389fc1c3d3be4479a154f63f495abff8

func FindNumsAppearOnce(nums []int) []int {
	// 边界条件，数组长度小于2直接返回
	if len(nums) < 2 {
		return nil
	}

	xor := 0
	// 第一步：所有数字异或，最终结果是两个只出现一次的数字的异或值
	for _, num := range nums {
		xor ^= num
	}

	diff := 1
	// 第二步：找到xor结果中最低位的1（即两个数字在这一位上不同）
	for (xor & diff) == 0 {
		diff <<= 1
	}

	num1, num2 := 0, 0
	// 第三步：根据diff这一位是否为1，将数组分为两组，分别异或
	for _, num := range nums {
		if (num & diff) == 0 {
			num1 ^= num // 该位为0的一组
		} else {
			num2 ^= num // 该位为1的一组
		}
	}

	// 返回两个只出现一次的数字
	return []int{num1, num2}
}
