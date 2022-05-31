package main

import "fmt"

//给你一个整数 n ，统计并返回各位数字都不同的数字 x 的个数，其中 0 <= x < 10n 。
//
//
//示例 1：
//
//输入：n = 2
//输出：91
//解释：答案应为除去 11、22、33、44、55、66、77、88、99 外，在 0 ≤ x < 100 范围内的所有数字。
//示例 2：
//
//输入：n = 0
//输出：1
//
//
//提示：
//
//0 <= n <= 8
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/count-numbers-with-unique-digits
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(countNumbersWithUniqueDigits(3))
}





//排列组合
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 10
	res := 9
	for i:=2;i<=n;i++ {
		if 11-n <=0 {
			res = 0
		}else {
			res *= 11-i
		}
		dp[i] = dp[i-1]+res
	}

	return dp[n]
}
