package main

import "fmt"

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
//示例 1:
//
//
//
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//示例 2：
//
//
//
//输入： heights = [2,4]
//输出： 4
//
//
//提示：
//
//1 <= heights.length <=105
//0 <= heights[i] <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/largest-rectangle-in-histogram
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

type s struct {
	index  int
	height int
}

func largestRectangleArea(heights []int) int {
	sArr := make([]*s, 0)
	l := make([]int, len(heights))
	r := make([]int, len(heights))

	//算出左右比自己小的柱子位置
	for i := 0; i < len(heights); i++ {
		for {
			if len(sArr) == 0 {
				l[i] = -1
				sArr = append(sArr, &s{
					index:  i,
					height: heights[i],
				})
				break
			} else if heights[i] > sArr[len(sArr)-1].height {
				l[i] = sArr[len(sArr)-1].index
				sArr = append(sArr, &s{
					index:  i,
					height: heights[i],
				})
				break
			} else {
				sArr = sArr[0 : len(sArr)-1]
			}
		}
	}

	sArr = nil
	for i := len(heights) - 1; i >= 0; i-- {
		for {
			if len(sArr) == 0 {
				r[i] = len(heights)
				sArr = append(sArr, &s{
					index:  i,
					height: heights[i],
				})
				break
			} else if heights[i] > sArr[len(sArr)-1].height {
				r[i] = sArr[len(sArr)-1].index
				sArr = append(sArr, &s{
					index:  i,
					height: heights[i],
				})
				break
			} else {
				sArr = sArr[0 : len(sArr)-1]
			}
		}
	}

	result := 0
	//遍历计算面积，比较大小
	for k, height := range heights {
		result = max(result, (r[k]-l[k]-1)*height)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
