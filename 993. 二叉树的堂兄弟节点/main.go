package main

//在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
//
//如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
//
//我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
//
//只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。
//
//
//
//示例 1：
//
//
//输入：root = [1,2,3,4], x = 4, y = 3
//输出：false
//示例 2：
//
//
//输入：root = [1,2,3,null,4,null,5], x = 5, y = 4
//输出：true
//示例 3：
//
//
//
//输入：root = [1,2,3,null,4], x = 2, y = 3
//输出：false
//
//
//提示：
//
//二叉树的节点数介于 2 到 100 之间。
//每个节点的值都是唯一的、范围为 1 到 100 的整数。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/cousins-in-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

}

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var xi, yi int
var paxi, payi *TreeNode

//深度优先搜索
func isCousins(root *TreeNode, x int, y int) bool {
	dfs(root, root.Left, 1, x, y)
	dfs(root, root.Right, 1, x, y)
	if xi == yi && paxi.Val != payi.Val {
		return true
	}

	return false
}

func dfs(pat, t *TreeNode, nowLevel int, x int, y int) {
	if t == nil {
		return
	}
	if t.Val == x {
		xi = nowLevel
		paxi = pat
	}

	if t.Val == y {
		yi = nowLevel
		payi = pat
	}
	dfs(t, t.Left, nowLevel+1, x, y)
	dfs(t, t.Right, nowLevel+1, x, y)
}
