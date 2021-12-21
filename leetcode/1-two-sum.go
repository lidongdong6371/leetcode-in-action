package leetcode

// https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if p, ok := m[target-v]; ok {
			return []int{p, i}
		}
		m[v] = i
	}
	return nil
}

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }
