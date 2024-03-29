package main

import "fmt"

//在大小为 n x n 的网格 grid 上，每个单元格都有一盏灯，最初灯都处于 关闭 状态。
//
//给你一个由灯的位置组成的二维数组 lamps ，其中 lamps[i] = [rowi, coli] 表示 打开 位于 grid[rowi][coli] 的灯。即便同一盏灯可能在 lamps 中多次列出，不会影响这盏灯处于 打开 状态。
//
//当一盏灯处于打开状态，它将会照亮 自身所在单元格 以及同一 行 、同一 列 和两条 对角线 上的 所有其他单元格 。
//
//另给你一个二维数组 queries ，其中 queries[j] = [rowj, colj] 。对于第 j 个查询，如果单元格 [rowj, colj] 是被照亮的，则查询结果为 1 ，否则为 0 。在第 j 次查询之后 [按照查询的顺序] ，关闭 位于单元格 grid[rowj][colj] 上及相邻 8 个方向上（与单元格 grid[rowi][coli] 共享角或边）的任何灯。
//
//返回一个整数数组 ans 作为答案， ans[j] 应等于第 j 次查询 queries[j] 的结果，1 表示照亮，0 表示未照亮。
//
// 
//
//示例 1：
//
//
//输入：n = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,0]]
//输出：[1,0]
//解释：最初所有灯都是关闭的。在执行查询之前，打开位于 [0, 0] 和 [4, 4] 的灯。第 0 次查询检查 grid[1][1] 是否被照亮（蓝色方框）。该单元格被照亮，所以 ans[0] = 1 。然后，关闭红色方框中的所有灯。
//
//第 1 次查询检查 grid[1][0] 是否被照亮（蓝色方框）。该单元格没有被照亮，所以 ans[1] = 0 。然后，关闭红色矩形中的所有灯。
//
//示例 2：
//
//输入：n = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,1]]
//输出：[1,1]
//示例 3：
//
//输入：n = 5, lamps = [[0,0],[0,4]], queries = [[0,4],[0,1],[1,4]]
//输出：[1,1,0]
// 
//
//提示：
//
//1 <= n <= 109
//0 <= lamps.length <= 20000
//0 <= queries.length <= 20000
//lamps[i].length == 2
//0 <= rowi, coli < n
//queries[j].length == 2
//0 <= rowj, colj < n
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/grid-illumination
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(gridIllumination(5, [][]int{{0, 0}, {0, 4}}, [][]int{{0, 4}, {0, 1}, {1, 4}}))
}

type point struct {
	x int
	y int
}

//把下面的lampsMap变成5个变量，空间换时间,节约了循环lamps的时间
func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	result := make([]int, len(queries))
	points := make(map[point]bool)
	row := make(map[int]int)
	col := make(map[int]int)
	diagonal := make(map[int]int)
	antiDiagonal := make(map[int]int)

	for _, lamp := range lamps {
		x  := lamp[0]
		y  := lamp[1]

		p := point{
			x: x,
			y: y,
		}
		if _, ok := points[p];ok{
			continue
		}
		points[p] = true
		row[x]++
		col[y]++
		diagonal[x-y]++
		antiDiagonal[x+y]++
	}

	for i, querie := range queries {
		x, y := querie[0], querie[1]
		//判定是否亮的
		if row[x] > 0 || col[y] > 0 || diagonal[x-y] > 0 || antiDiagonal[x+y] > 0 {
			result[i] = 1
		}

		for a := x - 1; a <= x+1; a++ {
			for b := y - 1; b <= y+1; b++ {
				p := point{
					x:a,
					y:b,
				}
				if _, ok := points[p];!ok{
					continue
				}
				delete(points,p)
				row[a]--
				col[b]--
				diagonal[a-b]--
				antiDiagonal[a+b]--
			}
		}
	}

	return result
}

//此方法会超时，因为每次判定是否点亮时要循环len(lamps)
//func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
//	result := make([]int, len(queries))
//	lampsMap := make(map[int]map[int]bool)
//	for _, lamp := range lamps {
//		if _, ok := lampsMap[lamp[0]]; !ok {
//			lampsMap[lamp[0]] = make(map[int]bool)
//		}
//		lampsMap[lamp[0]][lamp[1]] = true
//	}
//
//	for i, querie := range queries {
//		a, b := querie[0], querie[1]
//		//判定是否亮的
//		result[i] = isLamps(lampsMap, a, b)
//
//		//关灯，lampsMap
//		closeLamps(lampsMap, a, b)
//	}
//
//	return result
//}
//
////1 表示照亮，0 表示未照亮。
//func isLamps(lamps map[int]map[int]bool, a, b int) int {
//	for i, lamp := range lamps {
//		for j := range lamp {
//			//横
//			if i == a {
//				return 1
//			}
//			//竖
//			if j == b {
//				return 1
//			}
//			//斜
//			if a-b == i-j {
//				return 1
//			}
//			if a+b == i+j {
//				return 1
//			}
//		}
//
//	}
//	return 0
//}
//
//func closeLamps(lamps map[int]map[int]bool, a, b int) {
//	//熄灭9个位置
//	delete(lamps[a], b)
//	delete(lamps[a], b+1)
//	delete(lamps[a], b-1)
//	delete(lamps[a+1], b)
//	delete(lamps[a+1], b+1)
//	delete(lamps[a+1], b-1)
//	delete(lamps[a-1], b)
//	delete(lamps[a-1], b+1)
//	delete(lamps[a-1], b-1)
//}