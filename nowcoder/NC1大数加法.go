package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算两个数之和
 * @param s string字符串 表示第一个整数
 * @param t string字符串 表示第二个整数
 * @return string字符串
 */
func solve(s string, t string) string {
	i, j := len(s)-1, len(t)-1
	ans := make([]byte, 0)
	carry := 0

	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		x := s[i] - '0'
		y := t[j] - '0'
		tmp := x + y + byte(carry)
		ans = append(ans, (tmp%10)+'0')
		carry = int(tmp) / 10
	}

	for ; i >= 0; i-- {
		x := s[i] - '0'
		tmp := x + byte(carry)
		ans = append(ans, (tmp%10)+'0')
		carry = int(tmp) / 10
	}

	for ; j >= 0; j-- {
		y := t[j] - '0'
		tmp := y + byte(carry)
		ans = append(ans, (tmp%10)+'0')
		carry = int(tmp) / 10
	}

	if carry > 0 {
		ans = append(ans, byte(carry)+'0')
	}

	// Reverse the result to get the final answer
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return string(ans)
}
