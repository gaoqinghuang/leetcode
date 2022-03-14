package main

import "fmt"

//树可以看成是一个连通且 无环 的 无向 图。
//
//给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。
//
//请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的边。
//
//
//
//示例 1：
//
//
//
//输入: edges = [[1,2], [1,3], [2,3]]
//输出: [2,3]
//示例 2：
//
//
//
//输入: edges = [[1,2], [2,3], [3,4], [1,4], [1,5]]
//输出: [1,4]
//
//
//提示:
//
//n == edges.length
//3 <= n <= 1000
//edges[i].length == 2
//1 <= ai < bi <= edges.length
//ai != bi
//edges 中无重复元素
//给定的图是连通的
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/redundant-connection
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(findRedundantConnection([][]int{[]int{1, 2}, {1, 3}, {2, 3}, {2, 3}, {2, 3}}))
}

//先连两边，无关链接的
func findRedundantConnection(edges [][]int) []int {
	//如果两个都没有，起一个新的，然后发现在某2个有的，则两个相连
	indexEdges := make(map[int]int, 0) //在第几号切片中
	sEdges := make([][]int, 0)
	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		aIndex, aok := indexEdges[a]
		bIndex, bok := indexEdges[b]
		//没在任何一个里，创建一个新的
		if !aok && !bok {
			sEdges = append(sEdges, []int{a, b})
			indexEdges[a] = len(sEdges) - 1
			indexEdges[b] = len(sEdges) - 1
		} else if aok && bok {
			//如果两个在同一个里，则形成环了
			//在两个里，两个合并
			if aIndex == bIndex { //这里就不对了
				return edge
			}
			//sEdges[bIndex]
			for _, v := range sEdges[bIndex] {
				indexEdges[v] = aIndex
			}
			sEdges[aIndex] = append(sEdges[aIndex], sEdges[bIndex]...)

		} else {
			//在其中一个里，加入进去
			if aok { //把b加入进来
				sEdges[aIndex] = append(sEdges[aIndex], b)
				indexEdges[b] = aIndex
			} else { //把a加入进来
				sEdges[bIndex] = append(sEdges[bIndex], a)
				indexEdges[a] = bIndex
			}
		}

	}
	return nil
}
