package main

import (
	"fmt"
	"math"
	"sort"
)

//
//在一个长度 无限 的数轴上，第 i 颗石子的位置为 stones[i]。如果一颗石子的位置最小/最大，那么该石子被称作 端点石子 。
//
//每个回合，你可以将一颗端点石子拿起并移动到一个未占用的位置，使得该石子不再是一颗端点石子。
//
//值得注意的是，如果石子像 stones = [1,2,5] 这样，你将 无法 移动位于位置 5 的端点石子，因为无论将它移动到任何位置（例如 0 或 3），该石子都仍然会是端点石子。
//
//当你无法进行任何移动时，即，这些石子的位置连续时，游戏结束。
//
//要使游戏结束，你可以执行的最小和最大移动次数分别是多少？ 以长度为 2 的数组形式返回答案：answer = [minimum_moves, maximum_moves] 。
//
//
//
//示例 1：
//
//输入：[7,4,9]
//输出：[1,2]
//解释：
//我们可以移动一次，4 -> 8，游戏结束。
//或者，我们可以移动两次 9 -> 5，4 -> 6，游戏结束。
//示例 2：
//
//输入：[6,5,4,3,10]
//输出：[2,3]
//解释：
//我们可以移动 3 -> 8，接着是 10 -> 7，游戏结束。
//或者，我们可以移动 3 -> 7, 4 -> 8, 5 -> 9，游戏结束。
//注意，我们无法进行 10 -> 2 这样的移动来结束游戏，因为这是不合要求的移动。
//示例 3：
//
//输入：[100,101,104,102,103]
//输出：[0,0]
//
//
//提示：
//
//3 <= stones.length <= 10^4
//1 <= stones[i] <= 10^9
//stones[i] 的值各不相同。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/moving-stones-until-consecutive-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(numMovesStonesII([]int{7, 4, 9}))
	fmt.Println(numMovesStonesII([]int{6, 5, 4, 3, 10}))
	fmt.Println(numMovesStonesII([]int{100, 101, 104, 102, 103}))
}

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	//求最多
	length := len(stones)
	min := stones[0]
	max := stones[length-1]
	maxResult := max - min - length - minInt(stones[1]-min, max-stones[length-2]) + 2
	//lenI := len()
	//求最少
	start := 0
	startI := 0
	lenEmpit := math.MaxInt
	stones = append(stones,math.MaxInt)
	for index,stone := range stones {
		if index == 0 {
			start = stone
			startI = index
		}else {
			//这里可能是连着
			for stone >= start+length && startI<index{
				lenEmpit = minInt(lenEmpit,length-(index-startI))
				startI++
				start = stones[startI]
			}
		}
	}
	stones = stones[:length]

	//处理特殊情况
	if lenEmpit == 1 {
		//左右两边是否连着
		if (min+length-2 == stones[length-2] && stones[length-2]+2 != max) || (stones[1]+length-2 == stones[length-1] && min +2 != stones[1]){
			lenEmpit = 2
		}

	}


	return []int{lenEmpit, maxResult}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
