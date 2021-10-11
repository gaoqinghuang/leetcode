package main

import "fmt"

//给你一个 m x n 的网格 grid。网格里的每个单元都代表一条街道。grid[i][j] 的街道可以是：
//
//1 表示连接左单元格和右单元格的街道。
//2 表示连接上单元格和下单元格的街道。
//3 表示连接左单元格和下单元格的街道。
//4 表示连接右单元格和下单元格的街道。
//5 表示连接左单元格和上单元格的街道。
//6 表示连接右单元格和上单元格的街道。
//
//
//你最开始从左上角的单元格 (0,0) 开始出发，网格中的「有效路径」是指从左上方的单元格 (0,0) 开始、一直到右下方的 (m-1,n-1) 结束的路径。该路径必须只沿着街道走。
//
//注意：你 不能 变更街道。
//
//如果网格中存在有效的路径，则返回 true，否则返回 false 。
//
//
//
//示例 1：
//
//
//
//输入：grid = [[2,4,3],[6,5,2]]
//输出：true
//解释：如图所示，你可以从 (0, 0) 开始，访问网格中的所有单元格并到达 (m - 1, n - 1) 。
//示例 2：
//
//
//
//输入：grid = [[1,2,1],[1,2,1]]
//输出：false
//解释：如图所示，单元格 (0, 0) 上的街道没有与任何其他单元格上的街道相连，你只会停在 (0, 0) 处。
//示例 3：
//
//输入：grid = [[1,1,2]]
//输出：false
//解释：你会停在 (0, 1)，而且无法到达 (0, 2) 。
//示例 4：
//
//输入：grid = [[1,1,1,1,1,1,3]]
//输出：true
//示例 5：
//
//输入：grid = [[2],[2],[2],[2],[2],[2],[6]]
//输出：true
//
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m, n <= 300
//1 <= grid[i][j] <= 6
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/check-if-there-is-a-valid-path-in-a-grid
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(hasValidPath([][]int{
		{6, 1, 3},
		{4, 1, 5},
	}))
}

func hasValidPath(grid [][]int) bool {
	switch grid[0][0] {
	case 1:
		return find(grid, 3)
	case 2:
		return find(grid, 1)
	case 3:
		return find(grid, 3)
	case 4:
		return find(grid, 2) || find(grid, 4)
	case 5:
		return false
	case 6:
		return find(grid, 1)
	}
	return false
}

func find(grid [][]int, eType int) bool {
	//每个数字只有唯一的目的地，记录下哪些访问过了就好
	visitEd := make([][]bool, len(grid))
	for k := range visitEd {
		visitEd[k] = make([]bool, len(grid[0]))
	}
	startX := 0
	startY := 0
	for {
		//越界，循环
		if startX >= len(grid) || startY >= len(grid[0]) || startX < 0 || startY < 0 || visitEd[startX][startY] {
			break
		}
		//到达终点而且要能对接上
		switch grid[startX][startY] {
		case 1:
			if !(eType == 3 || eType == 4) {
				return false
			}
		case 2:
			if !(eType == 1 || eType == 2) {
				return false
			}
		case 3:
			if !(eType == 3 || eType == 2) {
				return false
			}
		case 4:
			if !(eType == 4 || eType == 2) {
				return false
			}
		case 5:
			if !(eType == 1 || eType == 3) {
				return false
			}
		case 6:
			if !(eType == 1 || eType == 4) {
				return false
			}
		}
		//到达最后一个
		if startX == len(grid)-1 && startY == len(grid[0])-1 {
			return true
		}
		visitEd[startX][startY] = true
		//取得下一个值
		switch grid[startX][startY] {
		case 1:
			if eType == 3 {
				startY++
				eType = 3
			} else if eType == 4 {
				startY--
				eType = 4
			} else {
				return false
			}
		case 2:
			if eType == 1 {
				startX++
				eType = 1
			} else if eType == 2 {
				startX--
				eType = 2
			} else {
				return false
			}
		case 3:
			if eType == 3 {
				startX++
				eType = 1
			} else if eType == 2 {
				startY--
				eType = 4
			} else {
				return false
			}
		case 4:
			if eType == 4 {
				startX++
				eType = 1
			} else if eType == 2 {
				startY++
				eType = 3
			} else {
				return false
			}
		case 5:
			if eType == 1 {
				startY--
				eType = 4
			} else if eType == 3 {
				startX--
				eType = 2
			} else {
				return false
			}
		case 6:
			if eType == 1 {
				startY++
				eType = 3
			} else if eType == 4 {
				startX--
				eType = 2
			} else {
				return false
			}
		}
	}
	return false
}
