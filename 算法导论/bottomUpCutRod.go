package main

import "fmt"

var maxPrice map[int]int

var price map[int]int

func main() {
	//钢条切割，动态规划
	price = map[int]int{
		1:  1,
		2:  5,
		3:  8,
		4:  9,
		5:  10,
		6:  17,
		7:  17,
		8:  20,
		9:  24,
		10: 30,
	}

	maxPrice = map[int]int{
		1: 1,
		2: 5,
	}

	//输出前10的最优解  //到底11的时候又不同了，因为没有0本身
	n := 12
	for i := 3; i <= n; i++ {
		//切割分案
		countMoney, _ := price[i] //不切割的得到的钱
		for j := 1; j <= i/2; j++ {
			if otherMoney := maxPrice[j] + maxPrice[i-j]; otherMoney > countMoney {
				countMoney = otherMoney
			}
		}
		maxPrice[i] = countMoney
	}
	fmt.Println(maxPrice)
}
