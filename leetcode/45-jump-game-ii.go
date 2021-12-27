package leetcode

// https://leetcode-cn.com/problems/jump-game-ii/
func jump2(nums []int) int {

	end, steps, maxProsition := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		maxProsition = max(maxProsition, end+nums[i])
		if i == end {
			end = maxProsition
			steps++
		}
	}
	return steps
}
