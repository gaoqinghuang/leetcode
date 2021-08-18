package main

import (
    "fmt"
    "math"
)

//给定两条线段（表示为起点start = {X1, Y1}和终点end = {X2, Y2}），如果它们有交点，请计算其交点，没有交点则返回空值。
//
//要求浮点型误差不超过10^-6。若有多个交点（线段重叠）则返回 X 值最小的点，X 坐标相同则返回 Y 值最小的点。
//
// 
//
//示例 1：
//
//输入：
//line1 = {0, 0}, {1, 0}
//line2 = {1, 1}, {0, -1}
//输出： {0.5, 0}
//示例 2：
//
//输入：
//line1 = {0, 0}, {3, 3}
//line2 = {1, 1}, {2, 2}
//输出： {1, 1}
//示例 3：
//
//输入：
//line1 = {0, 0}, {1, 1}
//line2 = {1, 0}, {2, 1}
//输出： {}，两条线段没有交点
// 
//
//提示：
//
//坐标绝对值不会超过 2^7
//输入的坐标均是有效的二维坐标
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/intersection-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
    //line1 = {0, 0}, {1, 0}
    //line2 = {1, 1}, {0, -1}
    //输出： {0.5, 0}

    //line1 = {0, 0}, {3, 3}
    //line2 = {1, 1}, {2, 2}

    //line1 = {0, 0}, {1, 1}
    //line2 = {1, 0}, {2, 1}

    //[0,3]
    //[0,6]
    //[0,1]
    //[0,5]

    //fmt.Println(intersection([]int{0, 0}, []int{1, 0}, []int{1, 1}, []int{0, -1}))
    //fmt.Println(intersection([]int{0, 0}, []int{3, 3}, []int{1, 1}, []int{2, 2}))
    //fmt.Println(intersection([]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{2, 1}))
    //fmt.Println(intersection([]int{0, 3}, []int{0, 6}, []int{0, 1}, []int{0, 5}))
    //[1,0]
    //[1,1]
    //[-1,0]
    //[3,2]

    //fmt.Println(intersection([]int{1, 0}, []int{1, 1}, []int{-1, 0}, []int{3, 2}))
    //[0,-1]
    //[0,1]
    //[-1,1]
    //[1,3]
    fmt.Println(intersection([]int{0, -1}, []int{0, 1}, []int{-1, 1}, []int{1, 3}))
}

//强硬列举各种情况
func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
    //交换位置 从左到右
    start1, end1, start2, end2 = change(start1, end1, start2, end2)

    //考虑竖线的情况

    if start1[0] == end1[0] { //a是竖线
        //b 也是竖线
        if start2[0] == end2[0] {
            if start1[0] == start2[0] {
                if start2[1] <= end1[1] {
                    return []float64{
                        float64(start2[0]), float64(start2[1]),
                    }
                }
            }

            return make([]float64, 0)
        }

        //b不是竖线
        if start2[0] <= start1[0] && start1[0] <= end2[0] {
            b := getEquation(start2, end2)
            intersectPoint := []float64{float64(start1[0]), b[1] + float64(start1[0])*b[0]}
            return in(intersectPoint, start1, end1, start2, end2)
        }
    } else if start2[0] == end2[0] { // b是竖线
        if start1[0] <= start2[0] && start2[0] <= end1[0] {
            b := getEquation(start1, end1)
            intersectPoint := []float64{float64(start2[0]), b[1] + float64(start2[0])*b[0]}
            return in(intersectPoint, start1, end1, start2, end2)
        }
    }

    //先判定是否是横线或者竖线  先不判定是否是竖线或者是横线
    a := getEquation(start1, end1)
    b := getEquation(start2, end2)
    //斜率一样
    if a[0] == b[0] {
        //重叠
        if a[1] == b[1] {
            //是否相交
            if start2[0] <= end1[0] {
                return []float64{
                    float64(start2[0]), float64(start2[1]),
                }
            }

        }
        return make([]float64, 0)
    }

    intersectPoint := getIntersectPoint(a, b)

    return in(intersectPoint, start1, end1, start2, end2)
}

func in(intersectPoint []float64, start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
    if start1[0] == end1[0] {
        if !(float64(start1[1]) <= intersectPoint[1] && intersectPoint[1] <= float64(end1[1])) {
            return make([]float64, 0)
        }
    }

    if start2[0] == end2[0] {
        if !(float64(start2[1]) <= intersectPoint[1] && intersectPoint[1] <= float64(end2[1])) {
            return make([]float64, 0)
        }
    }

    if intersectPoint[0] <= math.Max(float64(start1[0]), float64(end1[0])) && intersectPoint[0] >= math.Min(float64(start1[0]), float64(end1[0])) {
        //相交的点是否在b里
        if intersectPoint[0] <= math.Max(float64(start2[0]), float64(end2[0])) && intersectPoint[0] >= math.Min(float64(start2[0]), float64(end2[0])) {
            return intersectPoint
        }
    }

    return make([]float64, 0)
}

func change(start1 []int, end1 []int, start2 []int, end2 []int) ([]int, []int, []int, []int) {
    if start1[0] > end1[0] {
        return change(end1, start1, start2, end2)
    }
    if start2[0] > end2[0] {
        return change(start1, end1, end2, start2)
    }

    if start1[0] > start2[0] {
        return change(start2, end2, start1, end1)
    }

    if start1[0] == end1[0] && start1[1] > end1[1] {
        //左边都是竖线
        return change(end1, start1, start2, end2)
    }

    if start2[0] == end2[0] && start2[1] > end2[1] {
        //右边都是竖线
        return change(start1, end1, end2, start2)
    }

    if start2[0] == end2[0] && start1[0] == end1[0] && start1[0] == start2[0] && start1[1] > start2[1] {
        //两个都是竖线
        return change(start2, end2, start1, end1)
    }

    return start1, end1, start2, end2
}

func getEquation(start1 []int, end1 []int) []float64 {
    a := float64(start1[1]-end1[1]) / float64(start1[0]-end1[0])
    b := float64(start1[1]) - a*float64(start1[0])
    return []float64{a, b}
}

func getIntersectPoint(a, b []float64) []float64 {
    // 两条线重叠
    x := (b[1] - a[1]) / (a[0] - b[0])
    y := a[1] + x*a[0]
    return []float64{x, y}
}
