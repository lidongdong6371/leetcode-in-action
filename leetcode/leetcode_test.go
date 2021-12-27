package leetcode

import (
	"fmt"
	"math"
	"testing"
)

var esp float64 = 0.0000001

func Test152(t *testing.T) {
	// [-2,3,-4] 24
	// [-2,0,-1] 0
	// [2,-5,-2,-4,3] 24
	nums := []int{2, -5, -2, -4, 3}
	respect := 24
	result := maxProduct(nums)
	if result == respect {
		t.Logf("Success: result is: %d", result)
	} else {
		t.Errorf("Failed:respect %d， but result is: %d", respect, result)
	}
}

func Test55(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	respect := false
	result := canJump(nums)
	if result == respect {
		t.Logf("Success: result is: %t", result)
	} else {
		t.Errorf(" Failed:respect %t， but result is: %t", respect, result)
	}
}

func Test740(t *testing.T) {
	nums := []int{2, 2, 3, 3, 3, 4}
	respect := 9
	result := deleteAndEarn(nums)
	if result == respect {
		t.Logf("Success: result is: %d", result)
	} else {
		t.Errorf(" Failed:respect %d， but result is: %d", respect, result)
	}
}

func Test213(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	respect := 4
	result := rob2(nums)
	if result == respect {
		t.Logf("Success: result is: %d", result)
	} else {
		t.Errorf(" Failed:respect %d， but result is: %d", respect, result)
	}
}

func Test198(t *testing.T) {
	nums := []int{2, 1, 1, 2}
	respect := 4
	result := rob(nums)
	if result == respect {
		t.Logf("Success: result is: %d", result)
	} else {
		t.Errorf(" Failed:respect %d， but result is: %d", respect, result)
	}
}

func Test746(t *testing.T) {
	cost := []int{10, 15, 20}
	respect := 15
	result := minCostClimbingStairs(cost)
	if result == respect {
		t.Logf("Success: result is: %d", result)
	} else {
		t.Errorf(" Failed:respect %d， but result is: %d", respect, result)
	}
}

func Test1137(t *testing.T) {
	n := 4
	respect := 4
	result := tribonacci(n)
	if result == respect {
		t.Logf("Success: result is: %d", result)
	} else {
		t.Errorf(" Failed:respect %d， but result is: %d", respect, result)
	}
}

func Test509(t *testing.T) {
	n := 30
	respect := 2
	result := fib(n)
	if result == respect {
		t.Logf("Success: result is: %d", result)
	} else {
		t.Errorf(" Failed:respect %d， but result is: %d", respect, result)
	}
}

func Test121(t *testing.T) {
	// nums1 := []int{7, 1, 5, 3, 6, 4}
	nums2 := []int{7, 6, 4, 3, 1}
	respect := 0
	result := maxProfit(nums2)
	if result == respect {
		t.Logf("Success: result is: %d", result)
	} else {
		t.Errorf(" Failed:respect %d， but result is: %d", respect, result)
	}
}

func Test350(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	respect := []int{4, 9}
	result := intersect(nums1, nums2)
	if equalInt(result, respect) {
		t.Logf("Success: result is: %v", result)
	} else {
		t.Errorf(" Failed:respect %v， but result is: %v", respect, result)
	}
}

func Test01(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	respect := []int{1, 2}
	result := twoSum(nums, target)
	if equalInt(result, respect) {
		t.Logf("Success: result is: %v", result)
	} else {
		t.Errorf(" Failed:respect %v， but result is: %v", respect, result)
	}
}

func Test70(t *testing.T) {
	n := 10
	respect := 89
	result := climbStairs(n)
	if result == respect {
		t.Logf("Success: result is: %d", result)
	} else {
		t.Errorf(" Failed:respect %d， but result is: %d", respect, result)
	}
}

func Test963(t *testing.T) {
	points := [][]int{{0, 1}, {2, 1}, {1, 1}, {1, 0}, {2, 0}}
	// minAreaFreeRect(points)
	fmt.Println("Begin Test")
	respect := 1.00000
	result := minAreaFreeRect(points)
	if math.Dim(result, respect) < esp {
		t.Logf("Success: result is: %f", result)
	} else {
		t.Errorf(" Failed:respect 1， but result is: %f", result)
	}
}
func Test217(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	respect := true
	result := containsDuplicate(nums)
	if result == respect {
		t.Logf("Success: result is: %t", result)
	} else {
		t.Errorf(" Failed:respect %t， but result is: %t", respect, result)
	}
}

func Test53(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// nums := []int{1}
	// nums := []int{5, 4, -1, 7, 8}
	respect := 6
	result := maxSubArray(nums)
	if result == respect {
		t.Logf("Success: result is: %d", result)
	} else {
		t.Errorf(" Failed:respect %d， but result is: %d", respect, result)
	}
}
