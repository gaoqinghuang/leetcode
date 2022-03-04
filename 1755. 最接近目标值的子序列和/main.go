package main

import (
	"fmt"
	"sort"
)

//给你一个整数数组 nums 和一个目标值 goal 。
//
//你需要从 nums 中选出一个子序列，使子序列元素总和最接近 goal 。也就是说，如果子序列元素和为 sum ，你需要 最小化绝对差 abs(sum - goal) 。
//
//返回 abs(sum - goal) 可能的 最小值 。
//
//注意，数组的子序列是通过移除原始数组中的某些元素（可能全部或无）而形成的数组。
//
// 
//
//示例 1：
//
//输入：nums = [5,-7,3,5], goal = 6
//输出：0
//解释：选择整个数组作为选出的子序列，元素和为 6 。
//子序列和与目标值相等，所以绝对差为 0 。
//示例 2：
//
//输入：nums = [7,-9,15,-2], goal = -5
//输出：1
//解释：选出子序列 [7,-9,-2] ，元素和为 -4 。
//绝对差为 abs(-4 - (-5)) = abs(1) = 1 ，是可能的最小值。
//示例 3：
//
//输入：nums = [1,2,3], goal = -7
//输出：7
// 
//
//提示：
//
//1 <= nums.length <= 40
//-107 <= nums[i] <= 107
//-109 <= goal <= 109
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/closest-subsequence-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(minAbsDifference([]int{-5862,-3226,-5358,7874,-8960,356,8119,-3269,-4163,-205},-40729811))
	//fmt.Println(minAbsDifference([]int{5,-7,3,5},6))
}













//分两次
func minAbsDifference(nums []int, goal int) int {
	result := abs(goal)
	a := all(nums[0:len(nums)/2])
	b := all(nums[len(nums)/2:])
	sort.Ints(a)
	sort.Ints(b)
	nowSum := 0
	for _,v := range a {
		nowSum = abs(v-goal)
		if nowSum < result {
			result = nowSum
		}
	}
	for _,v := range b {
		nowSum = abs(v-goal)
		if nowSum < result {
			result = nowSum
		}
	}
	i := 0
	j := len(b)-1
	for  {
		nowSum = abs(a[i]+b[j]-goal)
		if nowSum < result {
			result = nowSum
		}

		if a[i]+b[j] > goal{
			j--
		}else {
			i++
		}
		if i == len(a) || j == -1{
			break
		}
	}
	return result
}

//暴力，加记住上一次的选择
func all(nums []int) []int {
	sums := []int{0}
	nowSum := 0
	for _,num := range nums {
		for _,sum := range sums {
			nowSum = sum+num
			sums = append(sums,nowSum)
		}
	}
	return sums
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//暴力循环，占用空间太大
func minAbsDifference1(nums []int, goal int) int {
	sort.Ints(nums)//从小到大
	minAbs,nowSum := abs(goal),0
	sums := make([]int,1)
	for _,num := range nums {
		for _,sum := range sums {
			nowSum = sum+num
			if nowSum == goal {
				return 0
			}

			if minAbs> abs(nowSum-goal){
				minAbs = abs(nowSum-goal)
			}

			sums = append(sums,nowSum)
		}
	}
	return minAbs
}