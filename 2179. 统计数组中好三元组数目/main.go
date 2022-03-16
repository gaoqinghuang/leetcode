package main

import "fmt"

//给你两个下标从 0 开始且长度为 n 的整数数组 nums1 和 nums2 ，两者都是 [0, 1, ..., n - 1] 的 排列 。
//
//好三元组 指的是 3 个 互不相同 的值，且它们在数组 nums1 和 nums2 中出现顺序保持一致。换句话说，如果我们将 pos1v 记为值 v 在 nums1 中出现的位置，pos2v 为值 v 在 nums2 中的位置，那么一个好三元组定义为 0 <= x, y, z <= n - 1 ，且 pos1x < pos1y < pos1z 和 pos2x < pos2y < pos2z 都成立的 (x, y, z) 。
//
//请你返回好三元组的 总数目 。
//
// 
//
//示例 1：
//
//输入：nums1 = [2,0,1,3], nums2 = [0,1,2,3]
//输出：1
//解释：
//总共有 4 个三元组 (x,y,z) 满足 pos1x < pos1y < pos1z ，分别是 (2,0,1) ，(2,0,3) ，(2,1,3) 和 (0,1,3) 。
//这些三元组中，只有 (0,1,3) 满足 pos2x < pos2y < pos2z 。所以只有 1 个好三元组。
//示例 2：
//
//输入：nums1 = [4,0,1,3,2], nums2 = [4,1,0,2,3]
//输出：4
//解释：总共有 4 个好三元组 (4,0,3) ，(4,0,2) ，(4,1,3) 和 (4,1,2) 。
// 
//
//提示：
//
//n == nums1.length == nums2.length
//3 <= n <= 105
//0 <= nums1[i], nums2[i] <= n - 1
//nums1 和 nums2 是 [0, 1, ..., n - 1] 的排列。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/count-good-triplets-in-an-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(goodTriplets([]int{4,0,1,3,2},[]int{4,1,0,2,3}))
}







//先映射，然后找左边小，右边大的
func goodTriplets(nums1 []int, nums2 []int) int64 {
	//映射
	newNums := make([]int,len(nums1))
	nums2New := make([]int,len(nums2))
	for k,v := range nums2 {
		nums2New[v] = k
	}
	for k,v := range nums1 {
		newNums[k] = nums2New[v]
	}

	//寻找左小右大的数
	result := int64(0)
	for i:=1;i<len(newNums)-1;i++ {
		leftLowCount := 0
		for j:=0;j<i;j++ {
			if newNums[j]<newNums[i]{
				leftLowCount++
			}
		}
		result += int64(leftLowCount)*int64(len(newNums)-newNums[i]-1-i+leftLowCount)
	}

	return result
}


func goodTriplets1(nums1, nums2 []int) (ans int64) {
	n := len(nums1)
	p := make([]int, n)
	for i, v := range nums1 {
		p[v] = i
	}
	tree := make([]int, n+1)
	for i := 1; i < n-1; i++ {
		for j := p[nums2[i-1]] + 1; j <= n; j += j & -j { // 将 p[nums2[i-1]]+1 加入树状数组
			tree[j]++
		}
		y, less := p[nums2[i]], 0
		for j := y; j > 0; j &= j - 1 { // 计算 less
			less += tree[j]
		}
		ans += int64(less) * int64(n-1-y-(i-less))
	}
	return
}
