//在二维平面上计算出两个由直线构成的矩形重叠后形成的总面积。
//
//每个矩形由其左下顶点和右上顶点坐标表示，如图所示。
//
//
//
//示例:
//
//输入: -3, 0, 3, 4, 0, -1, 9, 2
//输出: 45
//说明: 假设矩形面积不会超出 int 的范围。

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
}
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	area1 := (C - A) * (D - B)
	area2 := (G - E) * (H - F)
	//没有重叠的情况
	//上下左右
	if F > D || B > H || A > G || E > C {
		return area1 + area2
	}
	//重叠坐标
	left := int(math.Max(float64(A), float64(E)))
	right := int(math.Min(float64(C), float64(G)))
	top := int(math.Min(float64(H), float64(D)))
	down := int(math.Max(float64(B), float64(F)))
	return area1 + area2 - (right-left)*(top-down)
}
