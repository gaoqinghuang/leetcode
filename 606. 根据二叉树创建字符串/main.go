package main

import (
	"fmt"
)

//给你二叉树的根节点 root ，请你采用前序遍历的方式，将二叉树转化为一个由括号和整数组成的字符串，返回构造出的字符串。
//
//空节点使用一对空括号对 "()" 表示，转化后需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。
//
// 
//
//示例 1：
//
//
//输入：root = [1,2,3,4]
//输出："1(2(4))(3)"
//解释：初步转化后得到 "1(2(4)())(3()())" ，但省略所有不必要的空括号对后，字符串应该是"1(2(4))(3)" 。
//示例 2：
//
//
//输入：root = [1,2,3,null,4]
//输出："1(2()(4))(3)"
//解释：和第一个示例类似，但是无法省略第一个空括号对，否则会破坏输入与输出一一映射的关系。
// 
//
//提示：
//
//树中节点的数目范围是 [1, 104]
//-1000 <= Node.val <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/construct-string-from-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	a := &TreeNode{Val: 1}
	a.Left = &TreeNode{Val: 2}
	a.Right = &TreeNode{Val: 3}
	a.Left.Left = &TreeNode{Val: 4}
	fmt.Println(tree2str(a))
}

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil{
		return ""
	}else if root.Left == nil && root.Right == nil{
		return  fmt.Sprintf("%d",root.Val)
	}else if root.Right == nil {
		return fmt.Sprintf("%d(%s)",root.Val,tree2str(root.Left))
	}else {
		return fmt.Sprintf("%d(%s)(%s)",root.Val,tree2str(root.Left),tree2str(root.Right))
	}
}
