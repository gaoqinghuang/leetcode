package main

import (
	"fmt"
)

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
// 
//
//示例 1：
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
//[1,2,1],
//[2,1,1]]
//示例 2：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
//
//提示：
//
//1 <= nums.length <= 8
//-10 <= nums[i] <= 10
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/permutations-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(permuteUnique([]int{1,1,2}))
}


var res [][]int
var uiq map[string]bool
func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	uiq = make(map[string]bool, 0)
	backtrack(0, nums, len(nums))
	return res
}

//回溯法
func backtrack(first int, output []int,  n int) {
	if first == n {
		u := uq(output)
		if uiq[u] {
			return
		}
		res = append(res, append([]int(nil), output...))
		uiq[u] = true
		return
	}
	for i := first; i < n; i++ {
		output[i], output[first] = output[first], output[i]
		backtrack(first+1, output, n)
		output[i], output[first] = output[first], output[i]
	}
}

func uq(output []int) string {
	result := ""
	for _,v := range output {
		result += fmt.Sprintf("%d",v)
	}
	return result
}


