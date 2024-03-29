package main

import (
	"fmt"
	"sort"
)

//给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。
//
//
//
//示例 1：
//
//输入：nums = [5,2,6,1]
//输出：[2,1,1,0]
//解释：
//5 的右侧有 2 个更小的元素 (2 和 1)
//2 的右侧仅有 1 个更小的元素 (1)
//6 的右侧有 1 个更小的元素 (1)
//1 的右侧有 0 个更小的元素
//示例 2：
//
//输入：nums = [-1]
//输出：[0]
//示例 3：
//
//输入：nums = [-1,-1]
//输出：[0,0]
//
//
//提示：
//
//1 <= nums.length <= 105
//-104 <= nums[i] <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
}

func countSmaller(nums []int) []int {
	n := len(nums)
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)

	for i := 0; i < n; i++ {
		nums[i] = sort.SearchInts(tmp, nums[i]) + 1
	}

	bit := BIT{
		n:    n,
		tree: make([]int, n+1),
	}

	ans := make([]int, len(nums))
	for i := n - 1; i >= 0; i-- {
		ans[i] = bit.query(nums[i] - 1)
		bit.update(nums[i])
	}
	return ans
}

// BIT 数状数组
type BIT struct {
	n    int
	tree []int
}

func (b BIT) lowbit(x int) int { return x & (-x) }

func (b BIT) query(x int) int {
	ret := 0
	for x > 0 {
		ret += b.tree[x]
		x -= b.lowbit(x)
	}
	return ret
}

func (b BIT) update(x int) {
	for x <= b.n {
		b.tree[x]++
		x += b.lowbit(x)
	}
}
