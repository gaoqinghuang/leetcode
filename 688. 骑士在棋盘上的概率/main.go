package main

import (
	"fmt"
	"math"
)

//在一个 n x n 的国际象棋棋盘上，一个骑士从单元格 (row, column) 开始，并尝试进行 k 次移动。行和列是 从 0 开始 的，所以左上单元格是 (0,0) ，右下单元格是 (n - 1, n - 1) 。
//
//象棋骑士有8种可能的走法，如下图所示。每次移动在基本方向上是两个单元格，然后在正交方向上是一个单元格。
//
//
//
//每次骑士要移动时，它都会随机从8种可能的移动中选择一种(即使棋子会离开棋盘)，然后移动到那里。
//
//骑士继续移动，直到它走了 k 步或离开了棋盘。
//
//返回 骑士在棋盘停止移动后仍留在棋盘上的概率 。
//
//
//
//示例 1：
//
//输入: n = 3, k = 2, row = 0, column = 0
//输出: 0.0625
//解释: 有两步(到(1,2)，(2,1))可以让骑士留在棋盘上。
//在每一个位置上，也有两种移动可以让骑士留在棋盘上。
//骑士留在棋盘上的总概率是0.0625。
//示例 2：
//
//输入: n = 1, k = 0, row = 0, column = 0
//输出: 1.00000
//
//
//提示:
//
//1 <= n <= 25
//0 <= k <= 100
//0 <= row, column <= n
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/knight-probability-in-chessboard
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(float64(54617224235499260000)/float64(1237940039285380300000000000))
	//newS := make(map[int]map[int]int)
	//newS[1][2] = 1
	//fmt.Println(newS)
	//fmt.Println(knightProbability(8, 30, 6, 4))
}

func knightProbability(n int, k int, row int, column int) float64 {
	if k == 0 {
		return 1
	}
	count := math.Pow(8, float64(k)) //所有的可能性
	survivalCount := float64(0)      //存活的可能性

	//最后存活的位置
	lastS := make(map[int]map[int]int) //位置的次数
	lastS[row] = make(map[int]int)
	lastS[row][column] = 1

	newS := make(map[int]map[int]int) //新位置
	for i := 0; i < n; i++ {
		newS[i] = make(map[int]int)
	}
	for i := 1; i <= k; i++ {
		for row1, is := range lastS {
			for column1, v := range is {
				//8个方向
				row2 := row1 + 2
				column2 := column1 + 1
				if check(row2, column2, n) {
					newS[row2][column2] += v
				}

				row2 = row1 + 2
				column2 = column1 - 1
				if check(row2, column2, n) {
					newS[row2][column2] += v
				}

				row2 = row1 - 2
				column2 = column1 + 1
				if check(row2, column2, n) {
					newS[row2][column2] += v
				}

				row2 = row1 - 2
				column2 = column1 - 1
				if check(row2, column2, n) {
					newS[row2][column2] += v
				}

				row2 = row1 + 1
				column2 = column1 + 2
				if check(row2, column2, n) {
					newS[row2][column2] += v
				}

				row2 = row1 + 1
				column2 = column1 - 2
				if check(row2, column2, n) {
					newS[row2][column2] += v
				}

				row2 = row1 - 1
				column2 = column1 + 2
				if check(row2, column2, n) {
					newS[row2][column2] += v
				}

				row2 = row1 - 1
				column2 = column1 - 2
				if check(row2, column2, n) {
					newS[row2][column2] += v
				}
			}
		}
		lastS = newS
		newS = make(map[int]map[int]int)
		for ii := 0; ii < n; ii++ {
			newS[ii] = make(map[int]int)
		}
	}

	//合计
	for _, iv := range lastS {
		for _, v := range iv {
			survivalCount += float64(v)
		}
	}

	return survivalCount / count
}

//是否合法
func check(row, column, n int) bool {
	if row < 0 || row >= n {
		return false
	}

	if column < 0 || column >= n {
		return false
	}
	return true
}
