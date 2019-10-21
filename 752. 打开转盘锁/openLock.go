//你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
//
//锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
//
//列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
//
//字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。
//
//
//
//示例 1:
//
//输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
//输出：6
//解释：
//可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
//注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
//因为当拨动到 "0102" 时这个锁就会被锁定。
//示例 2:
//
//输入: deadends = ["8888"], target = "0009"
//输出：1
//解释：
//把最后一位反向旋转一次即可 "0000" -> "0009"。
//示例 3:
//
//输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
//输出：-1
//解释：
//无法旋转到目标数字且不被锁定。
//示例 4:
//
//输入: deadends = ["0000"], target = "8888"
//输出：-1
//
//
//提示：
//
//死亡列表 deadends 的长度范围为 [1, 500]。
//目标数字 target 不会在 deadends 之中。
//每个 deadends 和 target 中的字符串的数字会在 10,000 个可能的情况 '0000' 到 '9999' 中产生。

package main

import (
	"fmt"
)

func main() {

	fmt.Println(openLock([]string{"0201","0101","0102","1212","2002"}, "0202"))
}

var total map[string]uint8

func openLock(deadends []string, target string) int {
	total = map[string]uint8{}
	count := 0
	for _, v := range deadends {
		total[v] = 1
	}
	if inTotal("0000") {
		return -1
	}
	if target == "0000" {
		if inTotal(target) {
			return -1
		}
		return 0
	}

	//深度搜索
	//一延伸8，8记录在册，如果存在或者有问题就不记录，记录走了多少
	nextList := map[string]uint8{"0000": 1}
	for {
		nextList = getNext(nextList)
		if len(nextList) == 0 {
			return -1
		}
		count++
		if nextList[target] == 1 {
			break
		}
	}
	return count
}

func inTotal(target string) bool {
	return total[target] == 1
}
func getNext(startArr map[string]uint8) map[string]uint8 {
	result := map[string]uint8{}

	for start := range startArr {
		for i := 0; i < 4; i++ {
			a := start[0:i] + addString(start[i:i+1]) + start[i+1:4]
			b := start[0:i] + subString(start[i:i+1]) + start[i+1:4]
			if !inTotal(a) {
				result[a] = 1
				total[a] = 1
			}
			if !inTotal(b) {
				result[b] = 1
				total[b] = 1
			}
		}
	}

	return result
}

func addString(start string) string {
	switch start {
	case "1":
		return "2"
	case "2":
		return "3"
	case "3":
		return "4"
	case "4":
		return "5"
	case "5":
		return "6"
	case "6":
		return "7"
	case "7":
		return "8"
	case "8":
		return "9"
	case "9":
		return "0"
	case "0":
		return "1"
	}
	return ""
}

func subString(start string) string {
	switch start {
	case "1":
		return "0"
	case "2":
		return "1"
	case "3":
		return "2"
	case "4":
		return "3"
	case "5":
		return "4"
	case "6":
		return "5"
	case "7":
		return "6"
	case "8":
		return "7"
	case "9":
		return "8"
	case "0":
		return "9"
	}
	return ""
}
