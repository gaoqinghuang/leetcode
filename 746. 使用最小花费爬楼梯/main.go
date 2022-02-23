package main

import "fmt"

//给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
//
//你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
//
//请你计算并返回达到楼梯顶部的最低花费。
//
//
//
//示例 1：
//
//输入：cost = [10,15,20]
//输出：15
//解释：你将从下标为 1 的台阶开始。
//- 支付 15 ，向上爬两个台阶，到达楼梯顶部。
//总花费为 15 。
//示例 2：
//
//输入：cost = [1,100,1,1,1,100,1,1,100,1]
//输出：6
//解释：你将从下标为 0 的台阶开始。
//- 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
//- 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
//- 支付 1 ，向上爬一个台阶，到达楼梯顶部。
//总花费为 6 。
//
//
//提示：
//
//2 <= cost.length <= 1000
//0 <= cost[i] <= 999
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/min-cost-climbing-stairs
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

//不用找什么规律，直接动态规划，每一层的最小值，等于上面一层或者上上层的最小值加上自己
func minCostClimbingStairs(cost []int) int {
	preA := cost[0]
	preB := cost[1]
	newPreB := 0
	for i := 2; i < len(cost); i++ {
		newPreB = min(preA, preB) + cost[i]
		preA = preB
		preB = newPreB
	}

	return min(preA, preB)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
