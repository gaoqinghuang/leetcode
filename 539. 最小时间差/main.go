package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
//
// 
//
//示例 1：
//
//输入：timePoints = ["23:59","00:00"]
//输出：1
//示例 2：
//
//输入：timePoints = ["00:00","23:59","00:00"]
//输出：0
// 
//
//提示：
//
//2 <= timePoints <= 2 * 104
//timePoints[i] 格式为 "HH:MM"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-time-difference
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(findMinDifference([]string{
		"23:59","00:00","00:00",
	}))
}

func findMinDifference(timePoints []string) int {
	//转换成分钟数
	timeMinutePoints := make([]int, len(timePoints))
	for k, timePoint := range timePoints {
		timePointString := strings.Split(timePoint, ":")
		hour, _ := strconv.Atoi(timePointString[0])
		minute, _ := strconv.Atoi(timePointString[1])

		timeMinutePoints[k] = hour*60 + minute
	}
	//排序
	sort.Ints(timeMinutePoints)
	//处理最早最晚时间
	timeMinutePoints = append(timeMinutePoints, 24*60+timeMinutePoints[0])

	//寻找最小
	result := 24 * 60
	for k := 1; k < len(timeMinutePoints); k++ {
		subMinute := timeMinutePoints[k] - timeMinutePoints[k-1]
		if result > subMinute {
			result = subMinute
		}
	}

	return result
}