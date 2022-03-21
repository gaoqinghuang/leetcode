package main

import (
	"fmt"
	"sort"
)

//给你一个整型数组 nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：6
//示例 2：
//
//输入：nums = [1,2,3,4]
//输出：24
//示例 3：
//
//输入：nums = [-1,-2,-3]
//输出：-6
//
//
//提示：
//
//3 <= nums.length <= 104
//-1000 <= nums[i] <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-product-of-three-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(maximumProduct([]int{722, 634, -504, -379, 163, -613, -842, -578, 750, 951, -158, 30, -238, -392, -487, -797, -157, -374, 999, -5, -521, -879, -858, 382, 626, 803, -347, 903, -205, 57, -342, 186, -736, 17, 83, 726, -960, 343, -984, 937, -758, -122, 577, -595, -544, -559, 903, -183, 192, 825, 368, -674, 57, -959, 884, 29, -681, -339, 582, 969, -95, -455, -275, 205, -548, 79, 258, 35, 233, 203, 20, -936, 878, -868, -458, -882, 867, -664, -892, -687, 322, 844, -745, 447, -909, -586, 69, -88, 88, 445, -553, -666, 130, -640, -918, -7, -420, -368, 250, -786}))
}

func maximumProduct(nums []int) int {
	n := len(nums)
	if n == 3 {
		return nums[0] * nums[1] * nums[2]
	}
	sort.Ints(nums)
	return max(nums[n-1]*nums[n-2]*nums[n-3], nums[0]*nums[1]*nums[n-1])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
