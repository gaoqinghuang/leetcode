package main

import (
	"fmt"
	"sort"
)

//
//有一些工作：difficulty[i] 表示第 i 个工作的难度，profit[i] 表示第 i 个工作的收益。
//
//现在我们有一些工人。worker[i] 是第 i 个工人的能力，即该工人只能完成难度小于等于 worker[i] 的工作。
//
//每一个工人都最多只能安排一个工作，但是一个工作可以完成多次。
//
//举个例子，如果 3 个工人都尝试完成一份报酬为 1 的同样工作，那么总收益为 $3。如果一个工人不能完成任何工作，他的收益为 $0 。
//
//我们能得到的最大收益是多少？
//
//
//
//示例：
//
//输入: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
//输出: 100
//解释: 工人被分配的工作难度是 [4,4,6,6] ，分别获得 [20,20,30,30] 的收益。
//
//
//提示:
//
//1 <= difficulty.length = profit.length <= 10000
//1 <= worker.length <= 10000
//difficulty[i], profit[i], worker[i]  的范围是 [1, 10^5]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/most-profit-assigning-work
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(maxProfitAssignment([]int{5,50,92,21,24,70,17,63,30,53}, []int{68,100,3,99,56,43,26,93,55,25}, []int{96,3,55,30,11,58,68,36,26,1}))
}

type d struct {
	Difficulty int
	Profit int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	//排序
	jobs := make([]*d,len(difficulty))
	for k := range difficulty {
		jobs[k] = &d{
			Difficulty:difficulty[k],
			Profit:profit[k],
		}
	}
	sort.Ints(worker)
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Difficulty < jobs[j].Difficulty
	})
	result := 0

	var best, i int
	for _, w := range worker {
		for i < len(jobs) && w >= jobs[i].Difficulty   {
			best = maxInt(best, jobs[i].Profit)
			i++
		}
		result += best
	}

	return result
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

