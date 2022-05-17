package main

import (
	"fmt"
	"strconv"
)

//给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。
//
//表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
//
//示例 1:
//
//输入: "3+2*2"
//输出: 7
//示例 2:
//
//输入: " 3/2 "
//输出: 1
//示例 3:
//
//输入: " 3+5 / 2 "
//输出: 5
//说明：
//
//你可以假设所给定的表达式都是有效的。
//请不要使用内置的库函数 eval。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/calculator-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//fmt.Println(stringToNum("-11"))
	fmt.Println(calculate("3+2*2"))
}

type stack struct {
	arr []int
}

func (s *stack) RAdd(val int) {
	s.arr = append(s.arr, val)
}

func (s *stack) RDel() int {
	last := s.arr[len(s.arr)-1]
	s.arr = s.arr[0 : len(s.arr)-1]
	return last
}
func (s *stack) Sum() int {
	result := 0
	for _, v := range s.arr {
		result += v
	}
	return result
}

//直接循环
func calculate(s string) int {
	st := &stack{
		arr: make([]int, 0),
	}
	lastNumStr := ""
	lastC := 0 //0无 1* 2/
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			continue
		case '+':
			lastCS(lastNumStr, lastC, st)
			lastNumStr = ""
			lastC = 0
		case '-':
			lastCS(lastNumStr, lastC, st)
			lastNumStr = "-"
			lastC = 0
		case '*':
			lastCS(lastNumStr, lastC, st)
			//这里要往后找
			lastC = 1
			lastNumStr = ""
		case '/':
			lastCS(lastNumStr, lastC, st)
			lastNumStr = ""
			lastC = 2
		default:
			lastNumStr += string(s[i])
			continue
		}
	}
	if lastNumStr != "" {
		lastCS(lastNumStr, lastC, st)
	}
	return st.Sum()
}

func lastCS(lastNumStr string, lastC int, st *stack) {
	num := stringToNum(lastNumStr)
	if lastC != 0 {
		if lastC == 1 {
			num = num * st.RDel()
		} else {
			num = st.RDel() / num
		}
	}
	st.RAdd(num)
}

func stringToNum(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
