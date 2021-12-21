package leetcode

const MAX_INT = int(^uint64((0)) >> 1)

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	minPrice := MAX_INT
	maxProfit := 0
	for _, v := range prices {
		if minPrice > v {
			minPrice = v
		}
		if maxProfit < v-minPrice {
			maxProfit = v - minPrice
		}
	}
	return maxProfit
}
