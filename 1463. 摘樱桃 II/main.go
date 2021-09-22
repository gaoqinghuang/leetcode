package main

import "fmt"

//
//给你一个 rows x cols 的矩阵 grid 来表示一块樱桃地。 grid 中每个格子的数字表示你能获得的樱桃数目。
//
//你有两个机器人帮你收集樱桃，机器人 1 从左上角格子 (0,0) 出发，机器人 2 从右上角格子 (0, cols-1) 出发。
//
//请你按照如下规则，返回两个机器人能收集的最多樱桃数目：
//
//从格子 (i,j) 出发，机器人可以移动到格子 (i+1, j-1)，(i+1, j) 或者 (i+1, j+1) 。
//当一个机器人经过某个格子时，它会把该格子内所有的樱桃都摘走，然后这个位置会变成空格子，即没有樱桃的格子。
//当两个机器人同时到达同一个格子时，它们中只有一个可以摘到樱桃。
//两个机器人在任意时刻都不能移动到 grid 外面。
//两个机器人最后都要到达 grid 最底下一行。
//
//
//示例 1：
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/cherry-pickup-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

	//[[1,1],[1,1]]
	grid := make([][]int,0)
	grid  = append(grid,[]int{
		1,0,0,0,0,0,1,
	})
	grid  = append(grid,[]int{
		2,0,0,0,0,3,0,
	})
	grid  = append(grid,[]int{
		2,0,9,0,0,0,0,
	})
	grid  = append(grid,[]int{
		0,3,0,5,4,0,0,
	})
	grid  = append(grid,[]int{
		1,0,2,3,0,0,6,
	})
	fmt.Println(cherryPickup(grid))
}

var tems1 map[int]map[int]int
func cherryPickup(grid [][]int) int {
	//i比j小
	high := len(grid)
	length := len(grid[0])
	tems := make(map[int]map[int]int)
	for i := 0; i < high; i++ {
		tems[i] = make(map[int]int)
	}

	tems[0][len(grid[0])-1] = grid[0][len(grid[0])-1] + grid[0][0]//第一层
	for i := 1; i < high; i++ {
		tems1 = tems
		//tems = nil
		tems = makeTems1(high)
		//i 代表第几层
		//j 左边的数 k又边的数
		for j := 0; j < length; j++ {
			if j > i {
				continue
			}
			for k := j + 1; k < length; k++ {
				if j >= k {
					continue
				}
				if k < (length-i)-1 {
					continue
				}
				tems[j][k] = getMax(j,k,length,grid[i][j]+grid[i][k])
			}
		}

	}

	//最后一行累加
	result := 0
	for _,tem := range tems {
		for _,t := range tem{
			if t> result {
				result = t
			}
		}
	}
	return result
}
func makeTems1(high int) map[int]map[int]int{
	tems := make(map[int]map[int]int)
	for i := 0; i < high; i++ {
		tems[i] = make(map[int]int)
	}
	return tems
}

func getMax(j, k ,length,num int) int {
	result := 0
	for m:=j-1;m<j+2;m++{
		if m < 0 || m >=length {
			continue
		}
		for n:=k-1;n<k+2;n++ {
			if n < 0 || n >=length {
				continue
			}
			if m>=n {
				continue
			}
			if  tems1[m][n]+ num >result {
				result = tems1[m][n]+ num
			}
		}
	}


	return result
}
