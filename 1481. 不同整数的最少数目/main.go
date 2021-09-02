package main

import (
    "fmt"
    "sort"
)

//给你一个整数数组 arr 和一个整数 k 。现需要从数组中恰好移除 k 个元素，请找出移除后数组中不同整数的最少数目。
//
// 
//
//示例 1：
//
//输入：arr = [5,5,4], k = 1
//输出：1
//解释：移除 1 个 4 ，数组中只剩下 5 一种整数。
//示例 2：
//
//输入：arr = [4,3,1,1,3,3,2], k = 3
//输出：2
//解释：先移除 4、2 ，然后再移除两个 1 中的任意 1 个或者三个 3 中的任意 1 个，最后剩下 1 和 3 两种整数。
// 
//
//提示：
//
//1 <= arr.length <= 10^5
//1 <= arr[i] <= 10^9
//0 <= k <= arr.length
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/least-number-of-unique-integers-after-k-removals
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
    fmt.Println(findLeastNumOfUniqueInts([]int{5,5,4},1))
}

type s struct {
    Num int //数量
    Count int //个数
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
    if len(arr) == 0 {
        return 0
    }
    result := 0
    //统计数量
    mapA := make(map[int]int)
    for _, v := range arr {
        mapA[v]++
    }
    mapB := make(map[int]int)
    //从最小的移除  数量，按从小到大排序，然后还有这个数量的数字
    for _,v := range mapA {
        if v > k {
            continue
        }
        mapB[v]++
    }
    result = len(mapA)
    mapA = nil
    sliceA := make([]s,0,len(mapB))
    for k,v := range mapB {
        sliceA = append(sliceA,s{
            Num: k,
            Count: v,
        })
    }
    mapB = nil
    sort.Slice(sliceA, func(i, j int) bool {
        return sliceA[i].Num < sliceA[j].Num
    })
    for _,v := range sliceA {
        if k <= 0 {
            break
        }
        if v.Count*v.Num >= k {
            //最多扣
            result -= k/v.Num
        }else {
            result -= v.Count
        }
        k -= v.Count*v.Num
    }
    return result
}
