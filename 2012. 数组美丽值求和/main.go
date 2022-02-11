package main

import "fmt"

//
//给你一个下标从 0 开始的整数数组 nums 。对于每个下标 i（1 <= i <= nums.length - 2），nums[i] 的 美丽值 等于：
//
//2，对于所有 0 <= j < i 且 i < k <= nums.length - 1 ，满足 nums[j] < nums[i] < nums[k]
//1，如果满足 nums[i - 1] < nums[i] < nums[i + 1] ，且不满足前面的条件
//0，如果上述条件全部不满足
//返回符合 1 <= i <= nums.length - 2 的所有 nums[i] 的 美丽值的总和 。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：2
//解释：对于每个符合范围 1 <= i <= 1 的下标 i :
//- nums[1] 的美丽值等于 2
//示例 2：
//
//输入：nums = [2,4,6,4]
//输出：1
//解释：对于每个符合范围 1 <= i <= 2 的下标 i :
//- nums[1] 的美丽值等于 1
//- nums[2] 的美丽值等于 0
//示例 3：
//
//输入：nums = [3,2,1]
//输出：0
//解释：对于每个符合范围 1 <= i <= 1 的下标 i :
//- nums[1] 的美丽值等于 0
//
//
//提示：
//
//3 <= nums.length <= 105
//1 <= nums[i] <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-beauty-in-the-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(sumOfBeauties([]int{3, 13, 10, 3, 3, 12, 20, 16}))
}

func sumOfBeauties(nums []int) int {
	result := 0
	//或者右边最小
	lastLow := make([]int, len(nums))
	lastLow[len(nums)-2] = nums[len(nums)-1]
	for i := len(nums) - 3; i > 0; i-- {
		lastLow[i] = min(nums[i+1], lastLow[i+1])
	}

	preHigh := nums[0]
	for i := 1; i < len(nums)-1; i++ {

		//非上升
		if nums[i-1] >= nums[i] || nums[i+1] <= nums[i] {
			if preHigh < nums[i] {
				preHigh = nums[i]
			}
			continue
		}
		result++

		//是否再得一分
		if nums[i] > preHigh && nums[i] < lastLow[i] {
			result++
		}
		if preHigh < nums[i] {
			preHigh = nums[i]
		}
	}

	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
