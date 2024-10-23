package bitmanipulation

// https://leetcode.cn/problems/reverse-bits/description/

func reverseBits(n uint32) (rev uint32) {
	// 循环32次，对应32位无符号整数
	for i := 0; i < 32 && n > 0; i++ {
		// 获取 n 的最右位，并左移到对应的反转位置，然后与 rev 进行按位或操作
		rev |= (n & 1) << (31 - i)
		// n 右移一位，准备处理下一位
		n >>= 1
	}
	// 返回反转后的结果
	return
}
