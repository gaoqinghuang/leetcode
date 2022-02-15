package main

import (
	"fmt"
	"math/rand"
	"time"
)

//给你一个 下标从 0 开始 的正整数数组 w ，其中 w[i] 代表第 i 个下标的权重。
//
//请你实现一个函数 pickIndex ，它可以 随机地 从范围 [0, w.length - 1] 内（含 0 和 w.length - 1）选出并返回一个下标。选取下标 i 的 概率 为 w[i] / sum(w) 。
//
//例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3) = 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3) = 0.75（即，75%）。
//
//
//示例 1：
//
//输入：
//["Solution","pickIndex"]
//[[[1]],[]]
//输出：
//[null,0]
//解释：
//Solution solution = new Solution([1]);
//solution.pickIndex(); // 返回 0，因为数组中只有一个元素，所以唯一的选择是返回下标 0。
//示例 2：
//
//输入：
//["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
//[[[1,3]],[],[],[],[],[]]
//输出：
//[null,1,1,1,1,0]
//解释：
//Solution solution = new Solution([1, 3]);
//solution.pickIndex(); // 返回 1，返回下标 1，返回该下标概率为 3/4 。
//solution.pickIndex(); // 返回 1
//solution.pickIndex(); // 返回 1
//solution.pickIndex(); // 返回 1
//solution.pickIndex(); // 返回 0，返回下标 0，返回该下标概率为 1/4 。
//
//由于这是一个随机问题，允许多个答案，因此下列输出都可以被认为是正确的:
//[null,1,1,1,1,0]
//[null,1,1,1,1,1]
//[null,1,1,1,0,0]
//[null,1,1,1,0,1]
//[null,1,0,1,0,0]
//......
//诸若此类。
//
//
//提示：
//
//1 <= w.length <= 104
//1 <= w[i] <= 105
//pickIndex 将被调用不超过 104 次
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/random-pick-with-weight
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//for i := 0;i<20;i++ {
	//	fmt.Println(r.Intn(3))
	//
	//}

	solution := Constructor([]int{2})
	fmt.Println(solution.PickIndex())
	fmt.Println(solution.PickIndex())
	fmt.Println(solution.PickIndex())
	fmt.Println(solution.PickIndex())
}

type Solution struct {
	sum    int
	preSum []int
}

func Constructor(w []int) Solution {
	sum := 0
	preSum := make([]int, 0, len(w))
	for _, v := range w {
		sum += v

		preSum = append(preSum, sum)
	}
	return Solution{
		sum:    sum,
		preSum: preSum,
	}
}

func (this *Solution) PickIndex() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randInt := r.Intn(this.sum)
	for k, v := range this.preSum {
		if v >= randInt {
			return k
		}
	}
	return 100
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
