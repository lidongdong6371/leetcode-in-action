package leetcode

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit2(prices []int) int {
	ans := 0
	low := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[low] {
			if i+1 < len(prices) && prices[i+1] < prices[i] {
				ans += prices[i] - prices[low]
				low = i + 1
				i++
			} else if i+1 == len(prices) {
				ans += prices[i] - prices[low]
			}
		} else {
			low = i
		}
	}
	return ans
}
