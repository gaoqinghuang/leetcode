package main

import "fmt"

//给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
//
//回文字符串 是正着读和倒过来读一样的字符串。
//
//子字符串 是字符串中的由连续字符组成的一个序列。
//
//具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
// 
//
//示例 1：
//
//输入：s = "abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//示例 2：
//
//输入：s = "aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
// 
//
//提示：
//
//1 <= s.length <= 1000
//s 由小写英文字母组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/palindromic-substrings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(countSubstrings("aaa"))
}

//aaa中心拓展
func countSubstrings(s string) int {
	result := 0
	sBL := []byte(s)
	for i:=0;i<len(sBL);i++ {
		//命中字母
		result++
		left := i-1
		right := i+1
		for true {
			if left <0 || right >= len(sBL) || sBL[left] != sBL[right]{
				break
			}
			result++
			left--
			right++
		}
		//命中字母空隙
		left = i
		right = i+1
		for true {
			if left <0 || right >= len(sBL) || sBL[left] != sBL[right]{
				break
			}
			result++
			left--
			right++
		}
	}

	return result
}