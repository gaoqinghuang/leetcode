package main

//给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树不应该改变保留在树中的元素的相对结构（即，如果没有被移除，原有的父代子代关系都应当保留）。 可以证明，存在唯一的答案。
//
//所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。
//
// 
//
//示例 1：
//
//
//输入：root = [1,0,2], low = 1, high = 2
//输出：[1,null,2]
//示例 2：
//
//
//输入：root = [3,0,4,null,2,null,null,1], low = 1, high = 3
//输出：[3,2,null,1]
//示例 3：
//
//输入：root = [1], low = 1, high = 2
//输出：[1]
//示例 4：
//
//输入：root = [1,null,2], low = 1, high = 3
//输出：[1,null,2]
//示例 5：
//
//输入：root = [1,null,2], low = 2, high = 4
//输出：[2]
// 
//
//提示：
//
//树中节点数在范围 [1, 104] 内
//0 <= Node.val <= 104
//树中每个节点的值都是唯一的
//题目数据保证输入是一棵有效的二叉搜索树
//0 <= low <= high <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/trim-a-binary-search-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func main() {
    root := &TreeNode{
        Val: 3,
        Left: &TreeNode{
            Val: 0,
            Right: &TreeNode{
              Val: 2,
              Left: &TreeNode{
                  Val: 1,
              },
            },
        },
        Right: &TreeNode{
            Val: 4,
        },
    }
    trimBST(root, 1, 3)
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
    //low
    for root != nil && root.Val < low {
        root = root.Right
    }

    if root == nil {
        return nil
    }

    lowTree := root
    for lowTree.Left != nil {
        if lowTree.Left.Val < low {
            lowTree.Left = lowTree.Left.Right//todo
        } else {
            lowTree = lowTree.Left
        }
    }
    
    
    //high
    for root != nil && root.Val > high {
        root = root.Left
    }

    if root == nil {
        return nil
    }

    highTree := root
    for highTree.Right != nil {
        if highTree.Right.Val > high {
            highTree.Right = highTree.Right.Left//todo
        } else {
            highTree = highTree.Right
        }
    }

    return root
}
