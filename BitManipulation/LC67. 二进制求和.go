package bitmanipulation

// https://leetcode.cn/problems/number-of-1-bits/description/

func addBinary(a string, b string) string {
	flag := 0
	res := ""

	len_a := len(a)
	len_b := len(b)

	for i := 0; i < len_a-len_b; i++ {
		b = "0" + b
	}

	for i := 0; i < len_b-len_a; i++ {
		a = "0" + a
	}

	for i := len(a) - 1; i >= 0; i-- {
		num1 := a[i] - '0'
		num2 := b[i] - '0'

		num := flag + int(num1) + int(num2)

		if num == 0 {
			res = "0" + res
			flag = 0
		} else if num == 1 {
			res = "1" + res
			flag = 0
		} else if num == 2 {
			res = "0" + res
			flag = 1
		} else {
			res = "1" + res
			flag = 1
		}

	}

	if flag == 1 {
		res = "1" + res
	}

	return res
}
