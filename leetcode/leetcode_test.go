package leetcode

import (
	"fmt"
	"math"
	"testing"
)

var esp float64 = 0.0000001

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
