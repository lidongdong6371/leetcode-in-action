package leetcode

const INT_MIN = ^int(^uint64((0)) >> 1)

// https://leetcode-cn.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {

	return getMax(0, len(nums)-1, nums).mSum
}

func getMax(l, r int, nums []int) Status {
	if r == l {
		return Status{nums[r], nums[r], nums[r], nums[r]}
	}
	m := (1 + r) >> 1

	lSub := getMax(l, m, nums)
	rSub := getMax(m+1, r, nums)

	return pushUp(lSub, rSub)
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

type Status struct {
	lSum, rSum, mSum, iSum int
}

// 官方答案 动态规划
// func maxSubArray(nums []int) int {
// 	res := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i]+nums[i-1] > nums[i] {
// 			nums[i] += nums[i-1]
// 		}
// 		if res < nums[i] {
// 			res = nums[i]
// 		}
// 	}
// 	return res
// }

// 递归 动态规划  状态转移方程用的很差劲 与循环无异
// func maxValue(n, m int, nums []int) int {
// 	if n == m {
// 		return nums[n]
// 	}
// 	x := maxValue(n+1, m, nums)
// 	y := maxValue(n, m-1, nums)
// 	s := sum(nums[n : m+1])
// 	if s > x && s > y {
// 		return s
// 	} else if x > y {
// 		return x
// 	} else {
// 		return y
// 	}
// }

//----------- 朴实无华的循环, 结果也朴实无华,直接超时 -----
// func maxSubArray(nums []int) int {
// 	res := INT_MIN
// 	// 循环
// 	for i, l := 0, len(nums); i < l; i++ {
// 		for j := 0; j+i < l; j++ {
// 			s := sum(nums[j : j+i+1])
// 			if s > res {
// 				res = s
// 			}
// 		}
// 	}
// 	return res
// }

// func sum(nums []int) int {
// 	var res int
// 	for _, num := range nums {
// 		res += num
// 	}
// 	return res
// }

// ------------------------------------------------
