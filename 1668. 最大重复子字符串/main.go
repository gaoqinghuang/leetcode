package main

import (
	"fmt"
	"strings"
)

//给你一个字符串 sequence ，如果字符串 word 连续重复 k 次形成的字符串是 sequence 的一个子字符串，那么单词 word 的 重复值为 k 。单词 word 的 最大重复值 是单词 word 在 sequence 中最大的重复值。如果 word 不是 sequence 的子串，那么重复值 k 为 0 。
//
//给你一个字符串 sequence 和 word ，请你返回 最大重复值 k 。
//
// 
//
//示例 1：
//
//输入：sequence = "ababc", word = "ab"
//输出：2
//解释："abab" 是 "ababc" 的子字符串。
//示例 2：
//
//输入：sequence = "ababc", word = "ba"
//输出：1
//解释："ba" 是 "ababc" 的子字符串，但 "baba" 不是 "ababc" 的子字符串。
//示例 3：
//
//输入：sequence = "ababc", word = "ac"
//输出：0
//解释："ac" 不是 "ababc" 的子字符串。
// 
//
//提示：
//
//1 <= sequence.length <= 100
//1 <= word.length <= 100
//sequence 和 word 都只包含小写英文字母。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/maximum-repeating-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//fmt.Println(strings.Index("ab","a"))
	fmt.Println(maxRepeating("aaabaaaabaaabaaaabaaaabaaaabaaaaba","aaaba"))
}


func maxRepeating(sequence string, word string) int {
	return max(maxRepeatingZ(sequence,word),maxRepeatingZ(reserveString(sequence),reserveString(word)))
}

//maxRepeatingZ 正序
func maxRepeatingZ(sequence string, word string) int {
	count := 0
	n:=0
	for  {
		i := strings.Index(sequence,word)
		if i == -1 {
			return count
		}

		if n == 0 || i ==0 {
			n++
		}else {
			//不是连续的
			n = 1
		}
		count = max(n,count)
		sequence = sequence[i+len(word):]

	}
}

func max(a,b int)int  {
	if a<b {
		return b
	}

	return a
}

func reserveString(str string) string {
	reserveStr := []rune(str)
	length := len(reserveStr)

	for i := 0; i < length/2; i++ {
		reserveStr[i], reserveStr[length-i-1] = reserveStr[length-i-1], reserveStr[i]
	}

	return string(reserveStr)
}
