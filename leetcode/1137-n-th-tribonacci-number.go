package leetcode

//https://leetcode-cn.com/problems/n-th-tribonacci-number/
//泰波那契序列 Tn 定义如下：
// T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
// 超出时间限制
// func tribonacci(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 || n == 2 {
// 		return 1
// 	}
// 	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
// }

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	q, r, p, ans := 0, 1, 1, 2
	for i := 0; i < n-3; i++ {
		q = r
		r = p
		p = ans
		ans = q + r + p
	}
	return ans
}
