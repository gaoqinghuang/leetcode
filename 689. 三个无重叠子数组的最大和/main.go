package main

import "fmt"

//给你一个整数数组 nums 和一个整数 k ，找出三个长度为 k 、互不重叠、且全部数字和（3 * k 项）最大的子数组，并返回这三个子数组。
//
//以下标的数组形式返回结果，数组中的每一项分别指示每个子数组的起始位置（下标从 0 开始）。如果有多个结果，返回字典序最小的一个。
//
// 
//
//示例 1：
//
//输入：nums = [1,2,1,2,6,7,5,1], k = 2
//输出：[0,3,5]
//解释：子数组 [1, 2], [2, 6], [7, 5] 对应的起始下标为 [0, 3, 5]。
//也可以取 [2, 1], 但是结果 [1, 3, 5] 在字典序上更大。
//示例 2：
//
//输入：nums = [1,2,1,2,1,2,1,2,1], k = 2
//输出：[0,2,4]
// 
//
//提示：
//
//1 <= nums.length <= 2 * 104
//1 <= nums[i] < 216
//1 <= k <= floor(nums.length / 3)
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/maximum-sum-of-3-non-overlapping-subarrays
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//fmt.Println([]int{0,1,2,3,4}[0:2])
	fmt.Println(maxSumOfThreeSubarrays([]int{1,2,1,2,6,7,5,1},2))
	fmt.Println(maxSumOfThreeSubarrays([]int{1,2,1,2,1,2,1,2,1},2))
}

//时间复杂度为n的
func maxSumOfThreeSubarrays(nums []int, k int) (ans []int) {
	var sum1, maxSum1, maxSum1Idx int
	var sum2, maxSum12, maxSum12Idx1,maxSum12Idx2 int
	var sum3, maxSum123 int
	for i := 2*k; i < len(nums); i++ {
		sum1 += nums[i-2*k]
		sum2 += nums[i-k]
		sum3 += nums[i]
		if i >= k*3-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				maxSum1Idx = i - k*3 + 1
			}

			if maxSum1+sum2 > maxSum12 {
				maxSum12 = maxSum1+sum2
				maxSum12Idx1, maxSum12Idx2 = maxSum1Idx, i-k*2+1
			}
			if maxSum12+sum3 > maxSum123 {
				maxSum123 = maxSum12 + sum3
				ans = []int{maxSum12Idx1, maxSum12Idx2, i - k + 1}
			}

			sum1 -= nums[i-k*3+1]
			sum2 -= nums[i-k*2+1]
			sum3 -= nums[i-k+1]
		}
	}
	return
}

func maxSumOfTwoSubarrays(nums []int, k int) (ans []int) {
	var sum1, maxSum1, maxSum1Idx int
	var sum2, maxSum12 int
	for i := k; i < len(nums); i++ {
		sum1 += nums[i-k]
		sum2 += nums[i]
		if i >= k*2-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				maxSum1Idx = i - k*2 + 1
			}
			if maxSum1+sum2 > maxSum12 {
				maxSum12 = maxSum1 + sum2
				ans = []int{maxSum1Idx, i - k + 1}
			}
			sum1 -= nums[i-k*2+1]
			sum2 -= nums[i-k+1]
		}
	}
	return
}

type d struct {
	index int //起始位置
	num int //数值
}

// n^2都超时
func maxSumOfThreeSubarrays1(nums []int, k int) []int {
	//分两个分割,3部分,分别计算哪个大

	//算出每段最大的
	maxD := make(map[int]map[int]*d,len(nums))//左右包含的
	lastSum := 0
	nextSum := sumSlice(nums[0:k])
	for i:= 0;i<=len(nums)-k;i++{
		lastSum = nextSum
		maxD[i] = make(map[int]*d)
		maxD[i][i+k-1] = &d{
			index: i,
			num: lastSum,
		}
		if i+k < len(nums){
			nextSum = nextSum+nums[i+k]-nums[i]
		}
		for j:= i+k;j<len(nums);j++{
			//上一个计数
			lastSum = lastSum+nums[j]-nums[j-k]
			if lastSum > maxD[i][j-1].num{
				maxD[i][j] = &d{
					num: lastSum,
					index: j-k+1,
				}
			}else {
				maxD[i][j] = maxD[i][j-1]
			}

		}
	}

	result := make([]int,3)
	maxSum := 0
	//左分割线
	for i:= k;i<=len(nums)-2*k;i++{
		//右分割线
		for j:= i+k;j<=len(nums)-k;j++{
			sum := maxD[0][i-1].num+maxD[i][j-1].num+maxD[j][len(nums)-1].num
			if sum > maxSum {
				result[0] = maxD[0][i-1].index
				result[1] = maxD[i][j-1].index
				result[2] = maxD[j][len(nums)-1].index
				maxSum = sum
			}
		}
	}

	return result
}


func sumSlice(nums []int)  int{
	sumNum := 0
	for _,num :=range nums{
		sumNum += num
	}
	return sumNum
}