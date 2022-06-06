package main

import "encoding/json"

//
//序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
//
//设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。
//
//编码的字符串应尽可能紧凑。
//
// 
//
//示例 1：
//
//输入：root = [2,1,3]
//输出：[2,1,3]
//示例 2：
//
//输入：root = []
//输出：[]
// 
//
//提示：
//
//树中节点数范围是 [0, 104]
//0 <= Node.val <= 104
//题目数据 保证 输入的树是一棵二叉搜索树。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/serialize-and-deserialize-bst
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

}

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	a,_ :=  json.Marshal(root)
	return string(a)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == ""{
		return nil
	}
	t := &TreeNode{}
	json.Unmarshal([]byte(data),t)
	return t
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */