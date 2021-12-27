package leetcode

// https://leetcode-cn.com/problems/house-robber-ii/
func rob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	nums1 := nums[:n-1]
	n1 := len(nums1)
	dp1 := make([]int, n1)
	dp1[0], dp1[1] = nums1[0], max(nums1[0], nums1[1])
	for i := 2; i < n1; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums1[i])
	}
	nums2 := nums[1:]
	n2 := len(nums2)
	dp2 := make([]int, n2)
	dp2[0], dp2[1] = nums2[0], max(nums2[0], nums2[1])
	for i := 2; i < n2; i++ {
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums2[i])
	}
	return max(dp1[n1-1], dp2[n2-1])
}
