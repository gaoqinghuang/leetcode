package main

import (
	"fmt"
	"sort"
)

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
//请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//
//
//示例 1:
//
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//示例 2:
//
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
//
//提示：
//
//1 <= k <= nums.length <= 104
//-104 <= nums[i] <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 6))
}

//利用快排，每次确定一位
func findKthLargest1(nums []int, k int) int {
	//选择一个数，左边比他大，右边比他小
	temNum := make([]int, len(nums))
	copy(temNum, nums)
	for {
		nowI, left, right := quickSort(temNum)
		if len(left)+1 == k {
			return nowI //选中的值
		} else if len(left)+1 > k { //在左边
			temNum = left
		} else { //在右边
			temNum = right
			k -= len(left) + 1
		}
	}

}

func quickSort(nums []int) (int, []int, []int) {
	selectI := nums[0]
	left, right := make([]int, 0), make([]int, 0)
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		if v > selectI {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return selectI, left, right
}

//排序直接获取第K位
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
