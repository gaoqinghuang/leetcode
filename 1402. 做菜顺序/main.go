package main

import (
	"fmt"
	"sort"
)

//一个厨师收集了他 n 道菜的满意程度 satisfaction ，这个厨师做出每道菜的时间都是 1 单位时间。
//
//一道菜的 「喜爱时间」系数定义为烹饪这道菜以及之前每道菜所花费的时间乘以这道菜的满意程度，也就是 time[i]*satisfaction[i] 。
//
//请你返回做完所有菜 「喜爱时间」总和的最大值为多少。
//
//你可以按 任意 顺序安排做菜的顺序，你也可以选择放弃做某些菜来获得更大的总和。
//
// 
//
//示例 1：
//
//输入：satisfaction = [-1,-8,0,5,-9]
//输出：14
//解释：去掉第二道和最后一道菜，最大的喜爱时间系数和为 (-1*1 + 0*2 + 5*3 = 14) 。每道菜都需要花费 1 单位时间完成。
//示例 2：
//
//输入：satisfaction = [4,3,2]
//输出：20
//解释：按照原来顺序相反的时间做菜 (2*1 + 3*2 + 4*3 = 20)
//示例 3：
//
//输入：satisfaction = [-1,-4,-5]
//输出：0
//解释：大家都不喜欢这些菜，所以不做任何菜可以获得最大的喜爱时间系数。
//示例 4：
//
//输入：satisfaction = [-2,5,-1,0,3,-3]
//输出：35
// 
//
//提示：
//
//n == satisfaction.length
//1 <= n <= 500
//-10^3 <= satisfaction[i] <= 10^3
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reducing-dishes
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。





func main() {
	//arr :=[]int{1,2,3}
	//i := sort.SearchInts(arr,2)
	//fmt.Println(i,arr[i:])
	fmt.Println(maxSatisfaction([]int{-2,5,-1,0,3,-3}))
}












func maxSatisfaction(satisfaction []int) int {
	//取大于0的部分，从小到大排序，然后一直加-数，如果有一次负数开始小于以前，则到此
	sort.Ints(satisfaction)
	zeroIndex := sort.SearchInts(satisfaction,0)//>=index 的都是
	nowSum := sum(satisfaction[zeroIndex:])
	if nowSum == 0 {
		return 0
	}
	result := sumIndex(satisfaction[zeroIndex:])
	if zeroIndex == 0 {
		return result
	}
	for i := zeroIndex-1; i >= 0; i-- {
		nowSum = nowSum+satisfaction[i]
		if nowSum > 0 {
			result +=nowSum
		}else {
			return result
		}
	}

	return result
}

func sumIndex(arr []int) int {
	result := 0
	for k,v := range arr {
		result += (k+1)*v
	}
	return result
}

func sum(arr []int) int {
	result := 0
	for _,v := range arr {
		result += v
	}
	return result
}