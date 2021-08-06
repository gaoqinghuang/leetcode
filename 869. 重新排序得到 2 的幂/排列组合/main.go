package 排列组合

import "strconv"

//给定正整数 N ，我们按任何顺序（包括原始顺序）将数字重新排序，注意其前导数字不能为零。
//
//如果我们可以通过上述方式得到 2 的幂，返回 true；否则，返回 false。
//
// 
//
//示例 1：
//
//输入：1
//输出：true
//示例 2：
//
//输入：10
//输出：false
//示例 3：
//
//输入：16
//输出：true
//示例 4：
//
//输入：24
//输出：false
//示例 5：
//
//输入：46
//输出：true
// 
//
//提示：
//
//1 <= N <= 10^9
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reordered-power-of-2
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func reorderedPowerOf2(n int) bool {
    for _, number := range getReordered(strconv.Itoa(n)) {
        //验证是不是2的幕
        if isPowerOf2(number) {
            return true
        }
    }
    return false
}



//获取各种排序
func getReordered(n string) []string {
    if len(n) == 1 {
        return []string{n}
    }
    result := make([]string, 0)
    for num, value := range n {
        for _, nSlice := range getReordered(n[0:num]+n[num+1:]) {
            result = append(result, string(value)+nSlice)
        }
    }
    return result
}

//是否是2的幂
func isPowerOf2(nString string) bool {
    if nString[0:1] == "0" {
        return false
    }
    n, _ := strconv.Atoi(nString)
    for true {
        if n == 1 {
            return true
        }

        if n <= 0 {
            return false
        }

        if n%2 == 1 {
            return false
        }
        n = n >> 1
    }
    return false
}