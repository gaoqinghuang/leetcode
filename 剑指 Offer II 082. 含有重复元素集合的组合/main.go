package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//给定一个可能有重复数字的整数数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。 
//
// 
//
//示例 1:
//
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//示例 2:
//
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//]
// 
//
//提示:
//
//1 <= candidates.length <= 100
//1 <= candidates[i] <= 50
//1 <= target <= 30
// 
//
//注意：本题与主站 40 题相同： https://leetcode-cn.com/problems/combination-sum-ii/
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/4sjJUc
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。










func main() {
	//a := make([]int,0)
	//a = append(a,1)
	//a = append(a,1)
	//a = append(a,1)
	//temArr :=make([]tem,0)
	//temArr = append(temArr,tem{
	//	Arr: a,
	//})
	//s := make([]int,0)
	//s =  append(a,2)
	//temArr = append(temArr,tem{
	//	Arr: s,
	//})
	//s =  append(a,4)
	//temArr = append(temArr,tem{
	//	Arr: s,
	//})
	////b := append(a,2)
	////c := append(a,4)
	//fmt.Println(temArr)
	fmt.Println(combinationSum2([]int{1,2,3,4,5},10))
}

type tem struct {
	Sum int
	K string
}


//2,5,2,1,2
//算出所有可能，比较，如果已经大于了，就没必要再算，所有可能中，要借助上一次的结果,   还得去重·····
func combinationSum2(candidates []int, target int) [][]int{
	sort.Ints(candidates)
	result := make([][]int,0)
	temArr := make([]tem,1)
	temArr[0] = tem{
		0,"",
	}
	repeat := make(map[string]bool)
	repeat[""] = true
	//初始是0
	for _,c := range candidates {
		for _, v :=range temArr {
			newKey := v.K+","+strconv.Itoa(c)
			if repeat[newKey]{
				continue
			}
			if v.Sum+c > target {
				continue
			}

			repeat[newKey]=true
			if v.Sum+c == target {
				//直接string转[]int
				result = append(result,stringToIntSlice(newKey))
				continue
			}
			temArr = append(temArr,tem{
				Sum: v.Sum+c,
				K: newKey,
			})

		}
	}
	
	return result
}

func stringToIntSlice(s string) []int {
	s = strings.Trim(s,",")
	sS := strings.Split(s,",")
	result := make([]int,len(sS))
	for k,v := range sS  {
		result[k],_ = strconv.Atoi(v)
	}
	return result
}

