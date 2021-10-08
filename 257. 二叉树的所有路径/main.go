package main

import "fmt"

//给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
//叶子节点 是指没有子节点的节点。
//
//
//示例 1：
//
//
//输入：root = [1,2,3,null,5]
//输出：["1->2->5","1->3"]
//示例 2：
//
//输入：root = [1]
//输出：["1"]
//
//
//提示：
//
//树中节点的数目在范围 [1, 100] 内
//-100 <= Node.val <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-paths
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

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
	//a.Right.Left = &TreeNode{Val: 4}
	fmt.Println(binaryTreePaths(a))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []string

func binaryTreePaths(root *TreeNode) []string {
	result = make([]string, 0)

	getNextPath(fmt.Sprintf("%d", root.Val), root)
	return result
}

func getNextPath(path string, node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		result = append(result, path)
	}
	if node.Left != nil {
		getNextPath(fmt.Sprintf("%s->%d", path, node.Left.Val), node.Left)
	}
	if node.Right != nil {
		getNextPath(fmt.Sprintf("%s->%d", path, node.Right.Val), node.Right)
	}
}
