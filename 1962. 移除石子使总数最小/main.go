package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//给你一个整数数组 piles ，数组 下标从 0 开始 ，其中 piles[i] 表示第 i 堆石子中的石子数量。另给你一个整数 k ，请你执行下述操作 恰好 k 次：
//
//选出任一石子堆 piles[i] ，并从中 移除 floor(piles[i] / 2) 颗石子。
//注意：你可以对 同一堆 石子多次执行此操作。
//
//返回执行 k 次操作后，剩下石子的 最小 总数。
//
//floor(x) 为 小于 或 等于 x 的 最大 整数。（即，对 x 向下取整）。
//
// 
//
//示例 1：
//
//输入：piles = [5,4,9], k = 2
//输出：12
//解释：可能的执行情景如下：
//- 对第 2 堆石子执行移除操作，石子分布情况变成 [5,4,5] 。
//- 对第 0 堆石子执行移除操作，石子分布情况变成 [3,4,5] 。
//剩下石子的总数为 12 。
//示例 2：
//
//输入：piles = [4,3,6,7], k = 3
//输出：12
//解释：可能的执行情景如下：
//- 对第 2 堆石子执行移除操作，石子分布情况变成 [4,3,3,7] 。
//- 对第 3 堆石子执行移除操作，石子分布情况变成 [4,3,3,4] 。
//- 对第 0 堆石子执行移除操作，石子分布情况变成 [2,3,3,4] 。
//剩下石子的总数为 12 。
// 
//
//提示：
//
//1 <= piles.length <= 105
//1 <= piles[i] <= 104
//1 <= k <= 105
//通过次数6,763提交次数14,957
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/remove-stones-to-minimize-the-total
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//arr:= []int{1,2,3,11,5}
	//arr = sortI(arr)
	//fmt.Println(arr)
	fmt.Println(minStoneSum([]int{4,3,6,7},3    ))
}

func minStoneSum(piles []int, k int) (ans int) {
	h := &hp{piles}
	heap.Init(h)
	for ; k > 0; k-- {
		h.IntSlice[0] -= h.IntSlice[0] / 2
		heap.Fix(h, 0)
	}
	for _, v := range h.IntSlice {
		ans += v
	}
	return
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }

//
////minStoneSum 会超时，使用大根堆
//func minStoneSum(piles []int, k int) int {
//	//排序
//	l := len(piles)
//	sort.Ints(piles)
//	for i:=1;i<=k;i++ {
//		//减半
//		piles[l-1] = piles[l-1]-piles[l-1]/2
//		piles = sortI(piles)
//	}
//
//	//最终合计
//	sum:= 0
//	for _,pile := range piles {
//		sum+=pile
//	}
//	return sum
//}
//
//// 最后有问题
//func sortI(arr []int)  []int {
//	last := arr[len(arr)-1]
//	index := sort.SearchInts(arr[0:len(arr)-1],last)
//	arr = arr[0:len(arr)-1]
//	return append(arr[0:index],append([]int{last},arr[index:]...)...)
//}