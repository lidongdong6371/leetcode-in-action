package leetcode

// https://leetcode-cn.com/problems/fibonacci-number/
func fib(n int) int {
	// return fibRecursive(n)
	return fibArr(n)
}

// 递归
func fibRecursive(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// 数组
func fibArr(n int) int {
	if n == 0 {
		return 0
	}
	a, b, c := 0, 1, 1
	for i := 0; i < n-2; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}
