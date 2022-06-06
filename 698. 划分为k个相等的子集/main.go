package main

import (
	"fmt"
	"sort"
)

//给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
//
// 
//
//示例 1：
//
//输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
//输出： True
//说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
//示例 2:
//
//输入: nums = [1,2,3,4], k = 3
//输出: false
// 
//
//提示：
//
//1 <= k <= len(nums) <= 16
//0 < nums[i] < 10000
//每个元素的频率在 [1,4] 范围内
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/partition-to-k-equal-sum-subsets
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1},4))
}


func canPartitionKSubsets(nums []int, k int) bool {
	//从大到小
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	sum := 0
	for _,num := range nums {
		sum+=num
	}
	if sum %k !=0{
		return false
	}
	target := sum/k

	bucket :=make([]int,k)

	return backtrack(nums, 0, bucket, k, target)
}

func backtrack(nums []int,index int,bucket []int,k int,target int ) bool {
	if index == len(nums){
		return true
	}
	for  i := 0; i < k; i++ {

		// 剪枝：放入球后超过 target 的值，选择一下桶
		if bucket[i]+nums[index] > target {
			continue
		}

		if i > 0 && bucket[i] == bucket[i - 1] {
			continue
		}

		// 做选择：放入 i 号桶
		bucket[i] += nums[index]
		if backtrack(nums, index + 1, bucket, k, target) {
			return true
		}
		bucket[i] -= nums[index]
	}

	return false
}

