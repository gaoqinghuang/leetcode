package main

import (
	"fmt"
	"strconv"
)

//
//给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
//
//示例 1 :
//
//输入: 2736
//输出: 7236
//解释: 交换数字2和数字7。
//示例 2 :
//
//输入: 9973
//输出: 9973
//解释: 不需要交换。
//注意:
//
//给定数字的范围是 [0, 108]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-swap
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//交换字符串两个位置

	fmt.Println(maximumSwap(2736))
}

//记录每个数出现的位置（最后一个出现），然后从左到右找出可以交换的数
func maximumSwap(num int) int {
	numS := strconv.Itoa(num)
	lastN := make(map[int32]int)
	for k, v := range numS {
		lastN[v] = k
	}

	for k, v := range numS {
		for start := '9'; start > v; start-- {
			if lastK, ok := lastN[start]; ok && lastK > k {
				//找到了，交换lastK和k
				return swap([]byte(numS), k, lastK)
			}
		}
	}

	//交换a,b两位置的数，即可
	return num
}

func swap(num []byte, i, j int) int {
	num[i], num[j] = num[j], num[i]
	result, _ := strconv.Atoi(string(num))
	return result
}
