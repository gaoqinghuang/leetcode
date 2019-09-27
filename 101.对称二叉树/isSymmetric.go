//给定一个二叉树，检查它是否是镜像对称的。
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//1
/// \
//2   2
/// \ / \
//3  4 4  3
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//1
/// \
//2   2
//\   \
//3    3
//说明:
//
//如果你可以运用递归和迭代两种方法解决这个问题，会很加分。

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//[1,2,2,3,4,4,3]
	//1
	/// \
	//2   2
	/// \ / \
	//3  4 4  3
	a := &TreeNode{Val: 1}
	a.Left = &TreeNode{Val: 2}
	a.Right = &TreeNode{Val: 2}
	a.Left.Left = &TreeNode{Val: 3}
	a.Left.Right = &TreeNode{Val: 4}
	a.Right.Right = &TreeNode{Val: 3}
	a.Right.Left = &TreeNode{Val: 4}
	fmt.Println(isSymmetric(a))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return deepisSymmetric(root.Left, root.Right)
}
func deepisSymmetric(right, left *TreeNode) bool {
	if right == nil {
		if left == nil {
			return true
		} else {
			return false
		}
	}

	if left == nil {
		return false
	}

	if right.Val != left.Val {
		return false
	}

	if !deepisSymmetric(right.Right, left.Left) {
		return false
	}
	if !deepisSymmetric(right.Left, left.Right) {
		return false
	}

	return true

}

