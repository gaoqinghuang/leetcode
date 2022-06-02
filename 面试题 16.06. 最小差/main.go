package main

import (
	"fmt"
	"math"
	"sort"
)

//给定两个整数数组a和b，计算具有最小差绝对值的一对数值（每个数组中取一个值），并返回该对数值的差
//
//
//
//示例：
//
//输入：{1, 3, 15, 11, 2}, {23, 127, 235, 19, 8}
//输出：3，即数值对(11, 8)
//
//
//提示：
//
//1 <= a.length, b.length <= 100000
//-2147483648 <= a[i], b[i] <= 2147483647
//正确结果在区间 [0, 2147483647] 内
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/smallest-difference-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(smallestDifference([]int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8}))
}

func smallestDifference(a []int, b []int) int {
	//排序，双指针
	sort.Ints(a)
	sort.Ints(b)
	result := math.MaxInt

	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			return 0
		}
		result = min(result, uSub(a[i], b[j]))
		if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

func uSub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
