package leetcode

// https://leetcode-cn.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		if m[num] > 0 {
			return true
		}
		m[num]++
	}
	return false
}
