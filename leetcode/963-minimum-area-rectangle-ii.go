package leetcode

import (
	"math"
)

type coordinate struct {
	x    float64
	y    float64
	dist float64
}

// https://leetcode-cn.com/problems/minimum-area-rectangle-ii/
func minAreaFreeRect(points [][]int) float64 {
	// map【coordinate][]coordinate key 中点坐标， value端的坐标
	coor := make(map[coordinate][]coordinate)
	// 找出所有线段
	for i, len := 0, len(points); i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			// 计算点之间的长度 计算中点坐标
			x, y := midPoint(points[i], points[j])
			dist := dist(points[i], points[j])
			coor[coordinate{x: x, y: y, dist: dist}] = append(coor[coordinate{x: x, y: y, dist: dist}], coordinate{x: float64(points[i][0]), y: float64(points[i][1]), dist: dist})
		}
	}

	var area float64
	for k, v := range coor {
		// 只有一条中线表示不能构成矩形
		if len(v) < 2 {
			continue
		}
		// 计算边长
		for i, l := 0, len(v); i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				side1 := math.Sqrt(math.Pow(v[i].x-v[j].x, 2) + math.Pow(v[i].y-v[j].y, 2))
				side2 := math.Sqrt(math.Pow(k.dist/2, 2)-math.Pow(side1/2, 2)) * 2
				if side1*side2 < area || area == 0 {
					area = side1 * side2
				}
			}
		}
	}
	return area
}

func midPoint(p1 []int, p2 []int) (float64, float64) {
	return float64(p1[0]+p2[0]) / 2, float64(p1[1]+p2[1]) / 2.0
}

func dist(p1 []int, p2 []int) float64 {
	return math.Sqrt(math.Pow(float64(p1[0]-p2[0]), 2) + math.Pow(float64(p1[1]-p2[1]), 2))
}
