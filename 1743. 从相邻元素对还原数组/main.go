package main

import "fmt"

//存在一个由 n 个不同元素组成的整数数组 nums ，但你已经记不清具体内容。好在你还记得 nums 中的每一对相邻元素。
//
//给你一个二维整数数组 adjacentPairs ，大小为 n - 1 ，其中每个 adjacentPairs[i] = [ui, vi] 表示元素 ui 和 vi 在 nums 中相邻。
//
//题目数据保证所有由元素 nums[i] 和 nums[i+1] 组成的相邻元素对都存在于 adjacentPairs 中，存在形式可能是 [nums[i], nums[i+1]] ，也可能是 [nums[i+1], nums[i]] 。这些相邻元素对可以 按任意顺序 出现。
//
//返回 原始数组 nums 。如果存在多种解答，返回 其中任意一个 即可。
//
// 
//
//示例 1：
//
//输入：adjacentPairs = [[2,1],[3,4],[3,2]]
//输出：[1,2,3,4]
//解释：数组的所有相邻元素对都在 adjacentPairs 中。
//特别要注意的是，adjacentPairs[i] 只表示两个元素相邻，并不保证其 左-右 顺序。
//示例 2：
//
//输入：adjacentPairs = [[4,-2],[1,4],[-3,1]]
//输出：[-2,4,1,-3]
//解释：数组中可能存在负数。
//另一种解答是 [-3,1,4,-2] ，也会被视作正确答案。
//示例 3：
//
//输入：adjacentPairs = [[100000,-100000]]
//输出：[100000,-100000]
// 
//
//提示：
//
//nums.length == n
//adjacentPairs.length == n - 1
//adjacentPairs[i].length == 2
//2 <= n <= 105
//-105 <= nums[i], ui, vi <= 105
//题目数据保证存在一些以 adjacentPairs 作为元素对的数组 nums
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/restore-the-array-from-adjacent-pairs
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(restoreArray([][]int{{2,1},{3,4},{3,2}}))
}


func restoreArray(adjacentPairs [][]int) []int {
	data := make(map[int][]int)
	result := make([]int,0,len(adjacentPairs)+1)

	for _,adjacentPair := range adjacentPairs {
		if _,ok := data[adjacentPair[0]] ;!ok {
			data[adjacentPair[0]] = make([]int,0,2)
		}
		data[adjacentPair[0]] = append(data[adjacentPair[0]],adjacentPair[1])

		if _,ok := data[adjacentPair[1]] ;!ok {
			data[adjacentPair[1]] = make([]int,0,2)
		}
		data[adjacentPair[1]] = append(data[adjacentPair[1]],adjacentPair[0])
	}

	//找出开头
	for k,d := range data {
		if len(d) == 1 {
			result = append(result,k)
			result = append(result,d[0])
			break
		}
	}
	for  {
		if len(result) == len(adjacentPairs)+1 {
			break
		}
		pData := data[result[len(result)-1]]
		if len(pData) == 1{
			result = append(result,pData[0])
		}else {
			if pData[0] == result[len(result)-2]{
				result = append(result,pData[1])
			}else {
				result = append(result,pData[0])
			}
		}

	}

	return result
}