package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
//
//例如，
//
//[2,3,4] 的中位数是 3
//
//[2,3] 的中位数是 (2 + 3) / 2 = 2.5
//
//设计一个支持以下两种操作的数据结构：
//
//void addNum(int num) - 从数据流中添加一个整数到数据结构中。
//double findMedian() - 返回目前所有元素的中位数。
//示例：
//
//addNum(1)
//addNum(2)
//findMedian() -> 1.5
//addNum(3)
//findMedian() -> 2
//进阶:
//
//如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
//如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-median-from-data-stream
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	m := Constructor()
	m.AddNum(1)
	fmt.Println(m.FindMedian())
	m.AddNum(2)
	fmt.Println(m.FindMedian())
	m.AddNum(3)
	fmt.Println(m.FindMedian())
	m.AddNum(4)
	fmt.Println(m.FindMedian())
	m.AddNum(5)
	fmt.Println(m.FindMedian())
}

//
//type MedianFinder struct {
//	Arr []int
//}
//
//func Constructor() MedianFinder {
//	return MedianFinder{Arr: make([]int, 0)}
//}
//
//func (this *MedianFinder) AddNum(num int) {
//	index := sort.SearchInts(this.Arr, num)
//	this.Arr = append(this.Arr[0:index], append([]int{num}, this.Arr[index:]...)...)
//}
//
//func (this *MedianFinder) FindMedian() float64 {
//	if len(this.Arr)&1 == 1 { //奇数
//		return float64(this.Arr[len(this.Arr)/2])
//	}
//	//偶数
//	return (float64(this.Arr[len(this.Arr)/2]) + float64(this.Arr[len(this.Arr)/2-1])) / 2
//}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type MedianFinder struct {
	queMin, queMax hp
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	minQ, maxQ := &mf.queMin, &mf.queMax
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
