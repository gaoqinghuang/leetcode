package main

import "fmt"

func main() {
	//最长公共子序列
	a := "ABCBDAB"
	b := "BDCABA"

	record := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		record[i] = make([]int, len(b)+1)
	}
	//求record[]
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				record[i][j] = record[i-1][j-1] + 1
			} else if record[i-1][j] >= record[i][j-1] {
				record[i][j] = record[i-1][j]
			} else {
				record[i][j] = record[i][j-1]
			}
		}
	}

	fmt.Println(record)
}
