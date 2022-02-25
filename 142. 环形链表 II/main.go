package main

//给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
//不允许修改 链表。
//
// 
//
//示例 1：
//
//
//
//输入：head = [3,2,0,-4], pos = 1
//输出：返回索引为 1 的链表节点
//解释：链表中有一个环，其尾部连接到第二个节点。
//示例 2：
//
//
//
//输入：head = [1,2], pos = 0
//输出：返回索引为 0 的链表节点
//解释：链表中有一个环，其尾部连接到第一个节点。
//示例 3：
//
//
//
//输入：head = [1], pos = -1
//输出：返回 null
//解释：链表中没有环。
// 
//
//提示：
//
//链表中节点的数目范围在范围 [0, 104] 内
//-105 <= Node.val <= 105
//pos 的值为 -1 或者链表中的一个有效索引
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	a := a
}

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}


























//难点在于O(1)空间 直接hash
func detectCycle1(head *ListNode) *ListNode {
	had := make(map[*ListNode]bool)
	now :=head
	for  {
		if now == nil {
			return nil
		}
		if had[now] {
			return now
		}
		had[now]=true
		now = now.Next

	}
}

//快慢指针
func detectCycle(head *ListNode) *ListNode {
	x,y,now :=head,head,head
	for y!=nil {
		x = x.Next//慢指针
		//快指针
		if y.Next == nil {
			return nil
		}
		y = y.Next.Next
		if x ==y {
			//相遇
			for  {
				if now == y {
					return now
				}
				now = now.Next
				y = y.Next
			}

		}
	}
	return nil
}

