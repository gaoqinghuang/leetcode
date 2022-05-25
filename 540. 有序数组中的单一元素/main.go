package main

import "fmt"

//给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
//
//请你找出并返回只出现一次的那个数。
//
//你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
//
// 
//
//示例 1:
//
//输入: nums = [1,1,2,3,3,4,4,8,8]
//输出: 2
//示例 2:
//
//输入: nums =  [3,3,7,7,10,11,11]
//输出: 10
// 
//
//提示:
//
//1 <= nums.length <= 105
//0 <= nums[i] <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/single-element-in-a-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//fmt.Println(singleNonDuplicate([]int{1,1,2,3,3,4,4,8,8}))
	//fmt.Println(singleNonDuplicate([]int{3,3,7,7,10,11,11}))
	fmt.Println(singleNonDuplicate([]int{1,1,2,3,3}))
	//fmt.Println(10^1)
	//fmt.Println(9^1)
	//fmt.Println(8^1)
	//fmt.Println(7^1)
}

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	//取一半，看在左还是在右，一直这样
	index := len(nums)/2
	if index & 1 ==0 {
		if nums[index]== nums[index+1]{
			return singleNonDuplicate(nums[index:])
		}else {
			return singleNonDuplicate(nums[0:index+1])
		}
	}else {
		if nums[index]== nums[index-1]{
			return singleNonDuplicate(nums[index+1:])
		}else {
			return singleNonDuplicate(nums[0:index])
		}
	}
}
