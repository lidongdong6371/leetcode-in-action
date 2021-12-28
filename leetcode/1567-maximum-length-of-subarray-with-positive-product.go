package leetcode

// https://leetcode-cn.com/problems/maximum-length-of-subarray-with-positive-product/
func getMaxLen(nums []int) int {
	negative := []int{}
	// 参照,出现0就更新
	refer := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			ans = max(ans, getLen(refer, i, negative))
			refer = i + 1
			negative = []int{}
		}
		if nums[i] < 0 {
			negative = append(negative, i)
		}
	}
	return max(getLen(refer, len(nums), negative), ans)
}

// 默认当前位置 curr 为不合法值 refer 为合法值
func getLen(refer, curr int, negative []int) int {
	n := len(negative)
	if n%2 == 0 {
		return curr - refer
	} else {
		return max(max(negative[n-1]-refer, curr-negative[n-1]-1), max(negative[0]-refer, curr-negative[0]-1))
	}
}
