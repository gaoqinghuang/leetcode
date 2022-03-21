package main

import (
	"fmt"
	"strings"
)

//对于字符串 S 和 T，只有在 S = T + ... + T（T 自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。
//
//返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。
//
// 
//
//示例 1：
//
//输入：str1 = "ABCABC", str2 = "ABC"
//输出："ABC"
//示例 2：
//
//输入：str1 = "ABABAB", str2 = "ABAB"
//输出："AB"
//示例 3：
//
//输入：str1 = "LEET", str2 = "CODE"
//输出：""
// 
//
//提示：
//
//1 <= str1.length <= 1000
//1 <= str2.length <= 1000
//str1[i] 和 str2[i] 为大写英文字母
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/greatest-common-divisor-of-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。









func main() {
	fmt.Println(gcdOfStrings("ABCABCABC","ABC"))
}











//找到两则的公约数，然后判定是否都是这样的数
func gcdOfStrings(str1 string, str2 string) string {
	smallLen := min(len(str1),len(str2))
	for i := smallLen;i>0;i-- {
		if len(str1) % i == 0 && len(str2) % i == 0 {//公约数
			substr := str1[0:i]
			if strings.Count(str1,substr) == len(str1)/i && strings.Count(str2,substr) == len(str2)/i{
				return substr
			}
		}
	}
	return ""
}

func min(a,b int) int {
	if a <b {
		return a
	}
	return b
}
