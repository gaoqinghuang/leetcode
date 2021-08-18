package main

import (
    "fmt"
    "sort"
)

//给你一个字符串数组 words。
//
//一步操作中，你可以交换字符串 words[i] 的任意两个偶数下标对应的字符或任意两个奇数下标对应的字符。
//
//对两个字符串 words[i] 和 words[j] 而言，如果经过任意次数的操作，words[i] == words[j] ，那么这两个字符串是 特殊等价 的。
//
//例如，words[i] = "zzxy" 和 words[j] = "xyzz" 是一对 特殊等价 字符串，因为可以按 "zzxy" -> "xzzy" -> "xyzz" 的操作路径使 words[i] == words[j] 。
//现在规定，words 的 一组特殊等价字符串 就是 words 的一个同时满足下述条件的非空子集：
//
//该组中的每一对字符串都是 特殊等价 的
//该组字符串已经涵盖了该类别中的所有特殊等价字符串，容量达到理论上的最大值（也就是说，如果一个字符串不在该组中，那么这个字符串就 不会 与该组内任何字符串特殊等价）
//返回 words 中 特殊等价字符串组 的数量。
//
// 
//
//示例 1：
//
//输入：words = ["abcd","cdab","cbad","xyzz","zzxy","zzyx"]
//输出：3
//解释：
//其中一组为 ["abcd", "cdab", "cbad"]，因为它们是成对的特殊等价字符串，且没有其他字符串与这些字符串特殊等价。
//另外两组分别是 ["xyzz", "zzxy"] 和 ["zzyx"]。特别需要注意的是，"zzxy" 不与 "zzyx" 特殊等价。
//示例 2：
//
//输入：words = ["abc","acb","bac","bca","cab","cba"]
//输出：3
//解释：3 组 ["abc","cba"]，["acb","bca"]，["bac","cab"]
// 
//
//提示：
//
//1 <= words.length <= 1000
//1 <= words[i].length <= 20
//所有 words[i] 都只由小写字母组成。
//所有 words[i] 都具有相同的长度。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/groups-of-special-equivalent-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
    words := []string{"abc","acb","bac","bca","cab","cba"}
    fmt.Println(numSpecialEquivGroups(words))
}

func numSpecialEquivGroups(words []string) int {
    specialEquivGroups := make(map[string]bool)
    for _, word := range words {
        specialEquivGroups[toSpecialEquiv(word)] = true
    }

    return len(specialEquivGroups)
}

func toSpecialEquiv(word string) string {
    //按固定格式转换成固定string
    var a, b []rune
    for k, w := range word {
        //奇数
        if k%2 == 1 {
            a = append(a, w)
        } else {
            b = append(b, w)
        }
    }

    sort.Slice(a, func(i, j int) bool {
        return a[i] < a[j]
    })
    sort.Slice(b, func(i, j int) bool {
        return b[i] < b[j]
    })
    return string(append(a, b...))
}
