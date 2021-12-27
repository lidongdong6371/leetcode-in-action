package leetcode

// https://leetcode-cn.com/problems/delete-and-earn/

func deleteAndEarn(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// 转换
	trans := map[int]int{}
	// 边界
	n := 0
	for _, v := range nums {
		if n < v {
			n = v
		}
		trans[v] += v
	}
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, trans[1]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], trans[i]+dp[i-2])
	}
	return dp[n]
}
