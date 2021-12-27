package main

import "fmt"

//给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 K 。
//
//返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。
//
//
//
//示例 1：
//
//输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
//输出：[7,4,1]
//解释：
//所求结点为与目标结点（值为 5）距离为 2 的结点，
//值分别为 7，4，以及 1
//
//
//
//注意，输入的 "root" 和 "target" 实际上是树上的结点。
//上面的输入仅仅是对这些对象进行了序列化描述。
//
//
//提示：
//
//给定的树是非空的。
//树上的每个结点都具有唯一的值 0 <= node.val <= 500 。
//目标结点 target 是树上的结点。
//0 <= K <= 1000.
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//[1,2,2,3,4,4,3]
	//1
	///  \
	//2    3
	/// \  / \
	//     6
	//target := &TreeNode{Val: 1}
	a := &TreeNode{Val: 1}
	a.Left = &TreeNode{Val: 2}
	a.Right = &TreeNode{Val: 3}
	//a.Left.Left = &TreeNode{Val: 4}
	//a.Left.Right = &TreeNode{Val: 5}
	target := &TreeNode{Val: 6}
	a.Right.Right = target
	//target := &TreeNode{Val: 7}
	//a.Right.Left = target

	fmt.Println(distanceK(a, target, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	result = make([]int,0)
	parents = make(map[*TreeNode]*TreeNode)
	if k ==0 {
		return []int{target.Val}
	}
	//遍历
	setParents(root, root.Left)
	setParents(root, root.Right)
	distanceK1(target, target.Right, k-1)
	distanceK1(target, target.Left, k-1)
	distanceK1(target, parents[target], k-1)
	return result
}

var result []int
var parents map[*TreeNode]*TreeNode

func setParents(parent, target *TreeNode) {
	if target == nil {
		return
	}
	parents[target] = parent
	setParents(target, target.Left)
	setParents(target, target.Right)
}

func distanceK1(from, target *TreeNode, k int) {
	if target == nil {
		return
	}

	if k == 0 {
		result = append(result, target.Val)
		return
	}

	if parents[target] != from {
		distanceK1(target, parents[target], k-1)
	}
	if target.Left != from {
		distanceK1(target, target.Left, k-1)
	}
	if target.Right != from {
		distanceK1(target, target.Right, k-1)
	}
}



func distanceK2(root *TreeNode, target *TreeNode, k int) []int {
	var result []int
	var parents map[*TreeNode]*TreeNode

	parents = make(map[*TreeNode]*TreeNode)

	var setParents func(parent, target *TreeNode)
	setParents = func(parent, target *TreeNode) {
		if target == nil {
			return
		}
		parents[target] = parent
		setParents(target, target.Left)
		setParents(target, target.Right)
	}
	//遍历
	setParents(root, root.Left)
	setParents(root, root.Right)

	var distanceK1 func(from, target *TreeNode, k int)
	distanceK1 = func(from, target *TreeNode, k int) {
		if target == nil {
			return
		}

		if k == 0 {
			result = append(result, target.Val)
			return
		}

		if parents[target] != from {
			distanceK1(from, parents[target], k-1)
		}
		if target.Left != from {
			distanceK1(from, target.Left, k-1)
		}
		if target.Right != from {
			distanceK1(from, target.Right, k-1)
		}
	}
	distanceK1(target, target.Right, k-1)
	distanceK1(target, target.Left, k-1)
	distanceK1(target, parents[target], k-1)
	return result
}
