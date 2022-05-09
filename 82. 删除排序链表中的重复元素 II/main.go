package main

import "fmt"

//
//给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
//
//
//
//示例 1：
//
//
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
//示例 2：
//
//
//输入：head = [1,1,1,2,3]
//输出：[2,3]
//
//
//提示：
//
//链表中节点数目在范围 [0, 300] 内
//-100 <= Node.val <= 100
//题目数据保证链表已经按升序 排列
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	fmt.Println(deleteDuplicates(l1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead *ListNode
	newNow := newHead
	lastVal := head.Val
	lastCount := 1

	now := head.Next
	for now != nil {
		if now.Val == lastVal {
			lastCount++
		} else {
			if lastCount == 1 {
				if newNow == nil {
					newNow = &ListNode{
						Val: lastVal,
					}
					newHead = newNow
				} else {
					new1 := &ListNode{
						Val: lastVal,
					}
					newNow.Next = new1
					newNow = new1
				}

			}

			lastVal = now.Val
			lastCount = 1
		}
		now = now.Next
	}

	//最后一个
	if lastCount == 1 {
		if newNow == nil {
			return &ListNode{
				Val: lastVal,
			}
		}
		newNow.Next = &ListNode{
			Val: lastVal,
		}
	}

	return newHead
}
