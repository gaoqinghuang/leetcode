package main

import (
	"fmt"
	"github.com/kr/pretty"
)

//
//给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。
//
//请你将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置。
//
//下图中蓝色边和节点展示了操作后的结果：
//
//
//请你返回结果链表的头指针。
//
//
//
//示例 1：
//
//
//
//输入：list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
//输出：[0,1,2,1000000,1000001,1000002,5]
//解释：我们删除 list1 中下标为 3 和 4 的两个节点，并将 list2 接在该位置。上图中蓝色的边和节点为答案链表。
//示例 2：
//
//
//输入：list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
//输出：[0,1,1000000,1000001,1000002,1000003,1000004,6]
//解释：上图中蓝色的边和节点为答案链表。
//
//
//提示：
//
//3 <= list1.length <= 104
//1 <= a <= b < list1.length - 1
//1 <= list2.length <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-in-between-linked-lists
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	list1 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
	}
	list2 := &ListNode{
		Val: 1000000,
		Next: &ListNode{
			Val: 1000001,
			Next: &ListNode{
				Val: 1000002,
			},
		},
	}
	fmt.Printf("%# v", pretty.Formatter(mergeInBetween(list1, 3, 4, list2)))
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

var i int
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	//a,b可能相等，不是开头，也不是末尾
	i = 1
	//找到a
	aList := find(list1, a)
	//找到下一个位置b

	bList := find(aList, b).Next.Next
	//对接上一个位置
	aList.Next = list2
	bLastList := getLastB(list2)

	bLastList.Next = bList
	//对接
	return list1
}

func find(list1 *ListNode, a int) *ListNode {
	//....跟val值无关啊,更位置有关啊···
	if i == a {
		return list1
	}
	i ++
	return find(list1.Next,a)
}

func getLastB(list2 *ListNode) *ListNode {
	if list2.Next == nil {
		return list2
	}
	return getLastB(list2.Next)
}
