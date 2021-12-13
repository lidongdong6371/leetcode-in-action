package leetcode

import (
	"fmt"
	"math"
	"testing"
)

var esp float64 = 0.0000001

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
