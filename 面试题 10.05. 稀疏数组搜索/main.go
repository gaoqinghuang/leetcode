package main

import (
	"fmt"
	"strings"
)

//稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。
//
//示例1:
//
//输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
//输出：-1
//说明: 不存在返回-1。
//示例2:
//
//输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
//输出：4
//提示:
//
//words的长度在[1, 1000000]之间
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sparse-array-search-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(findString([]string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}, "ta"))
}

//排好序的，那就二分查找
func findString(words []string, s string) int {
	left, mid, right := 0, 0, len(words)-1
	for left <= right {
		//还有可能是空字符串
		mid = (left + right) / 2
		for left < mid && words[mid] == "" {
			mid--
		}

		switch strings.Compare(s, words[mid]) {
		case 0:
			return mid
		case -1:
			right = mid - 1
		case 1:
			left = mid + 1
		}
	}
	return -1
}

//直接循环
func findString1(words []string, s string) int {
	for k, word := range words {
		if word == s {
			return k
		}
	}
	return -1
}
