package main

import "fmt"

//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
//示例 1：
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//示例 2：
//
//输入：s = "cbbd"
//输出："bb"
//示例 3：
//
//输入：s = "a"
//输出："a"
//示例 4：
//
//输入：s = "ac"
//输出："a"
//
//
//提示：
//
//1 <= s.length <= 1000
//s 仅由数字和英文字母（大写和/或小写）组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindromic-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(longestPalindrome("ac"))
}

//暴力解法,总有一个中心吧，从中心往两边扩张,假设这个位置是中心，然后计算
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	result := string(s[0])
	tem := ""
	for i := 1; i < len(s); i++ {
		//中心是空
		tem = getByEmpty(s, i)
		if len(tem) > len(result) {
			result = tem
		}
		//中心有值
		tem = getByOne(s, i)
		if len(tem) > len(result) {
			result = tem
		}
	}

	return result
}

func getByEmpty(s string, i int) string {
	result := ""
	left := i - 1
	right := i
	for {
		if left < 0 || right >= len(s) {
			break
		}
		if s[left] != s[right] {
			break
		}
		result = string(s[left]) + result + string(s[right])
		left--
		right++
	}
	return result
}

func getByOne(s string, i int) string {
	result := string(s[i])
	left := i - 1
	right := i + 1
	for {
		if left < 0 || right >= len(s) {
			break
		}
		if s[left] != s[right] {
			break
		}
		result = string(s[left]) + result + string(s[right])
		left--
		right++
	}
	return result
}
