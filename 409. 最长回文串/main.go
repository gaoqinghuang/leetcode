package main

import "fmt"

//给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
//
//在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
//
//注意:
//假设字符串的长度不会超过 1010。
//
//示例 1:
//
//输入:
//"abccccdd"
//
//输出:
//7
//
//解释:
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindrome
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(longestPalindrome("ccc"))
}

func longestPalindrome(s string) int {
	dataCount := make(map[int32]int)
	for _, v := range s {
		dataCount[v]++
	}
	result := 0
	hasOdd := false
	for _, count := range dataCount {
		if count&1 == 0 {
			result += count
		} else {
			hasOdd = true
			result += count - 1
		}
	}
	if hasOdd {
		result++
	}
	return result
}
