package main

import (
	"fmt"
)

//给你两个 非递增 的整数数组 nums1​​​​​​ 和 nums2​​​​​​ ，数组下标均 从 0 开始 计数。
//
//下标对 (i, j) 中 0 <= i < nums1.length 且 0 <= j < nums2.length 。如果该下标对同时满足 i <= j 且 nums1[i] <= nums2[j] ，则称之为 有效 下标对，该下标对的 距离 为 j - i​​ 。​​
//
//返回所有 有效 下标对 (i, j) 中的 最大距离 。如果不存在有效下标对，返回 0 。
//
//一个数组 arr ，如果每个 1 <= i < arr.length 均有 arr[i-1] >= arr[i] 成立，那么该数组是一个 非递增 数组。
//
// 
//
//示例 1：
//
//输入：nums1 = [55,30,5,4,2], nums2 = [100,20,10,10,5]
//输出：2
//解释：有效下标对是 (0,0), (2,2), (2,3), (2,4), (3,3), (3,4) 和 (4,4) 。
//最大距离是 2 ，对应下标对 (2,4) 。
//示例 2：
//
//输入：nums1 = [2,2,2], nums2 = [10,10,1]
//输出：1
//解释：有效下标对是 (0,0), (0,1) 和 (1,1) 。
//最大距离是 1 ，对应下标对 (0,1) 。
//示例 3：
//
//输入：nums1 = [30,29,19,5], nums2 = [25,25,25,25,25]
//输出：2
//解释：有效下标对是 (2,2), (2,3), (2,4), (3,3) 和 (3,4) 。
//最大距离是 2 ，对应下标对 (2,4) 。
//示例 4：
//
//输入：nums1 = [5,4], nums2 = [3,2]
//输出：0
//解释：不存在有效下标对，所以返回 0 。
// 
//
//提示：
//
//1 <= nums1.length <= 105
//1 <= nums2.length <= 105
//1 <= nums1[i], nums2[j] <= 105
//nums1 和 nums2 都是 非递增 数组
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-distance-between-a-pair-of-values
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
//[55,30,5,4,2], nums2 = [100,20,10,10,5]
	fmt.Println(maxDistance([]int{2,2,2},[]int{10,10,1}))

	//tem2 := []int{100,20,10,10,5}
	////tem2 := []int{5,10,10,20,100}
	//fmt.Println(Search(len(tem2), func(j int) bool { return tem2[j] >= 101 }))
}













func maxDistance(nums1 []int, nums2 []int) int {
	result := 0
	tem2:= make([]int,0)
	for i,num1 := range nums1 {
		if i > len(nums2)-1-result {
			break
		}
		tem2 = nums2[i:]
		f :=Search(len(tem2), func(j int) bool { return tem2[j] >= num1 })//找到比后面小的一个数
		if f == -1 {
			continue
		}
		result = max(result,f)
	}

	return result
}
func max(a,b int) int {
	if a<b {
		return b
	}
	return a
}

func Search(n int, f func(int) bool) int {
	i, j := -1, n-1
	for i < j {
		h := int(uint(i+j+1) >> 1)
		if !f(h) {
			j = h - 1
		} else {
			i = h
		}
	}
	return i
}