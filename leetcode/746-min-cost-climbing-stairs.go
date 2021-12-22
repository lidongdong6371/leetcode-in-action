package leetcode

// https://leetcode-cn.com/problems/min-cost-climbing-stairs/
func minCostClimbingStairs(cost []int) int {
	// cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]

	// return min(minCost(cost, len(cost)-1), minCost(cost, len(cost)-2))
	return minCost2(cost)
}

// 动态规划
func minCost2(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

// 递归超出时间限制
func minCost(cost []int, n int) int {
	if n == 0 {
		return cost[0]
	}
	if n == 1 {
		return cost[1]
	}

	return min(cost[n]+minCost(cost, n-1), cost[n]+minCost(cost, n-2))

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
