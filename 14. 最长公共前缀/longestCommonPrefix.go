//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母 a-z 。

package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"abcd", "abc", "ab"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	for i, v := range strs {
		if i == 0 {
			continue
		}
		prefix = twoCommonPrefix(prefix, v)
	}

	return prefix
}

//比较两个字符串的公共前缀
func twoCommonPrefix(a, b string) string {
	prefix := ""
	if len(a) > len(b) {
		a, b = b, a
	}
	for k, v := range a {
		if v == int32(b[k]) {
			prefix += string(v)
		} else {
			break
		}
	}
	return prefix
}
