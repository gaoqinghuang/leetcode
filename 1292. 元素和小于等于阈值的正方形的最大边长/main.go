package main

import "fmt"

//给你一个大小为 m x n 的矩阵 mat 和一个整数阈值 threshold。
//
//请你返回元素总和小于或等于阈值的正方形区域的最大边长；如果没有这样的正方形区域，则返回 0 。
//
//
//示例 1：
//
//
//
//输入：mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4
//输出：2
//解释：总和小于或等于 4 的正方形的最大边长为 2，如图所示。
//示例 2：
//
//输入：mat = [[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]], threshold = 1
//输出：0
//
//
//提示：
//
//m == mat.length
//n == mat[i].length
//1 <= m, n <= 300
//0 <= mat[i][j] <= 104
//0 <= threshold <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//[[1,1,1,1],[1,0,0,0],[1,0,0,0],[1,0,0,0]]
	fmt.Println(maxSideLength([][]int{{1,1,1,1}, {1,0,0,0}, {1,0,0,0}, {1,0,0,0}}, 6))
}

var result int

//超级遍历 ，只保留合格的，不合格的直接淘汰
func maxSideLength(mat [][]int, threshold int) int {
	result = 0
	m := len(mat)
	n := len(mat[0])

	//前缀和
	preS := make([][]int, m+1)
	for i := range preS {
		preS[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			preS[i][j] = preS[i-1][j] + preS[i][j-1] - preS[i-1][j-1] + mat[i-1][j-1]
		}
	}

	preT := make(map[int]map[int]bool)
	//一阶满足的数
	for i, iv := range mat {
		for j, jv := range iv {
			if jv <= threshold {
				if _, ok := preT[i]; !ok {
					preT[i] = make(map[int]bool)
				}
				preT[i][j] = true
			}
		}
	}
	getNextT(preT, 1, threshold, preS)
	return result
}

func getNextT(preT map[int]map[int]bool, level int, threshold int, preS [][]int) {
	if len(preT) == 0 {
		return
	}
	result = level
	T := make(map[int]map[int]bool)

	for i, iv := range preT {
		for j := range iv {
			if preT[i+1][j] && preT[i][j+1] && preT[i+1][j+1] {
				//A[x1..x2][y1..y2]
				x1 := i + 1
				x2 := i + 1 + level
				y1 := j + 1
				y2 := j + 1 + level

				sum := preS[x2][y2] - preS[x1-1][y2] - preS[x2][y1-1] + preS[x1-1][y1-1]
				if sum <= threshold {
					if _, ok := T[i]; !ok {
						T[i] = make(map[int]bool)
					}
					T[i][j] = true
				}
			}
		}
	}
	getNextT(T, level+1, threshold, preS)
}
