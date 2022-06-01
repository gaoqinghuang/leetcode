package main

import (
	"fmt"
	"math"
)

//某乐团的演出场地可视作 num * num 的二维矩阵 grid（左上角坐标为 [0,0])，每个位置站有一位成员。乐团共有 9 种乐器，乐器编号为 1~9，每位成员持有 1 个乐器。
//
//为保证声乐混合效果，成员站位规则为：自 grid 左上角开始顺时针螺旋形向内循环以 1，2，...，9 循环重复排列。例如当 num = 5 时，站位如图所示
//
//
//
//请返回位于场地坐标 [Xpos,Ypos] 的成员所持乐器编号。
//
//示例 1：
//
//输入：num = 3, Xpos = 0, Ypos = 2
//
//输出：3
//
//解释：
//
//
//示例 2：
//
//输入：num = 4, Xpos = 1, Ypos = 2
//
//输出：5
//
//解释：
//
//
//提示：
//
//1 <= num <= 10^9
//0 <= Xpos, Ypos < num
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/SNJvJP
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(orchestraLayout(4, 1, 2))
}

//算出每圈的数量，问题转化为第n圈，怎么算他在第几位
func orchestraLayout(num int, xPos int, yPos int) int {
	//判定第几圈
	indexCirleNum := minSlice([]int{xPos, yPos, num - xPos - 1, num - yPos - 1}) + 1

	index := (num*num-(num-(indexCirleNum-1)*2)*(num-(indexCirleNum-1)*2))%9 + 1

	right := num - indexCirleNum

	left := indexCirleNum - 1
	if xPos == left {
		// 上横
		index += yPos - left
	} else if yPos == right {
		//  右竖
		index += right - left
		index += xPos - left
	} else if xPos == right {
		//  下横
		index += 2 * (right - left)
		index += right - yPos
	} else {
		//  左竖
		index += 3 * (right - left)
		index += right - xPos
	}

	//算出数字
	if index%9 == 0 {
		index = 9
	} else {
		index = index % 9
	}
	return index
}

func minSlice(nums []int) int {
	minNum := math.MaxInt
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}

	return minNum
}
