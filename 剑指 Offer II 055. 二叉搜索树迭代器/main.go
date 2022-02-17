package main

import "fmt"

//实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
//
//BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
//boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
//int next()将指针向右移动，然后返回指针处的数字。
//注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。
//
//可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 的中序遍历中至少存在一个下一个数字。
//
//
//
//示例：
//
//
//
//输入
//inputs = ["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
//inputs = [[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
//输出
//[null, 3, 7, true, 9, true, 15, true, 20, false]
//
//解释
//BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
//bSTIterator.next();    // 返回 3
//bSTIterator.next();    // 返回 7
//bSTIterator.hasNext(); // 返回 True
//bSTIterator.next();    // 返回 9
//bSTIterator.hasNext(); // 返回 True
//bSTIterator.next();    // 返回 15
//bSTIterator.hasNext(); // 返回 True
//bSTIterator.next();    // 返回 20
//bSTIterator.hasNext(); // 返回 False
//
//
//提示：
//
//树中节点的数目在范围 [1, 105] 内
//0 <= Node.val <= 106
//最多调用 105 次 hasNext 和 next 操作
//
//
//进阶：
//
//你可以设计一个满足下述条件的解决方案吗？next() 和 hasNext() 操作均摊时间复杂度为 O(1) ，并使用 O(h) 内存。其中 h 是树的高度。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/kTOapQ
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//[1,2,2,3,4,4,3]
	//7
	/// \
	//3   15
	//    / \
	//    9  20
	a := &TreeNode{Val: 1}
	//a.Left = &TreeNode{Val: 3}
	//a.Right = &TreeNode{Val: 15}
	////a.Left.Left = &TreeNode{Val: 3}
	////a.Left.Right = &TreeNode{Val: 4}
	//a.Right.Right = &TreeNode{Val: 9}
	//a.Right.Left = &TreeNode{Val: 20}
	bSTIterator := Constructor(a)
	//fmt.Println(bSTIterator.Next())   // 返回 3
	//fmt.Println(bSTIterator.Next())    // 返回 7
	fmt.Println(bSTIterator.HasNext()) // 返回 True
	fmt.Println(bSTIterator.Next())    // 返回 9
	fmt.Println(bSTIterator.HasNext()) // 返回 True
	//fmt.Println(bSTIterator.Next())    // 返回 15
	//fmt.Println(bSTIterator.HasNext()) // 返回 True
	//fmt.Println(bSTIterator.Next())    // 返回 20
	//fmt.Println(bSTIterator.HasNext()) // 返回 False
}

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	result []int
	index int
}

func Constructor(root *TreeNode) BSTIterator {
	//直接中序遍历
	bSTIterator := BSTIterator{
	}
	bSTIterator.inorderTreeWalk(root)
	return bSTIterator
}


func (b *BSTIterator)inorderTreeWalk(t *TreeNode)  {
	if t != nil {
		b.inorderTreeWalk(t.Left)
		b.result = append(b.result,t.Val)
		b.inorderTreeWalk(t.Right)
	}
	return
}


func (this *BSTIterator) Next() int {
	num:= this.result[this.index]
	this.index++
	return num
}

func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.result)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

/**
 * 中序遍历排序
 * @param $x
 * @return array
 */
