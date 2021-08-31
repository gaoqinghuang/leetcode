package main

import (
    "fmt"
)

//在一个整数数组中，“峰”是大于或等于相邻整数的元素，相应地，“谷”是小于或等于相邻整数的元素。例如，在数组{5, 8, 4, 2, 3, 4, 6}中，{8, 6}是峰， {5, 2}是谷。现在给定一个整数数组，将该数组按峰与谷的交替顺序排序。
//
//示例:
//
//输入: [5, 3, 1, 2, 3]
//输出: [5, 1, 3, 2, 3]
//提示：
//
//nums.length <= 10000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/peaks-and-valleys-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
    nums := []int{5, 3, 1, 2, 3}

    wiggleSort(nums)

    fmt.Println(nums)
}

func wiggleSort(nums []int) {
    if len(nums) == 0 {
        return
    }

    for k, num := range nums {
        if k+1 == len(nums) {
            break
        }
        if k%2 == 0 {
            //偶 峰
            if num < nums[k+1] {

                nums[k], nums[k+1] = nums[k+1], nums[k]
            }

        } else {
            //奇  谷
            if num > nums[k+1] {
                nums[k], nums[k+1] = nums[k+1], nums[k]
            }
        }

    }

}
