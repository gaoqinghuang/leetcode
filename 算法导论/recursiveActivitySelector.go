package main

import (
	"fmt"
)

func main() {
	//活动选择问题 贪心算法
	s := []int{0, 1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12}
	f := []int{0, 4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}
	fmt.Println(recursiveActivitySelector(s, f, 0, len(s)-1))
}
func recursiveActivitySelector(s, f []int, k, n int) []int {
	m := k + 1
	for m <= n && s[m] < f[k] {
		m++
	}
	if m <= n {
		return append([]int{m}, recursiveActivitySelector(s, f, m, n)...)
	} else {
		return []int{}
	}
}
