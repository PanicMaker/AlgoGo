package nowcoder

// https://www.nowcoder.com/practice/9be0172896bd43948f8a32fb954e1be1

var arr []int

// Insert 将新元素插入到有序数组中的正确位置
func Insert(num int) {
	if arr == nil {
		arr = make([]int, 0)
	}
	// 使用条件安全的插入方式查找位置
	i := 0
	for i < len(arr) && num > arr[i] {
		i++
	}
	arr = append(arr[:i], append([]int{num}, arr[i:]...)...)
}

// GetMedian 返回当前数组的中位数
func GetMedian() float64 {
	n := len(arr)
	if n == 0 {
		return 0.0 // 或者根据业务场景报错
	}
	mid := n / 2
	if n%2 == 0 {
		// 注意修正为 mid - 1 和 mid 两个元素
		return float64(arr[mid-1]+arr[mid]) / 2.0
	}
	return float64(arr[mid])
}
