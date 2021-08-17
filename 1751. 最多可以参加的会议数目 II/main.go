package main

import (
    "fmt"
    "sort"
)

//给你一个 events 数组，其中 events[i] = [startDayi, endDayi, valuei] ，表示第 i 个会议在 startDayi 天开始，第 endDayi 天结束，如果你参加这个会议，你能得到价值 valuei 。同时给你一个整数 k 表示你能参加的最多会议数目。
//
//你同一时间只能参加一个会议。如果你选择参加某个会议，那么你必须 完整 地参加完这个会议。会议结束日期是包含在会议内的，也就是说你不能同时参加一个开始日期与另一个结束日期相同的两个会议。
//
//请你返回能得到的会议价值 最大和 。
//
//
//
//示例 1：
//
//
//
//输入：events = [[1,2,4],[3,4,3],[2,3,1]], k = 2
//输出：7
//解释：选择绿色的活动会议 0 和 1，得到总价值和为 4 + 3 = 7 。
//示例 2：
//
//
//
//输入：events = [[1,2,4],[3,4,3],[2,3,10]], k = 2
//输出：10
//解释：参加会议 2 ，得到价值和为 10 。
//你没法再参加别的会议了，因为跟会议 2 有重叠。你 不 需要参加满 k 个会议。
//示例 3：
//
//
//
//输入：events = [[1,1,1],[2,2,2],[3,3,3],[4,4,4]], k = 3
//输出：9
//解释：尽管会议互不重叠，你只能参加 3 个会议，所以选择价值最大的 3 个会议。
//
//
//提示：
//
//1 <= k <= events.length
//1 <= k * events.length <= 10^6
//1 <= startDayi <= endDayi <= 10^9
//1 <= valuei <= 10^6

func main() {
    //for i := 0; i < 0; i++ {
    //    fmt.Println(1)
    //}
    //
    //return
    //a := make([]int,0,0)
    //
    //a = append(a,1)
    //a = append(a,2)
    //a = append(a,3)
    //
    //b := append(a,4)
    //c := append(a,5)
    //c[3] =6
    //fmt.Println(a,b,c,b[3])
    //return
    //events = [[1,2,4],[3,4,3],[2,3,1]], k = 2
    //[[3,68,97],[12,46,13],[21,24,75],[64,85,74],[10,98,15],[23,84,62],[87,96,29],[80,85,39],[52,89,77],[31,63,91],[29,40,48],[30,96,42],[69,81,68],[52,58,65],[41,52,37]]
    fmt.Println(maxValue([][]int{{3, 68, 97}, {12, 46, 13}, {21, 24, 75}, {64, 85, 74}, {10, 98, 15}, {23, 84, 62}, {87, 96, 29}, {80, 85, 39}, {52, 89, 77}, {31, 63, 91}, {29, 40, 48}, {30, 96, 42}, {69, 81, 68}, {52, 58, 65}, {41, 52, 37}}, 10))
    //fmt.Println(maxValue([][]int{{3, 68, 97}, {21, 24, 75}, {29, 40, 48}}, 2))
    //fmt.Println(maxValue([][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}, 2))
    //fmt.Println(maxValue([][]int{{1,2,4},{3,4,3},{2,3,10}}, 2))
    //fmt.Println(maxValue([][]int{{1,1,1},{2,2,2},{3,3,3},{4,4,4}}, 3))
}

func maxValue(events [][]int, k int) int {
    // 借助前面算好的计算后面的东西
    //按end排序
    sort.Slice(events, func(i, j int) bool {
        return events[i][1] < events[j][1]
    })

    //dp[i][j]  前i个会议，在限选j个的情况下的价值，背包问题
    dp := make([][]int, len(events)+1)
    //初始化
    for i := 0; i <= len(events); i++ {
        dp[i] = make([]int, k+1)
    }

    //空间换时间，记录上一个不冲突的
    pre := make(map[int]int, 0)
    for i := 1; i <= len(events); i++ {
        for j := i - 1; j >= 1; j-- {
            if events[i-1][0] > events[j-1][1] {
                pre[i] = j
                break
            }
        }
    }

    for key, event := range events {
        i := key + 1
        for j := 1; j <= k; j++ {
            dp[i][j] = maxInt(dp[i-1][j], dp[pre[i]][j-1]+event[2])
        }
    }

    return dp[len(events)][k]
}

func maxInt(a, b int) int {
    if b > a {
        return b
    }
    return a
}
