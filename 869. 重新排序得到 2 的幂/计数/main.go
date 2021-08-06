package 计数

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
import "strconv"

func reorderedPowerOf2(n int) bool {
    for _, v := range getPowerOf2Slice() {
        if isDisorderEq(n, v) {
            return true
        }
    }
    return false
}

//是否乱序相等
func isDisorderEq(a, b int) bool {
    if len(strconv.Itoa(a)) != len(strconv.Itoa(b)) {
        return false
    }
    //两个值都变成map，然后对比
    aMap := nShowFrequency(a)
    bMap := nShowFrequency(b)
    for k, v := range aMap {
        if bMap[k] != v {
            return false
        }
    }
    return true
}

func nShowFrequency(a int) map[int32]int {
    result := make(map[int32]int)
    for _, v := range strconv.Itoa(a) {
        result[v]++
    }

    return result
}

//获取所有2的幕
func getPowerOf2Slice() []int {
    result := []int{1}
    for i := 0; i <= 28; i++ { //最大到10^9
        result = append(result, 2<<i)
    }
    return result
}
