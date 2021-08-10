package main

import (
    "fmt"
    "github.com/kr/pretty"
)

//给定一个二叉树

//struct Node {
//int val;
//Node *left;
//Node *right;
//Node *next;
//}
//填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//
//初始状态下，所有 next 指针都被设置为 NULL。
//
// 
//
//进阶：
//
//你只能使用常量级额外空间。
//使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
// 
//
//示例：
//
//
//
//输入：root = [1,2,3,4,5,null,7]
//输出：[1,#,2,3,#,4,5,7,#]
//解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化输出按层序遍历顺序（由 next 指针连接），'#' 表示每层的末尾。
// 
//
//提示：
//
//树中的节点数小于 6000
//-100 <= node.val <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//*
//* Definition for a Node.
type Node struct {
    Val   int
    Left  *Node
    Right *Node
    Next  *Node
}

func main() {
    a := &Node{
        Val: 1,
        Left: &Node{
            Val: 2,
            Left: &Node{
              Val: 4,
            },
            Right: &Node{
                Val: 5,
            },
        },
        Right: &Node{
            Val: 3,
            Right: &Node{
                Val: 7,
            },
        },
    }
    fmt.Printf("%# v", pretty.Formatter(connect(a)))
}

func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    forNode(root)

    return root
}

func forNode(n *Node) {
    //找到每一层的最左边
    if n == nil {
        return
    }

    //左边
    if n.Left != nil {
        setNext(n,n.Left)
        forNode(n.Left)
        return
    }

    //右边
    if n.Right != nil {
        setNext(n,n.Right)
        forNode(n.Right)
        return
    }

    //两边都空
    if n.Next != nil {
        forNode(n.Next)
    }
}

func setNext(pn, n *Node) { //这里是一层一层的
    if n == nil {
        return
    }

    //左孩子
    if pn.Left == n && pn.Right != nil{
        n.Next = pn.Right
        setNext(pn,n.Next)
        return
    }

    if pn.Next == nil {
        return
    }

    if pn.Next.Left != nil {
        n.Next = pn.Next.Left
        setNext(pn.Next,n.Next)
        return
    }
    if pn.Next.Right != nil {
        n.Next = pn.Next.Right
        setNext(pn.Next,n.Next)
        return
    }
    setNext(pn.Next, n)
}
