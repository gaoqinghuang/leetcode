package main

import (
	"fmt"
	"math"
	"sort"
)

//给你一个 正 整数数组 beans ，其中每个整数表示一个袋子里装的魔法豆的数目。
//
//请你从每个袋子中 拿出 一些豆子（也可以 不拿出），使得剩下的 非空 袋子中（即 至少 还有 一颗 魔法豆的袋子）魔法豆的数目 相等 。一旦魔法豆从袋子中取出，你不能将它放到任何其他的袋子中。
//
//请你返回你需要拿出魔法豆的 最少数目。
//
//
//
//示例 1：
//
//输入：beans = [4,1,6,5]
//输出：4
//解释：
//- 我们从有 1 个魔法豆的袋子中拿出 1 颗魔法豆。
//剩下袋子中魔法豆的数目为：[4,0,6,5]
//- 然后我们从有 6 个魔法豆的袋子中拿出 2 个魔法豆。
//剩下袋子中魔法豆的数目为：[4,0,4,5]
//- 然后我们从有 5 个魔法豆的袋子中拿出 1 个魔法豆。
//剩下袋子中魔法豆的数目为：[4,0,4,4]
//总共拿出了 1 + 2 + 1 = 4 个魔法豆，剩下非空袋子中魔法豆的数目相等。
//没有比取出 4 个魔法豆更少的方案。
//示例 2：
//
//输入：beans = [2,10,3,2]
//输出：7
//解释：
//- 我们从有 2 个魔法豆的其中一个袋子中拿出 2 个魔法豆。
//剩下袋子中魔法豆的数目为：[0,10,3,2]
//- 然后我们从另一个有 2 个魔法豆的袋子中拿出 2 个魔法豆。
//剩下袋子中魔法豆的数目为：[0,10,3,0]
//- 然后我们从有 3 个魔法豆的袋子中拿出 3 个魔法豆。
//剩下袋子中魔法豆的数目为：[0,10,0,0]
//总共拿出了 2 + 2 + 3 = 7 个魔法豆，剩下非空袋子中魔法豆的数目相等。
//没有比取出 7 个魔法豆更少的方案。
//
//
//提示：
//
//1 <= beans.length <= 105
//1 <= beans[i] <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/removing-minimum-number-of-magic-beans
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(minimumRemoval([]int{4, 1, 6, 5}))
}

func minimumRemoval(beans []int) int64 {
	//排序
	sort.Ints(beans)
	//前缀和
	preSum := getPreSum(beans)
	//算出每个数最小
	last := -1
	result := int64(math.MaxInt64)

	before := int64(0)
	after := int64(0)
	length := len(beans)
	for k, bean := range beans {
		if last == bean {
			continue
		}
		last = bean

		if k-1 >= 0 {
			before = preSum[k-1]
		}
		//这里可以不用前缀和，直接用totle-bean*xn (xn为大于等于bean的个数)
		after = preSum[len(beans)-1] - preSum[k] - int64(bean*(length-k-1))

		result = min(result, before+after)
	}
	return result
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func getPreSum(nums []int) []int64 {
	result := make([]int64, len(nums))
	last := int64(0)
	for k, num := range nums {
		last += int64(num)
		result[k] = last
	}

	return result
}
