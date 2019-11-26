package main

import (
	"fmt"
	"math"
)

var record [][]uint64
var p = []uint64{10, 100, 5, 50}

//矩阵链乘法
func main() {
	record = make([][]uint64, len(p))
	for i := 0; i < len(p); i++ {
		record[i] = make([]uint64, len(p))
	}
	fmt.Println(matrixChainOrder(1, len(p)-1))
}

func matrixChainOrder(i, j int) uint64 {
	if i == j {
		return 0
	}
	if data := record[i][j]; data > 0 {
		return data
	}
	result := uint64(math.MaxUint64)
	for k := i; k < j; k++ {
		//取得区分值
		tempResult := matrixChainOrder(i, k) + matrixChainOrder(k+1, j) + p[i-1]*p[k]*p[j]
		//比较
		if tempResult < result {
			result = tempResult
		}
	}
	//储存
	record[i][j] = result
	return result
}
