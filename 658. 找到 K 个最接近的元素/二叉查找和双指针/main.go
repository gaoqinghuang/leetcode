package main

import (
    "fmt"
    "math"
)

//给定一个排序好的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
//
//整数 a 比整数 b 更接近 x 需要满足：
//
//|a - x| < |b - x| 或者
//|a - x| == |b - x| 且 a < b
// 
//
//示例 1：
//
//输入：arr = [1,2,3,4,5], k = 4, x = 3
//输出：[1,2,3,4]
//示例 2：
//
//输入：arr = [1,2,3,4,5], k = 4, x = -1
//输出：[1,2,3,4]
// 
//
//提示：
//
//1 <= k <= arr.length
//1 <= arr.length <= 104
//数组里的每个元素与 x 的绝对值不超过 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-k-closest-elements
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
    arr := []int{1, 2, 3, 4, 5}
    k := 4
    x := -1
    fmt.Println(findClosestElements(arr, k, x))
}
func findClosestElements(arr []int, k int, x int) []int {
    for len(arr) > k {
        l, r := 0, len(arr)-1
        if math.Abs(float64(arr[l]-x)) <= math.Abs(float64(arr[r]-x)) {
            arr = arr[:r]
        } else {
            arr = arr[1:]
        }
    }
    return arr
}

func absInt(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
