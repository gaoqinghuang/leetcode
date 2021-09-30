package main

import (
	"fmt"
	"strings"
)

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
//
//
//示例：
//
//输入："Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
//
//
//提示：
//
//在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	var result string
	sSlice := strings.Split(s, " ")
	for k, value := range sSlice {
		for i := len(value)-1; i >= 0; i-- {
			result += string(value[i])
		}
		if k == len(sSlice) -1 {
			continue
		}
		result += " "
	}
	return result
}
