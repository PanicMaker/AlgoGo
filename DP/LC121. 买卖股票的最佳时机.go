package DP

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/

func maxProfit(prices []int) int {
	ans := 0
	min_price := prices[0]

	for _, price := range prices {
		min_price = min(min_price, price)
		ans = max(ans, price-min_price)
	}

	return ans
}
