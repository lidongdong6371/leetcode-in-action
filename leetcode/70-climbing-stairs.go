package leetcode

// https://leetcode-cn.com/problems/climbing-stairs/

// 斐波那契数列
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

// 超时
// func climbStairs(n int) int {
// 	// 状态转移方程 n层楼梯 需要 (n-2)层 或者 (n-1)层楼梯到达
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	return climbStairs(n-2) + climbStairs(n-1)
// }
