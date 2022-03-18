package main

import "fmt"

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//示例 2：
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//示例 3：
//
//输入：nums = [1]
//输出：[[1]]
//
//
//提示：
//
//1 <= nums.length <= 6
//-10 <= nums[i] <= 10
//nums 中的所有整数 互不相同
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/permutations
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	backtrack(0, nums, res, len(nums))
	return res
}

//回溯法
func backtrack(first int, output []int, res [][]int, n int) {
	if first == n {
		res = append(res, output)
		return
	}
	for i := first; i < n; i++ {
		output[i], output[first] = output[first], output[i]
		backtrack(first+1, output, res, n)
		output[i], output[first] = output[first], output[i]
	}
}

//直接暴力
func permute1(arr []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{arr[0]})
	temResult := make([][]int, 0)
	jvTem := make([]int, 0)
	for i, iv := range arr {
		if i == 0 {
			continue
		}
		for _, jv := range result {
			for l := 0; l <= len(jv); l++ {
				jvTem = make([]int, len(jv))
				copy(jvTem, jv)
				temResult = append(temResult, append(jvTem[0:l], append([]int{iv}, jvTem[l:]...)...))
				jvTem = nil
			}
		}
		result = temResult
		temResult = make([][]int, 0)
	}
	return result
}
