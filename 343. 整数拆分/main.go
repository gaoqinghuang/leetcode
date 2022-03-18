package main

import (
	"fmt"
	"math"
)

//给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
//
//返回 你可以获得的最大乘积 。
//
// 
//
//示例 1:
//
//输入: n = 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1。
//示例 2:
//
//输入: n = 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
// 
//
//提示:
//
//2 <= n <= 58
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/integer-break
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(integerBreak(10))
}





//大于4以后，全部换成2和3最大，2不超过2个
func integerBreak(n int) int {
	if n <= 4 {
		return [5]int{0,1,1,2,4}[n]
	}

	remainder := n%3
	if remainder == 0 {
		return int(math.Pow(3,float64(n/3)))
	}else if remainder == 1 {
		return int(math.Pow(3,float64(n/3-1)))*4
	}else {
		return int(math.Pow(3,float64(n/3)))*2
	}

}