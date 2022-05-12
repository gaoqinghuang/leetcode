package main

import "fmt"

//一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
//
// 
//
//示例 1：
//
//输入：nums = [4,1,4,6]
//输出：[1,6] 或 [6,1]
//示例 2：
//
//输入：nums = [1,2,10,4,1,4,3,3]
//输出：[2,10] 或 [10,2]
// 
//
//限制：
//
//2 <= nums.length <= 10000
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(singleNumbers([]int{4,1,4,6}))
}

//异或能找出一个不同的数字，关键点要怎么找出两个不同的数字，得分成2组
func singleNumbers(nums []int) []int {
	//全部异或
	ret :=0
	for _,num := range nums {
		ret ^= num
	}
	//找出第几位是1
	re := 1
	for ret&1 == 0 {
		re <<=  1
		ret >>= 1
	}
	a,b := 0,0
	//找出位数不同的，分成两种
	for _,num := range nums {
		if num&re  == 0{
			a ^=num
		}else {
			b ^=num
		}
	}
	return []int{a,b}
}
