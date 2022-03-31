package main

import (
	"fmt"
	"strings"
)

//给你一个二进制字符串 binary ，它仅有 0 或者 1 组成。你可以使用下面的操作任意次对它进行修改：
//
//操作 1 ：如果二进制串包含子字符串 "00" ，你可以用 "10" 将其替换。
//比方说， "00010" -> "10010"
//操作 2 ：如果二进制串包含子字符串 "10" ，你可以用 "01" 将其替换。
//比方说， "00010" -> "00001"
//请你返回执行上述操作任意次以后能得到的 最大二进制字符串 。如果二进制字符串 x 对应的十进制数字大于二进制字符串 y 对应的十进制数字，那么我们称二进制字符串 x 大于二进制字符串 y 。
//
// 
//
//示例 1：
//
//输入：binary = "000110"
//输出："111011"
//解释：一个可行的转换为：
//"000110" -> "000101"
//"000101" -> "100101"
//"100101" -> "110101"
//"110101" -> "110011"
//"110011" -> "111011"
//示例 2：
//
//输入：binary = "01"
//输出："01"
//解释："01" 没办法进行任何转换。
// 
//
//提示：
//
//1 <= binary.length <= 105
//binary 仅包含 '0' 和 '1' 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-binary-string-after-change
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(maximumBinaryString("11"))
}












func maximumBinaryString(binary string) string {
	findFristZeroK := -1
	//找到第一个0，然后看后面有几个0
	zeroN := 0
	for k,v := range binary {
		if v == '0'{
			if findFristZeroK == -1 {//0前面都是1，中间有zeroN个0
				findFristZeroK= k
			}
			zeroN++
		}
	}
	result := strings.Builder{}
	for i:=0;i<len(binary);i++ {
		if i == (findFristZeroK+zeroN-1){
			result.WriteString("0")
		}else {
			result.WriteString("1")
		}
	}
	return result.String()
}