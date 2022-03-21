package main

import "fmt"

//设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。这里的“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。
//
//注意：本题相对原题做了扩展
//
//示例:
//
//输入：4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释: 4 皇后问题存在如下两个不同的解法。
//[
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/eight-queens-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(solveNQueens(4))
}

var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	resTem := make([][]string,0)
	backtrack(0, n, resTem)
	return res
}

//first  第二行的了，第一个位置
func backtrack(first,n int,resTem [][]string)  {
	if len(resTem) == n{
		res = resTem
		return
	}
	if len(res)>0 {
		return
	}
	for i:= 0;i<n;i++ {//first第几行  i就是第几列了
		if len(res)>0 {
			return
		}
		if !check(resTem,i) {
			continue
		}
		re := make([]string,n)
		for k:=0;k<n ;k++ {
			re[k]= "."
		}
		re[i] = "Q"
		resTem = append(resTem,re)
		backtrack(first+1,n,resTem)
		if len(res)>0 {
			return
		}
		if first == len(resTem)-1{
			resTem = resTem[0:len(resTem)-1]
		}
	}
}

func check(resTem [][]string,i int) bool {
	if len(resTem) == 0 {
		return true
	}
	//相同列
	for _,re := range resTem {
		if re[i] == "Q"{
			return false
		}

	}
	//左上角
	if i != 0 {//一定有左上角
		if resTem[len(resTem)-1][i-1] == "Q" {
			return false
		}
	}
	//右上角
	if i != len(resTem[0])-1 {//一定有右上角
		if resTem[len(resTem)-1][i+1] == "Q" {
			return false
		}
	}
	return true
}