package main

import "fmt"

//给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
//请你找出符合题意的 最短 子数组，并输出它的长度。
//
// 
//
//示例 1：
//
//输入：nums = [2,6,4,8,10,9,15]
//输出：5
//解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//示例 2：
//
//输入：nums = [1,2,3,4]
//输出：0
//示例 3：
//
//输入：nums = [1]
//输出：0
// 
//
//提示：
//
//1 <= nums.length <= 104
//-105 <= nums[i] <= 105
// 
//
//进阶：你可以设计一个时间复杂度为 O(n) 的解决方案吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
    //for i:=1 ;i<1;i++ {
    //    fmt.Println(1111)
    //}
    fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
}

func findUnsortedSubarray(nums []int) int {
    if len(nums) == 0 || len(nums) == 1 {
        return 0
    }
    //处理小的
    //找出顺序
    minSortK := 0
    for k, num := range nums {
        //找出从小到大的最小值
        if k == 0 {
            continue
        }
        if num < nums[minSortK] {
            break
        }
        minSortK = k
    }

    //找出最小值
    minK := nums[minSortK]
    for k := minSortK + 1; k <= len(nums)-1; k++ {
        if minK > nums[k] {
            minK = nums[k]
        }
    }
    //顺序里的最小值

    lK := 0
    for k := 0; k <= minSortK; k++ {
        //找出从小到大的最小值
        if nums[k] > minK {
            break
        }
        lK++
    }

    //大

    maxSortK := len(nums) - 1
    for k := len(nums) - 2; k >= 0; k-- {
        //找出从大到小的最大值
        if nums[k] > nums[maxSortK] {
            break
        }
        maxSortK = k
    }

    //找出最大值
    maxK := nums[maxSortK]
    for k := 0; k <= maxSortK-1; k++ {
        if maxK < nums[k] {
            maxK = nums[k]
        }
    }
    //顺序里的最大值

    rK := 0
    for k := len(nums) - 1; k >= maxSortK; k-- {
        if nums[k] < maxK {
            break
        }
        rK++
    }
    if rK+lK > len(nums) {
        return 0
    }

    return len(nums) - rK - lK
}
