package main

import "fmt"

//在一个 m*n 的二维字符串数组中输出二叉树，并遵守以下规则：
//
//行数 m 应当等于给定二叉树的高度。
//列数 n 应当总是奇数。
//根节点的值（以字符串格式给出）应当放在可放置的第一行正中间。根节点所在的行与列会将剩余空间划分为两部分（左下部分和右下部分）。你应该将左子树输出在左下部分，右子树输出在右下部分。左下和右下部分应当有相同的大小。即使一个子树为空而另一个非空，你不需要为空的子树输出任何东西，但仍需要为另一个子树留出足够的空间。然而，如果两个子树都为空则不需要为它们留出任何空间。
//每个未使用的空间应包含一个空的字符串""。
//使用相同的规则输出子树。
//示例 1:
//
//输入:
//1
///
//2
//输出:
//[["", "1", ""],
//["2", "", ""]]
//示例 2:
//
//输入:
//1
/// \
//2   3
//\
//4
//输出:
//[["", "", "", "1", "", "", ""],
//["", "2", "", "", "", "3", ""],
//["", "", "4", "", "", "", ""]]
//示例 3:
//
//输入:
//1
/// \
//2   5
///
//3
///
//4
//输出:
//[["",  "",  "", "",  "", "", "", "1", "",  "",  "",  "",  "", "", ""]
//["",  "",  "", "2", "", "", "", "",  "",  "",  "",  "5", "", "", ""]
//["",  "3", "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]
//["4", "",  "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]]
//注意: 二叉树的高度在范围 [1, 10] 中。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/print-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
    a := &TreeNode{
        Val: 1,
        Left: &TreeNode{
            Val: 2,
            Left: &TreeNode{
                Val: 4,
            },
            Right: &TreeNode{
                Val: 5,
            },
        },
        Right: &TreeNode{
            Val: 3,
            Right: &TreeNode{
                Val: 7,
            },
        },
    }
    fmt.Println(printTree(a))
}

//Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

var result [][]string

func printTree(root *TreeNode) [][]string {
    if root == nil {
        return result
    }

    //获取层数
    height := GetTreeHeight(root)
    result = make([][]string, height)
    for k := range result {
        result[k] = make([]string, (2<<(height-1))-1)
    }
    //写入第一个位置

    setResult(height,1,1<<(height-1),root)

    return result
}

func setResult(sumHeight, height int, num int, tree *TreeNode) {
    if tree == nil {
        return
    }
    result[height-1][num-1] = fmt.Sprintf("%d", tree.Val)
    n := 1
    if sumHeight-height-1 >=0 {
        n = 1<<(sumHeight-height-1)

    }

    setResult(sumHeight, height+1, num-n, tree.Left)
    setResult(sumHeight, height+1, num+n, tree.Right)
}

//获取层数
func GetTreeHeight(node *TreeNode) int {
    if node == nil {
        return 0
    }
    lHight := GetTreeHeight(node.Left)
    rHight := GetTreeHeight(node.Right)
    if lHight >= rHight {
        return lHight + 1
    } else {
        return rHight + 1
    }
}
