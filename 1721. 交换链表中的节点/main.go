package main

//给你链表的头节点 head 和一个整数 k 。
//
//交换 链表正数第 k 个节点和倒数第 k 个节点的值后，返回链表的头节点（链表 从 1 开始索引）。
//
//
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[1,4,3,2,5]
//示例 2：
//
//输入：head = [7,9,6,6,7,8,3,0,9,5], k = 5
//输出：[7,9,6,6,8,7,3,0,9,5]
//示例 3：
//
//输入：head = [1], k = 1
//输出：[1]
//示例 4：
//
//输入：head = [1,2], k = 1
//输出：[2,1]
//示例 5：
//
//输入：head = [1,2,3], k = 2
//输出：[1,2,3]
//
//
//提示：
//
//链表中节点的数目是 n
//1 <= k <= n <= 105
//0 <= Node.val <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/swapping-nodes-in-a-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						//Next: &ListNode{
						//	Val: 5,
						//},
					},
				},
			},
		},
	}
	swapNodes(list1, 2)

}

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	//简单暴力，直接找到2个位置，然后交换val值
	length := getLength(head)
	preVal := getValByNum(head, k)
	lastVal := getValByNum(head, length-k+1)
	changeValByNum(head, k, lastVal)
	changeValByNum(head, length-k+1, preVal)
	return head
}

func getLength(head *ListNode) int {
	length := 0
	now := head
	for now != nil {
		now = now.Next
		length++
	}
	return length
}

func getValByNum(head *ListNode, num int) int {
	length := 0
	now := head
	for now != nil {
		length++
		if num == length {
			return now.Val
		}
		now = now.Next
	}
	return length
}

func changeValByNum(head *ListNode, num, val int) {
	length := 0
	now := head
	for now != nil {
		length++
		if num == length {
			now.Val = val
			return
		}
		now = now.Next
	}
	return
}
