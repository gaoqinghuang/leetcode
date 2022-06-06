package main

import (
	"fmt"
	"sort"
)

//给你一个整数 n ，表示有 n 节课，课程编号从 1 到 n 。同时给你一个二维整数数组 relations ，其中 relations[j] = [prevCoursej, nextCoursej] ，表示课程 prevCoursej 必须在课程 nextCoursej 之前 完成（先修课的关系）。同时给你一个下标从 0 开始的整数数组 time ，其中 time[i] 表示完成第 (i+1) 门课程需要花费的 月份 数。
//
//请你根据以下规则算出完成所有课程所需要的 最少 月份数：
//
//如果一门课的所有先修课都已经完成，你可以在 任意 时间开始这门课程。
//你可以 同时 上 任意门课程 。
//请你返回完成所有课程所需要的 最少 月份数。
//
//注意：测试数据保证一定可以完成所有课程（也就是先修课的关系构成一个有向无环图）。
//
// 
//
//示例 1:
//
//
//
//输入：n = 3, relations = [[1,3],[2,3]], time = [3,2,5]
//输出：8
//解释：上图展示了输入数据所表示的先修关系图，以及完成每门课程需要花费的时间。
//你可以在月份 0 同时开始课程 1 和 2 。
//课程 1 花费 3 个月，课程 2 花费 2 个月。
//所以，最早开始课程 3 的时间是月份 3 ，完成所有课程所需时间为 3 + 5 = 8 个月。
//示例 2：
//
//
//
//输入：n = 5, relations = [[1,5],[2,5],[3,5],[3,4],[4,5]], time = [1,2,3,4,5]
//输出：12
//解释：上图展示了输入数据所表示的先修关系图，以及完成每门课程需要花费的时间。
//你可以在月份 0 同时开始课程 1 ，2 和 3 。
//在月份 1，2 和 3 分别完成这三门课程。
//课程 4 需在课程 3 之后开始，也就是 3 个月后。课程 4 在 3 + 4 = 7 月完成。
//课程 5 需在课程 1，2，3 和 4 之后开始，也就是在 max(1,2,3,7) = 7 月开始。
//所以完成所有课程所需的最少时间为 7 + 5 = 12 个月。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/parallel-courses-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。







func main() {
	//zeross := []int{0,1,3,4}
	//index := sort.Search(len(zeross), func(i int) bool {
	//	return zeross[i] >= 2
	//})
	//fmt.Println(append(zeross[0:index],append([]int{2},zeross[index:]...)...))
	//fmt.Println(minimumTime(3,[][]int{{1,3},{2,3}},[]int{3,2,5}))
	fmt.Println(minimumTime(5,[][]int{{1,5},{2,5},{3,5},{3,4},{4,5}},[]int{1,2,3,4,5}))
	//n = 3, relations = [[1,3],[2,3]], time = [3,2,5]

}

type zeros struct {
	time int
	index int
}


func minimumTime(n int, relations [][]int, time []int) int {
	inDegree := make(map[int]int,n)

	for i:=0;i<n;i++ {
		inDegree[i] = 0
	}

	relationsMap := make(map[int][]int,0)

	for _,relation := range relations {
		prev    :=relation[0]-1
		next    :=relation[1]-1
		inDegree[next]++
		relationsMap[prev] = append(relationsMap[prev],next)
	}


	result := 0
	zeross := make([]zeros,0)
	for key,num := range inDegree{
		if num != 0 {
			continue
		}
		zeross = append(zeross,zeros{
			time: time[key],
			index: key,
		})
	}

	//按时间排序
	sort.Slice(zeross, func(i, j int) bool {
		return zeross[i].time<zeross[j].time
	})

	for  {
		//每次取一个最小时间来完成，完成后更新，
		if len(zeross) == 0 {
			break
		}
		//取出最小的，完成
		now := zeross[0]
		zeross = zeross[1:]
		result = now.time
		for _,r := range relationsMap[now.index] {
			inDegree[r]--
			if inDegree[r] == 0 {//加入到队列中
				newTime := now.time+time[r]
				index := sort.Search(len(zeross), func(i int) bool {
					return zeross[i].time >= newTime
				})
				zeross = append(zeross[0:index],append([]zeros{{time:newTime, index:r}},zeross[index:]...)...)
			}

		}

	}


	return result
}
