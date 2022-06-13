package main

import "fmt"

//给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
//
//我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
//
//所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。
//
//请你返回「表现良好时间段」的最大长度。
//
// 
//
//示例 1：
//
//输入：hours = [9,9,6,0,6,6,9]
//输出：3
//解释：最长的表现良好时间段是 [9,9,6]。
//示例 2：
//
//输入：hours = [6,6,6]
//输出：0
// 
//
//提示：
//
//1 <= hours.length <= 104
//0 <= hours[i] <= 16
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/longest-well-performing-interval
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(longestWPI([]int{6,9,9}))
	fmt.Println(longestWPI([]int{6,6,9}))
	fmt.Println(longestWPI([]int{9,9,6,0,6,6,9}))
}



func longestWPI(hours []int) int {
	//整理成前缀和
	pres := make([]int,1,len(hours)+1)
	pres[0] = 0
	for _,hour := range hours {
		score := -1
		if hour > 8 {
			score = 1
		}
		pres = append(pres,pres[len(pres)-1]+score)
	}
	//单调栈
	stock := make([]int,0)
	for index,pre := range pres {
		if len(stock) == 0 {
			stock = append(stock,index)
		}else {
			last := pres[stock[len(stock)-1]]
			if pre < last {
				stock = append(stock,index)
			}
		}
	}

	result := 0
	for i:=len(pres)-1;i>=0;i-- {
		for  {
			if len(stock) == 0 {
				break
			}

			if pres[i] > pres[stock[len(stock)-1]] {
				result = max(result,i-stock[len(stock)-1])
				stock = stock[0:len(stock)-1]
			}else {
				break
			}
		}


	}
	return result
}




// o(n^2),想下能找到一个o(n)的吗
func longestWPI2(hours []int) int {
	result := 0
	score := 0
	for i:=0;i<len(hours);i++ {
		score = 0
		for j:=i;j>=0;j-- {
			if hours[j] >8 {
				score++
			}else {
				score--
			}
			if score > 0 {
				result = max(result,i-j+1)
			}

		}
	}
	return result
}

func max(a,b int) int {
	if a<b {
		return b
	}

	return a
}