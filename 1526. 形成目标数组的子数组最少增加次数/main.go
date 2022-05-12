package main

import (
	"fmt"
)

//给你一个整数数组 target 和一个数组 initial ，initial 数组与 target  数组有同样的维度，且一开始全部为 0 。
//
//请你返回从 initial 得到  target 的最少操作次数，每次操作需遵循以下规则：
//
//在 initial 中选择 任意 子数组，并将子数组中每个元素增加 1 。
//答案保证在 32 位有符号整数以内。
//
//
//
//示例 1：
//
//输入：target = [1,2,3,2,1]
//输出：3
//解释：我们需要至少 3 次操作从 intial 数组得到 target 数组。
//[0,0,0,0,0] 将下标为 0 到 4 的元素（包含二者）加 1 。
//[1,1,1,1,1] 将下标为 1 到 3 的元素（包含二者）加 1 。
//[1,2,2,2,1] 将下表为 2 的元素增加 1 。
//[1,2,3,2,1] 得到了目标数组。
//示例 2：
//
//输入：target = [3,1,1,2]
//输出：4
//解释：(initial)[0,0,0,0] -> [1,1,1,1] -> [1,1,1,2] -> [2,1,1,2] -> [3,1,1,2] (target) 。
//示例 3：
//
//输入：target = [3,1,5,4,2]
//输出：7
//解释：(initial)[0,0,0,0,0] -> [1,1,1,1,1] -> [2,1,1,1,1] -> [3,1,1,1,1]
//-> [3,1,2,2,2] -> [3,1,3,3,2] -> [3,1,4,4,2] -> [3,1,5,4,2] (target)。
//示例 4：
//
//输入：target = [1,1,1,1]
//输出：1
//
//
//提示：
//
//1 <= target.length <= 10^5
//1 <= target[i] <= 10^5
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(minNumberOperations([]int{1, 2, 3, 2, 1}))
	fmt.Println(minNumberOperations([]int{3, 1, 1, 2}))
	fmt.Println(minNumberOperations([]int{3, 1, 5, 4, 2}))
}

//minNumberOperations 通过下面发现，只要算波峰的
func minNumberOperations(target []int) int {
	ans := target[0]
	for i := 1; i < len(target); i++ {
		ans += max(target[i]-target[i-1], 0)
	}
	return ans

}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//
//func minNumberOperations(target []int) int {
//	//每次找最左的数
//	initial := make([]int, len(target))
//	i := 0
//	count:=0
//	for i < len(target) {
//		//如果是递增的,找最后一位
//		if initial[i] == target[i] || (i != len(target)-1 && target[i]<=target[i+1]){
//			i++
//			continue
//		}
//		count +=target[i]-initial[i]
//		subNum := math.MaxInt
//		for j:=i;j<len(target);j++ {
//          一次性最少能加subNum
//			subNum = min(target[j]-initial[j],subNum)
//			if subNum == 0 {
//				break
//			}
//			initial[j] += subNum
//		}
//		i++
//	}
//
//	return count
//}
//
//func min(a,b int) int {
//	if a<b {
//		return a
//	}
//	return b
//}
