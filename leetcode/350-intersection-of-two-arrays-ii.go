package leetcode

import "sort"

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
// 排序双指针
func intersect(nums1 []int, nums2 []int) []int {
	p1, p2 := 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	l1 := len(nums1)
	l2 := len(nums2)
	res := []int{}
	for {
		if p1 == l1 || p2 == l2 {
			return res
		}
		if nums1[p1] == nums2[p2] {
			res = append(res, nums1[p1])
			p1++
			p2++
			continue
		}
		if nums1[p1] < nums2[p2] {
			p1++
		} else {
			p2++
		}
	}
}

// hash table 方式
// func intersect(nums1 []int, nums2 []int) []int {
// 	m := map[int]int{}
// 	for _, v := range nums1 {
// 		m[v]++
// 	}
// 	result := []int{}
// 	for _, v := range nums2 {
// 		if val, ok := m[v]; ok && val > 0 {
// 			result = append(result, v)
// 			m[v]--
// 		}
// 	}
// 	return result
// }
