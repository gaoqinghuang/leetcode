package main

import "fmt"

//给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
//
//下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。
//
//
//
//示例 1：
//
//
//
//输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
//输出：13
//解释：如图所示，为和最小的两条下降路径
//示例 2：
//
//
//
//输入：matrix = [[-19,57],[-40,-5]]
//输出：-59
//解释：如图所示，为和最小的下降路径
//
//
//提示：
//
//n == matrix.length == matrix[i].length
//1 <= n <= 100
//-100 <= matrix[i][j] <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-falling-path-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	a := make([][]int, 0)
	a = append(a, []int{100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100})
	a = append(a, []int{100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100})
	//a = append(a, []int{69,33})
	//a = append(a, []int{7, 8, 9})
	fmt.Println(minFallingPathSum(a))
}

//遍历每一层最小的，算出下一层最小的，一层一层更新
func minFallingPathSum(matrix [][]int) int {
	preMin := make([]int, len(matrix[0]))
	newPreMin := make([]int, len(matrix[0]))
	var a, b, c int
	for _, ivalue := range matrix {
		for j, value := range ivalue {
			b = preMin[j] + value
			a = b+1
			c = b+1
			if j >= 1  {
				a = preMin[j-1] + value
			}
			if j+1 <= len(matrix[0])-1  {
				c = preMin[j+1] + value
			}
			fmt.Println(j,111)
			newPreMin[j] = min3(a, b, c)
		}
		copy(preMin,newPreMin)
	}

	return minArr(preMin)
}

func min3(a, b, c int) int {
	d := min(a, b)
	return min(c, d)
}


func minArr(arr []int) int {
	minValve := arr[0]
	for _,value := range arr{
		minValve = min(value,minValve)
	}
	return minValve
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
