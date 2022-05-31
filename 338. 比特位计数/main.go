package main

import "fmt"

//给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
//
// 
//
//示例 1：
//
//输入：n = 2
//输出：[0,1,1]
//解释：
//0 --> 0
//1 --> 1
//2 --> 10
//示例 2：
//
//输入：n = 5
//输出：[0,1,1,2,1,2]
//解释：
//0 --> 0
//1 --> 1
//2 --> 10
//3 --> 11
//4 --> 100
//5 --> 101
// 
//
//提示：
//
//0 <= n <= 105
// 
//
//进阶：
//
//很容易就能实现时间复杂度为 O(n log n) 的解决方案，你可以在线性时间复杂度 O(n) 内用一趟扫描解决此问题吗？
//你能不使用任何内置函数解决此问题吗？（如，C++ 中的 __builtin_popcount ）
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/counting-bits
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(countBits(5))
}









func countBits(n int) []int {
	result := make([]int,n+1)

	if n ==0 {
		return []int{0}
	}
	for i:=1;i<=n;i++ {
		if i&1 == 1 {
			result[i] = result[i>>1]+1
		}else {
			result[i] = result[i>>1]
		}
	}

	return result
}
