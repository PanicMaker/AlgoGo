package nowcoder

// https://www.nowcoder.com/practice/e19927a8fd5d477794dac67096862042

func solveBM97(n int, m int, a []int) []int {
	step := m % n

	return append(a[n-step:], a[:n-step]...)
}

// 反转数组 arr 的从 start 到 end 范围内的元素
func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start] // 交换首尾元素
		start++
		end--
	}
}

// 循环右移 m 位
func rotateRight(arr []int, m int) {
	n := len(arr)
	if n == 0 || m <= 0 {
		return // 空数组或无偏移
	}
	m = m % n // 处理 m > n 的情况，例如 n=5, m=7 → m=2
	// 第一步：反转前 n - m 个元素
	reverse(arr, 0, n-m-1)
	// 第二步：反转后 m 个元素
	reverse(arr, n-m, n-1)
	// 第三步：反转整个数组
	reverse(arr, 0, n-1)
}
