package main

import "fmt"

//给你一个二维字符网格数组 grid ，大小为 m x n ，你需要检查 grid 中是否存在 相同值 形成的环。
//
//一个环是一条开始和结束于同一个格子的长度 大于等于 4 的路径。对于一个给定的格子，你可以移动到它上、下、左、右四个方向相邻的格子之一，可以移动的前提是这两个格子有 相同的值 。
//
//同时，你也不能回到上一次移动时所在的格子。比方说，环  (1, 1) -> (1, 2) -> (1, 1) 是不合法的，因为从 (1, 2) 移动到 (1, 1) 回到了上一次移动时的格子。
//
//如果 grid 中有相同值形成的环，请你返回 true ，否则返回 false 。
//
// 
//
//示例 1：
//
//
//
//输入：grid = [['a','a','a','a'],['a','b','b','a'],['a','b','b','a'],['a','a','a','a']]
//输出：true
//解释：如下图所示，有 2 个用不同颜色标出来的环：
//
//示例 2：
//
//
//
//输入：grid = [['c','c','c','a'],['c','d','c','c'],['c','c','e','c'],['f','c','c','c']]
//输出：true
//解释：如下图所示，只有高亮所示的一个合法环：
//
//示例 3：
//
//
//
//输入：grid = [['a','b','b'],['b','z','b'],['b','b','a']]
//输出：false
// 
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m <= 500
//1 <= n <= 500
//grid 只包含小写英文字母。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/detect-cycles-in-2d-grid
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
    //alrMap := make(map[int]map[int]bool)
    //alrMap[1][2] = true
    //alrMap := make(map[int]bool)
    //alrMap[1][2] = true
    //fmt.Println(alrMap)
    a := [][]byte{
        {'a', 'b', 'b'},
        {'b', 'z', 'b'},
        {'b', 'b', 'a'},
    }
    //a := [][]byte{
    //    {'b', 'b'},
    //    {'b', 'b'},
    //}
    fmt.Println(containsCycle(a))

}

func containsCycle(grid [][]byte) bool {
    if len(grid) == 0 {
        return false
    }
    alrGrid := [500][500]bool{} //已经走过的map
    for m, gri := range grid {
        for n, gr := range gri {
            if alrGrid[m][n] {
                continue
            }
            if getNext(gr, m, n, 0, grid, &alrGrid) {
                return true
            }
        }
    }

    return false
}

//pre 0无 1上 2下 3左 4右
func getNext(preGr byte, m, n int, pre uint32, grid [][]byte, alrGrid *[500][500]bool) bool {
    //判定超界
    if m < 0 || n < 0 || m >= len(grid) || n >= len(grid[0]) || preGr != grid[m][n] {
        return false
    }
    if alrGrid[m][n] {
        return true
    }
    alrGrid[m][n] = true
    switch pre {
    case 0:
        //上
        if getNext(preGr, m-1, n, 1, grid, alrGrid) {
            return true
        }
        //下
        if getNext(preGr, m+1, n, 2, grid, alrGrid) {
            return true
        }
        //左
        if getNext(preGr, m, n-1, 3, grid, alrGrid) {
            return true
        }
        //右
        if getNext(preGr, m, n+1, 4, grid, alrGrid) {
            return true
        }
    case 1:
        //上
        if getNext(preGr, m-1, n, 1, grid, alrGrid) {
            return true
        }
        //左
        if getNext(preGr, m, n-1, 3, grid, alrGrid) {
            return true
        }
        //右
        if getNext(preGr, m, n+1, 4, grid, alrGrid) {
            return true
        }
    case 2:
        //下
        if getNext(preGr, m+1, n, 2, grid, alrGrid) {
            return true
        }
        //左
        if getNext(preGr, m, n-1, 3, grid, alrGrid) {
            return true
        }
        //右
        if getNext(preGr, m, n+1, 4, grid, alrGrid) {
            return true
        }
    case 3:
        //上
        if getNext(preGr, m-1, n, 1, grid, alrGrid) {
            return true
        }
        //下
        if getNext(preGr, m+1, n, 2, grid, alrGrid) {
            return true
        }
        //左
        if getNext(preGr, m, n-1, 3, grid, alrGrid) {
            return true
        }
    case 4:
        //上
        if getNext(preGr, m-1, n, 1, grid, alrGrid) {
            return true
        }
        //下
        if getNext(preGr, m+1, n, 2, grid, alrGrid) {
            return true
        }
        //右
        if getNext(preGr, m, n+1, 4, grid, alrGrid) {
            return true
        }
    }

    return false
}
