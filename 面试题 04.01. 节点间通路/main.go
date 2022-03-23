package main

import "fmt"

//节点间通路。给定有向图，设计一个算法，找出两个节点之间是否存在一条路径。
//
//示例1:
//
//输入：n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]], start = 0, target = 2
//输出：true
//示例2:
//
//输入：n = 5, graph = [[0, 1], [0, 2], [0, 4], [0, 4], [0, 1], [1, 3], [1, 4], [1, 3], [2, 3], [3, 4]], start = 0, target = 4
//输出 true
//提示：
//
//节点数量n在[0, 1e5]范围内。
//节点编号大于等于 0 小于 n。
//图中可能存在自环和平行边。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/route-between-nodes-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(findWhetherExistsPath(3,[][]int{{0,1},{0,2},{1,2},{1,2}},0,2))
}

var findMap map[int]bool
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	graphMap := make(map[int][]int)
	for _,gr := range graph {
		if gr[0] == gr[1] {
			continue
		}
		if _,ok := graphMap[gr[0]];!ok{
			graphMap[gr[0]] = make([]int,0)
		}
		graphMap[gr[0]]= append(graphMap[gr[0]],gr[1])
	}
	findMap = make(map[int]bool)
	return find(start,graphMap,target)
}

func find(start int, arrMap  map[int][]int,target int) bool {
	if findMap[start] {
		return false
	}
	findMap[start] = true
	if _,ok := arrMap[start]; !ok {
		return false
	}
	for _,v := range arrMap[start] {
		if v == target {
			return true
		}
		finded := find(v,arrMap,target)
		if finded {
			return true
		}
	}
	return false
}