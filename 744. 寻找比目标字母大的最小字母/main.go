package main

import "fmt"

//给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母 target，请你寻找在这一有序列表里比目标字母大的最小字母。
//
//在比较时，字母是依序循环出现的。举个例子：
//
//如果目标字母 target = 'z' 并且字符列表为 letters = ['a', 'b']，则答案返回 'a'
//
//
//示例：
//
//输入:
//letters = ["c", "f", "j"]
//target = "a"
//输出: "c"
//
//输入:
//letters = ["c", "f", "j"]
//target = "c"
//输出: "f"
//
//输入:
//letters = ["c", "f", "j"]
//target = "d"
//输出: "f"
//
//输入:
//letters = ["c", "f", "j"]
//target = "g"
//输出: "j"
//
//输入:
//letters = ["c", "f", "j"]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'a', 'b'}, 'z')))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	result := letters[0]
	for k := 1; k < len(letters); k++ {
		if letters[k] > target && letters[k-1] <= target {
			result = letters[k]
		}
	}
	return result
}
