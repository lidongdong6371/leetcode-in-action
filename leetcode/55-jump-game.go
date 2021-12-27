package leetcode

// https://leetcode-cn.com/problems/jump-game/
func canJump(nums []int) bool {
	return jump(len(nums)-1, nums)
}

func jump(n int, nums []int) bool {
	if n == 0 {
		return true
	}
	// 判断周边是否可达

	for i := 1; i <= n; i++ {
		if nums[n-i] >= i {
			return jump(n-1, nums)
		}
	}
	return false
}
