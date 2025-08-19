package nowcoder

import "math"

// https://www.nowcoder.com/practice/3911a20b3f8743058214ceaa099eeb45

func minMoney(arr []int, aim int) int {
	// 动态规划数组 dp
	// dp[i] 表示凑成金额 i 所需的最少硬币数量。
	// 大小为 aim + 1，因为索引 0 到 aim 都需要存储。
	dp := make([]int, aim+1)

	// 初始化 dp 数组
	// dp[0] = 0：凑成金额 0 需要 0 个硬币。
	dp[0] = 0
	// 其他所有 dp[i] (i > 0) 都初始化为一个非常大的值，表示目前无法凑成。
	// 这样在后续的 min 比较中，任何有效的数量都会小于这个“无穷大”。
	// 理论上使用 math.MaxInt32 或 math.MaxInt64 都是可以的，
	// 也可以使用一个自定义的足够大的值，例如 aim + 1 (因为最多需要 aim 个 1元硬币，如果存在的话)。
	// 为了鲁棒性，使用 math.MaxInt32 是一个好的选择。
	const INF = math.MaxInt32 // 定义一个常量表示无穷大

	for i := 1; i <= aim; i++ {
		dp[i] = INF // 将所有非0的目标金额初始化为无穷大
	}

	// 动态规划核心逻辑
	// 外层循环：遍历所有可能的金额 i，从 1 到 aim。
	// 内层循环：遍历所有可用的硬币面额 coin。
	// 状态转移方程：dp[i] = min(dp[i], dp[i - coin] + 1)
	// 含义：凑成金额 i 的最少硬币数量，可以通过两种方式得到：
	// 1. 之前已经找到的凑成金额 i 的最优解 (dp[i])。
	// 2. 使用一个面额为 coin 的硬币，再加上凑成金额 (i - coin) 所需的最少硬币数量。
	for i := 1; i <= aim; i++ { // 遍历目标金额
		for _, coin := range arr { // 遍历硬币面额
			// 确保当前金额 i 减去硬币面额 coin 后不会小于 0 (避免数组越界)
			// 并且确保 dp[i - coin] 不是无穷大（即 i - coin 是可达的）
			if i-coin >= 0 && dp[i-coin] != INF {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// 最终结果判断
	// 如果 dp[aim] 仍然是 INF，说明目标金额 aim 无法通过给定的硬币凑成。
	if dp[aim] == INF {
		return -1 // 返回 -1 表示无法凑成
	}
	return dp[aim] // 返回凑成目标金额的最少硬币数量
}
