package main

import (
	"fmt"
	"sort"
)

//一个序列的 宽度 定义为该序列中最大元素和最小元素的差值。
//
//给你一个整数数组 nums ，返回 nums 的所有非空 子序列 的 宽度之和 。由于答案可能非常大，请返回对 109 + 7 取余 后的结果。
//
//子序列 定义为从一个数组里删除一些（或者不删除）元素，但不改变剩下元素的顺序得到的数组。例如，[3,6,2,7] 就是数组 [0,3,1,6,2,2,7] 的一个子序列。
//
//
//
//示例 1：
//
//输入：nums = [2,1,3]
//输出：6
//解释：子序列为 [1], [2], [3], [2,1], [2,3], [1,3], [2,1,3] 。
//相应的宽度是 0, 0, 0, 1, 1, 2, 2 。
//宽度之和是 6 。
//示例 2：
//
//输入：nums = [2]
//输出：0
//
//
//提示：
//
//1 <= nums.length <= 105
//1 <= nums[i] <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-subsequence-widths
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//fmt.Println(math.Pow(2,2),math.Pow(2,0))
	fmt.Println(sumSubseqWidths([]int{96,87,191,197,40,101,108,35,169,50,168,182,95,80,144,43,18,60,174,13,77,173,38,46,80,117,13,19,11,6,13,118,39,80,171,36,86,156,165,190,53,49,160,192,57,42,97,35,124,200,84,70,145,180,54,141,159,42,66,66,25,95,24,136,140,159,71,131,5,140,115,76,151,137,63,47,69,164,60,172,153,183,6,70,40,168,133,45,116,188,20,52,70,156,44,27,124,59,42,172}))
}

//排好序后一个数要嘛在左边要嘛在右边，在左边贡献的是负值，在右边贡献的是正值
func sumSubseqWidths(nums []int) int  {
	sort.Ints(nums)
	n := len(nums)
	m, sum := 1000000007, 0
	// 这里可以提前计算好幂运算
	pow := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			pow[i] = 1
		} else {
			pow[i] = pow[i-1] * 2 % m
		}
	}

	for i := 0; i < n; i++ {
		sum = sum + (pow[i]-pow[n-i-1])*nums[i]
		sum = sum % m
	}
	return sum
}



//这个会超时
//两个数，左边前没有，右边有全都没有，其它无所谓有没有
func sumSubseqWidths1(nums []int) int {
	mod := int(1000000007)
	//一共 2^len(nums)-1  全部都不见的，没有
	result := int(0)
	//从小到大
	sort.Ints(nums)
	pow := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			pow[i] = 1
		} else {
			pow[i] = pow[i-1] * 2 % mod
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			result += pow[j-i-1] * (nums[j] - nums[i])%mod
		}
	}

	return result % mod
}