package main

import (
	"fmt"
)

//
//给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。
//
//请你找出层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。
//
//
//
//示例 1：
//
//
//
//输入：root = [1,7,0,7,-8,null,null]
//输出：2
//解释：
//第 1 层各元素之和为 1，
//第 2 层各元素之和为 7 + 0 = 7，
//第 3 层各元素之和为 7 + -8 = -1，
//所以我们返回第 2 层的层号，它的层内元素之和最大。
//示例 2：
//
//输入：root = [989,null,10250,98693,-89388,null,null,null,-32127]
//输出：2
//
//
//提示：
//
//树中的节点数介于 1 和 10^4 之间
//-10^5 <= node.val <= 10^5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//[1,2,2,3,4,4,3]
	//-100
	/// \
	//-200   -300
	/// \ / \
	//3  4 4  3
	a := &TreeNode{Val: -100}
	a.Left = &TreeNode{Val: -200}
	a.Right = &TreeNode{Val: -300}
	a.Left.Left = &TreeNode{Val: -20}
	a.Left.Right = &TreeNode{Val: -5}
	a.Right.Right = &TreeNode{Val: -10}
	//a.Right.Left = &TreeNode{Val: 4}
	fmt.Println(maxLevelSum(a))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//一层一层遍历，找到最大层
func maxLevelSum(root *TreeNode) int {
	max := root.Val
	level := 1
	maxLevel := 1
	list1 := make([]*TreeNode, 0)
	list1 = append(list1, root)

	for len(list1) != 0 {
		list2 := make([]*TreeNode, 0)
		sum := 0
		for _, v := range list1 {
			if v.Left != nil {
				sum += v.Left.Val
				list2 = append(list2, v.Left)
			}
			if v.Right != nil {
				sum += v.Right.Val
				list2 = append(list2, v.Right)
			}
		}
		level++
		if sum > max && len(list2) != 0 {
			max = sum
			maxLevel = level
		}
		list1 = list2

	}

	return maxLevel
}
