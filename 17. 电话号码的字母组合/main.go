package main

import "fmt"

//
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
// 
//
//示例 1：
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//示例 2：
//
//输入：digits = ""
//输出：[]
//示例 3：
//
//输入：digits = "2"
//输出：["a","b","c"]
// 
//
//提示：
//
//0 <= digits.length <= 4
//digits[i] 是范围 ['2', '9'] 的一个数字。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
    fmt.Println(letterCombinations("2"))
}

func letterCombinations(digits string) []string {
    result := make([]string, 0)
    if digits == "" {
        return result
    }

    dict := map[int32][]int32{
        '2': {'a', 'b', 'c'},
        '3': {'d', 'e', 'f'},
        '4': {'g', 'h', 'i'},
        '5': {'j', 'k', 'l'},
        '6': {'m', 'n', 'o'},
        '7': {'p', 'q', 'r', 's'},
        '8': {'t', 'u', 'v'},
        '9': {'w', 'x', 'y', 'z'},
    }

    midResult := make([]string, 0)
    for k, digit := range digits {
        midResult = make([]string, len(result))
        copy(midResult, result)
        result = make([]string, 0, len(midResult)*3)

        if k == 0 {
            for _, i := range dict[digit] {
                result = append(result, string(i))
            }
            continue
        }

        for _, j := range midResult {
            for _, i := range dict[digit] {
                result = append(result, j+string(i))
            }
        }
    }
    return result
}
