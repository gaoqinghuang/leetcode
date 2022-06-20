package main

//请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
//实现 LRUCache 类：
//LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
// 
//
//示例：
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
// 
//
//提示：
//
//1 <= capacity <= 3000
//0 <= key <= 10000
//0 <= value <= 105
//最多调用 2 * 105 次 get 和 put
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/lru-cache
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



func main() {

}

type node struct {
	key,value int
	prev,next *node
}

type LRUCache struct {
	n int
	m map[int]*node
	head,tail *node
}

func Constructor(capacity int) LRUCache {
	 l := LRUCache{
		n:capacity,
		m:make(map[int]*node),
		head: &node{},
		tail: &node{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}


func (this *LRUCache) Get(key int) int {
	value,ok := this.m[key]
	if !ok {
		return -1
	}
	this.moveToHead(value)
	return value.value
}


func (this *LRUCache) Put(key int, value int)  {
	//有
	val,ok := this.m[key]
	if ok {
		//改变他的值，并移动到队首
		val.value = value
		this.moveToHead(val)
		return
	}

	no := &node{
		key: key,
		value: value,
	}
	this.m[key] = no
	this.addToHead(no)
	if len(this.m) > this.n {
		//超过了，找到队尾，删除链表和hash
		tail := this.removeTail()
		delete(this.m,tail.key)
	}



}

func (this *LRUCache) moveToHead(no *node) {
	this.removeNode(no)
	this.addToHead(no)
}

func (this *LRUCache) removeTail() *node {
	no := this.tail.prev
	this.removeNode(no)
	return no
}

func (this *LRUCache) removeNode(no *node) {
	no.prev.next = no.next
	no.next.prev = no.prev
}

func (this *LRUCache) addToHead(no *node) {
	no.prev = this.head
	no.next = this.head.next
	this.head.next.prev = no
	this.head.next = no
}



/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */