package main

import "fmt"

//请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
//
//若队列为空，pop_front 和 max_value 需要返回 -1
//
//示例 1：
//
//输入:
//["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
//[[],[1],[2],[],[],[]]
//输出: [null,null,null,2,1,2]
//示例 2：
//
//输入:
//["MaxQueue","pop_front","max_value"]
//[[],[],[]]
//输出: [null,-1,-1]
// 
//
//限制：
//
//1 <= push_back,pop_front,max_value的总操作数 <= 10000
//1 <= value <= 10^5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。








func main() {
	maxQueue := Constructor()
	maxQueue.Push_back(1)
	maxQueue.Push_back(2)
	fmt.Println(maxQueue.Max_value())
	fmt.Println(maxQueue.Pop_front())
	fmt.Println(maxQueue.Max_value())
}

type MaxQueue struct {
	arr []int
	max int
}

func Constructor() MaxQueue {
	return MaxQueue{
		arr: make([]int, 0),
		max: -1,
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.arr) == 0 {
		return -1
	}
	return this.max
}

func (this *MaxQueue) Push_back(value int) {
	if value > this.max{
		this.max = value
	}
	this.arr = append(this.arr,value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.arr) == 0 {
		return 0
	}
	num := this.arr[0]
	this.arr = this.arr[1:]
	if num == this.max{
		this.max = getMax(this.arr)
	}
	return num
}

func getMax(arr []int)  int{
	max := 0
	for _,v := range arr {
		if v >max {
			max = v
		}
	}
	return max
}
/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
