package main

import (
	"fmt"
	"math/rand"
)

//实现RandomizedSet 类：
//
//RandomizedSet() 初始化 RandomizedSet 对象
//bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
//bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
//int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
//你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。
//
// 
//
//示例：
//
//输入
//["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
//[[], [1], [2], [2], [], [1], [2], []]
//输出
//[null, true, false, true, 2, true, false, 2]
//
//解释
//RandomizedSet randomizedSet = new RandomizedSet();
//randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
//randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
//randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
//randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
//randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
//randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
//randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
// 
//
//提示：
//
//-231 <= val <= 231 - 1
//最多调用 insert、remove 和 getRandom 函数 2 * 105 次
//在调用 getRandom 方法时，数据结构中 至少存在一个 元素。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/insert-delete-getrandom-o1
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {


	//for i:=0;i<100;i++ {
	//
	//}
	//a := []int{1,2,3,4,5,6}
	//index := 5
	//a = append(a[0:index],a[index+1:]...)
	//a = append(a,1)
	//fmt.Println(a[0:len(a)-1])
	r:=Constructor()
	r.Insert(1)
	r.Remove(2)
	r.Insert(2)
	fmt.Println(r.GetRandom())
	r.Remove(1)
	r.Insert(2)
	fmt.Println(r.GetRandom())
}

type RandomizedSet struct {
	m map[int]int
	r []int
}


func Constructor() RandomizedSet {
	return RandomizedSet{
		m:make(map[int]int,0),
		r:make([]int,0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.r = append(this.r, val)
	this.m[val] = len(this.r) - 1
	return true
}


func (this *RandomizedSet) Remove(val int) bool {
	index,ok := this.m[val]
	if !ok{
		return false
	}
	delete(this.m,val)
	if index == len(this.r)-1 {
		this.r = this.r[0:len(this.r)-1]
	}else {
		//把最后一位移动过来，并删除最后一位
		this.r[index] = this.r[len(this.r)-1]
		this.m[this.r[index]] = index
	}
	return true
}


func (this *RandomizedSet) GetRandom() int {
	num := rand.Intn(len(this.r))
	return this.r[num]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */