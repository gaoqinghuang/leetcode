package main

import "fmt"

//给你一棵树（即，一个连通的无环无向图），这棵树由编号从 0  到 n - 1 的 n 个节点组成，且恰好有 n - 1 条 edges 。树的根节点为节点 0 ，树上的每一个节点都有一个标签，也就是字符串 labels 中的一个小写字符（编号为 i 的 节点的标签就是 labels[i] ）
//
//边数组 edges 以 edges[i] = [ai, bi] 的形式给出，该格式表示节点 ai 和 bi 之间存在一条边。
//
//返回一个大小为 n 的数组，其中 ans[i] 表示第 i 个节点的子树中与节点 i 标签相同的节点数。
//
//树 T 中的子树是由 T 中的某个节点及其所有后代节点组成的树。
//
//
//
//示例 1：
//
//
//
//输入：n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], labels = "abaedcd"
//输出：[2,1,1,1,1,1,1]
//解释：节点 0 的标签为 'a' ，以 'a' 为根节点的子树中，节点 2 的标签也是 'a' ，因此答案为 2 。注意树中的每个节点都是这棵子树的一部分。
//节点 1 的标签为 'b' ，节点 1 的子树包含节点 1、4 和 5，但是节点 4、5 的标签与节点 1 不同，故而答案为 1（即，该节点本身）。
//示例 2：
//
//
//
//输入：n = 4, edges = [[0,1],[1,2],[0,3]], labels = "bbbb"
//输出：[4,2,1,1]
//解释：节点 2 的子树中只有节点 2 ，所以答案为 1 。
//节点 3 的子树中只有节点 3 ，所以答案为 1 。
//节点 1 的子树中包含节点 1 和 2 ，标签都是 'b' ，因此答案为 2 。
//节点 0 的子树中包含节点 0、1、2 和 3，标签都是 'b'，因此答案为 4 。
//示例 3：
//
//
//
//输入：n = 5, edges = [[0,1],[0,2],[1,3],[0,4]], labels = "aabab"
//输出：[3,2,1,1,1]
//示例 4：
//
//输入：n = 6, edges = [[0,1],[0,2],[1,3],[3,4],[4,5]], labels = "cbabaa"
//输出：[1,2,1,1,2,1]
//示例 5：
//
//输入：n = 7, edges = [[0,1],[1,2],[2,3],[3,4],[4,5],[5,6]], labels = "aaabaaa"
//输出：[6,5,4,1,3,2,1]
//
//
//提示：
//
//1 <= n <= 10^5
//edges.length == n - 1
//edges[i].length == 2
//0 <= ai, bi < n
//ai != bi
//labels.length == n
//labels 仅由小写英文字母组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

//[[4,0],[5,4],[12,5],[3,12],[18,3],[10,18],[8,5],[16,8],[14,16],[13,16],[9,13],[22,9],[2,5],[6,2],[1,6],[11,1],[15,11],[20,11],[7,20],[19,1],[17,19],[23,19],[24,2],[21,24]]
//	"hcheiavadwjctaortvpsflssg"
	//输入：n = 5, edges = [[0,1],[0,2],[1,3],[0,4]], labels = "aabab"
	//输出：[3,2,1,1,1]
	//示例 4：
	fmt.Println(countSubTrees(25, [][]int{{4,0},{5,4},{12,5},{3,12},{18,3},{10,18},{8,5},{16,8},{14,16},{13,16},{9,13},{22,9},{2,5},{6,2},{1,6},{11,1},{15,11},{20,11},{7,20},{19,1},{17,19},{23,19},{24,2},{21,24}}, "hcheiavadwjctaortvpsflssg"))
	//fmt.Println(countSubTrees(4, [][]int{{0,2},{1,2},{0,3}}, "aeed"))
}

//通过edges深度搜索已有的值，然后比较是否相等,但是这样只能得到一个，这里是需要的全部，从头到尾相同数字只需要一次
func countSubTrees(n int, edges [][]int, labels string) []int {
	edgesMB := make(map[int][]int, 0)
	edgesM := make(map[int][]int, 0)

	for _, edge := range edges {
		k := edge[0]
		v := edge[1]

		if _, ok := edgesMB[k]; !ok {
			edgesMB[k] = make([]int, 0)
		}
		edgesMB[k] = append(edgesMB[k], v)
		k = edge[1]
		v = edge[0]
		if _, ok := edgesMB[k]; !ok {
			edgesMB[k] = make([]int, 0)
		}
		edgesMB[k] = append(edgesMB[k], v)
	}
	for _,v := range edgesMB[0] {
		if _, ok := edgesM[0]; !ok {
			edgesM[0] = make([]int, 0)
		}
		edgesM[0] = append(edgesM[0],v)
		getPositive(edgesMB,v,0,edgesM)
	}

	result := make([]int, n)

	//只知道根节点是0，如果从1开始算，那就不知道是多少了
	for i := 0; i < n; i++ {
		if result[i] >=1 {
			continue
		}
		result[i] = getChildLNum(edgesM, labels, labels[i], i,result)+1
	}

	return result
}

func getPositive (edgesMB map[int][]int,now,pre int,result map[int][]int)  {
	for _,v := range edgesMB[now] {
		if pre == v {
			continue
		}
		if _, ok := result[now]; !ok {
			result[now] = make([]int, 0)
		}
		result[now] = append(result[now],v)
		getPositive(edgesMB,v,now,result)
	}
}

//得用递归
func getChildLNum(edgesM map[int][]int, labels string, letter uint8, j int,overResult []int) int {
	if len(edgesM[j]) == 0 {
		return 0
	}
	re := 0
	for _, v := range edgesM[j] {
		if labels[v] == letter {
			overResult[v] = getChildLNum(edgesM, labels, letter, v,overResult)+1
			re += overResult[v]
		}else {
			re += getChildLNum(edgesM, labels, letter, v,overResult)
		}
		//
		//if labels[v] == letter {
		//	re++
		//}
		//re += getChildLNum(edgesM, labels, letter, v,overResult)
	}
	return re
}
