package main

import "fmt"

//给定一个数组 nums ，将其划分为两个连续子数组 left 和 right， 使得：
//
//left 中的每个元素都小于或等于 right 中的每个元素。
//left 和 right 都是非空的。
//left 的长度要尽可能小。
//在完成这样的分组后返回 left 的 长度 。
//
//用例可以保证存在这样的划分方法。
//
//
//
//示例 1：
//
//输入：nums = [5,0,3,8,6]
//输出：3
//解释：left = [5,0,3]，right = [8,6]
//示例 2：
//
//输入：nums = [1,1,1,0,6,12]
//输出：4
//解释：left = [1,1,1,0]，right = [6,12]
//
//
//提示：
//
//2 <= nums.length <= 105
//0 <= nums[i] <= 106
//可以保证至少有一种方法能够按题目所描述的那样对 nums 进行划分。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/partition-array-into-disjoint-intervals
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(partitionDisjoint([]int{6,0,8,30,37,6,75,98,39,90,63,74,52,92,64}))
}

//记录左边的最大，右边的最小，然后直接对比这两个数组
func partitionDisjoint(nums []int) int {
	maxL := make([]int, len(nums))
	minR := make([]int, len(nums))
	var max, min int
	for key, num := range nums {
		if key == 0 {
			max = num
		} else {
			max = maxInt(maxL[key-1], num)
		}
		maxL[key] = max
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if len(nums)-1 == i {
			min = nums[i]
		} else {
			min = minInt(minR[i+1], nums[i])
		}
		minR[i] = min
	}
	for i := 1; i < len(nums); i++ {
		if maxL[i-1] <= minR[i] {
			return i
		}
	}

	return 1
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
