package main

import (
	"fmt"
	"sort"
)

//给你一个非负整数数组 nums 。如果存在一个数 x ，使得 nums 中恰好有 x 个元素 大于或者等于 x ，那么就称 nums 是一个 特殊数组 ，而 x 是该数组的 特征值 。
//
//注意： x 不必 是 nums 的中的元素。
//
//如果数组 nums 是一个 特殊数组 ，请返回它的特征值 x 。否则，返回 -1 。可以证明的是，如果 nums 是特殊数组，那么其特征值 x 是 唯一的 。
//
//
//
//示例 1：
//
//输入：nums = [3,5]
//输出：2
//解释：有 2 个元素（3 和 5）大于或等于 2 。
//示例 2：
//
//输入：nums = [0,0]
//输出：-1
//解释：没有满足题目要求的特殊数组，故而也不存在特征值 x 。
//如果 x = 0，应该有 0 个元素 >= x，但实际有 2 个。
//如果 x = 1，应该有 1 个元素 >= x，但实际有 0 个。
//如果 x = 2，应该有 2 个元素 >= x，但实际有 0 个。
//x 不能取更大的值，因为 nums 中只有两个元素。
//示例 3：
//
//输入：nums = [0,4,3,0,4]
//输出：3
//解释：有 3 个元素大于或等于 3 。
//示例 4：
//
//输入：nums = [3,6,7,7,0]
//输出：-1
//
//
//提示：
//
//1 <= nums.length <= 100
//0 <= nums[i] <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/special-array-with-x-elements-greater-than-or-equal-x
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(specialArray([]int{3}))
}

func specialArray(nums []int) int {
	if len(nums) == 1 {
		return -1
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	}) //从大到小排序
	for k, v := range nums {
		i := k + 1
		if v >= i {
			if i == len(nums) {
				return i
			} else {
				if nums[k+1] < i {
					return i
				}
			}
		}
	}

	return -1
}
