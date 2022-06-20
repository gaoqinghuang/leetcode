package main

import "fmt"

//请你设计一个队列，支持在前，中，后三个位置的 push 和 pop 操作。
//
//请你完成 FrontMiddleBack 类：
//
//FrontMiddleBack() 初始化队列。
//void pushFront(int val) 将 val 添加到队列的 最前面 。
//void pushMiddle(int val) 将 val 添加到队列的 正中间 。
//void pushBack(int val) 将 val 添加到队里的 最后面 。
//int popFront() 将 最前面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
//int popMiddle() 将 正中间 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
//int popBack() 将 最后面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
//请注意当有 两个 中间位置的时候，选择靠前面的位置进行操作。比方说：
//
//将 6 添加到 [1, 2, 3, 4, 5] 的中间位置，结果数组为 [1, 2, 6, 3, 4, 5] 。
//从 [1, 2, 3, 4, 5, 6] 的中间位置弹出元素，返回 3 ，数组变为 [1, 2, 4, 5, 6] 。
// 
//
//示例 1：
//
//输入：
//["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle", "popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
//[[], [1], [2], [3], [4], [], [], [], [], []]
//输出：
//[null, null, null, null, null, 1, 3, 4, 2, -1]
//
//解释：
//FrontMiddleBackQueue q = new FrontMiddleBackQueue();
//q.pushFront(1);   // [1]
//q.pushBack(2);    // [1, 2]
//q.pushMiddle(3);  // [1, 3, 2]
//q.pushMiddle(4);  // [1, 4, 3, 2]
//q.popFront();     // 返回 1 -> [4, 3, 2]
//q.popMiddle();    // 返回 3 -> [4, 2]
//q.popMiddle();    // 返回 4 -> [2]
//q.popBack();      // 返回 2 -> []
//q.popFront();     // 返回 -1 -> [] （队列为空）
// 
//
//提示：
//
//1 <= val <= 109
//最多调用 1000 次 pushFront， pushMiddle， pushBack， popFront， popMiddle 和 popBack 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/design-front-middle-back-queue
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//a := []int{1,2,3,4,5,6}
	//fmt.Println(append(a[:2],a[3:]...))
	//return
	f := Constructor()
	f.PushFront(1)
	f.PushBack(2)
	f.PushMiddle(3)
	f.PushMiddle(4)
	fmt.Println(f.PopFront())
	fmt.Println(f.PopMiddle())
	fmt.Println(f.PopMiddle())
	fmt.Println(f.PopBack())
	fmt.Println(f.PopFront())
}



type FrontMiddleBackQueue struct {
	m []int
}


func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{m:make([]int,0)}
}


func (this *FrontMiddleBackQueue) PushFront(val int)  {
	this.m = append([]int{val},this.m...)
}


func (this *FrontMiddleBackQueue) PushMiddle(val int)  {
	index := len(this.m)/2

	this.m = append(this.m[0:index],append([]int{val},this.m[index:]...)...)
}


func (this *FrontMiddleBackQueue) PushBack(val int)  {
	this.m = append(this.m,val)
}


func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.m) == 0 {
		return -1
	}
	front := this.m[0]
	this.m = this.m[1:]
	return front
}


func (this *FrontMiddleBackQueue) PopMiddle() int {
	if len(this.m) == 0 {
		return -1
	}
	index := (len(this.m)-1)/2
	val := this.m[index]
	this.m = append(this.m[:index],this.m[index+1:]...)
	return val
}


func (this *FrontMiddleBackQueue) PopBack() int {
	if len(this.m) == 0 {
		return -1
	}
	back := this.m[len(this.m)-1]
	this.m = this.m[:len(this.m)-1]
	return back
}


/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */