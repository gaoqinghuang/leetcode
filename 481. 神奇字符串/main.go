package main

import "fmt"

//神奇字符串 s 仅由 '1' 和 '2' 组成，并需要遵守下面的规则：
//
//神奇字符串 s 的神奇之处在于，串联字符串中 '1' 和 '2' 的连续出现次数可以生成该字符串。
//s 的前几个元素是 s = "1221121221221121122……" 。如果将 s 中连续的若干 1 和 2 进行分组，可以得到 "1 22 11 2 1 22 1 22 11 2 11 22 ......" 。每组中 1 或者 2 的出现次数分别是 "1 2 2 1 1 2 1 2 2 1 2 2 ......" 。上面的出现次数正是 s 自身。
//
//给你一个整数 n ，返回在神奇字符串 s 的前 n 个数字中 1 的数目。
//
//
//
//示例 1：
//
//输入：n = 6
//输出：3
//解释：神奇字符串 s 的前 6 个元素是 “122112”，它包含三个 1，因此返回 3 。
//示例 2：
//
//输入：n = 1
//输出：1
//
//
//提示：
//
//1 <= n <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/magical-string
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(magicalString(9))
}

//模拟
func magicalString(n int) int {
	if n == 1 {
		return 1
	}
	s := make([]byte, 0, n+2)
	s = append(s, '1', '2', '2')
	next := 2
	res := 1
	for i := 2; i < n; i++ {
		if s[i] == '1' {
			res++
			if s[next] == '2' {
				s = append(s, '1')
			} else {
				s = append(s, '2')
			}
		} else {
			if s[next] == '2' {
				s = append(s, '1', '1')
			} else {
				s = append(s, '2', '2')
			}
			next++
		}
		next++
	}

	return res
}

