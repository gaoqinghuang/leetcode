package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

//
//给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
//
//示例 1:
//
//输入:
//"tree"
//
//输出:
//"eert"
//
//解释:
//'e'出现两次，'r'和't'都只出现一次。
//因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
//示例 2:
//
//输入:
//"cccaaa"
//
//输出:
//"cccaaa"
//
//解释:
//'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
//注意"cacaca"是不正确的，因为相同的字母必须放在一起。
//示例 3:
//
//输入:
//"Aabb"
//
//输出:
//"bbAa"
//
//解释:
//此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
//注意'A'和'a'被认为是两种不同的字符。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sort-characters-by-frequency
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//fmt.Println('z')
	//fmt.Println('a')
	//fmt.Println('Z')
	//fmt.Println('A')
	//fmt.Println('0')
	//fmt.Println('9')
	fmt.Println(frequencySort("2a554442f544asfasssffffasss"))
}

type LetterCount struct {
	Count int
	Letter byte
}


//只能排数字和字母，如果有其它符号，先转map，再转arr
func frequencySort(s string) string {
	temCount := make([]*LetterCount,75)
	for k := range temCount {
		temCount[k] = &LetterCount{
			Letter:byte(k),
		}
	}
	for _,v := range s {
		temCount[v-'0'].Count++
	}
	sort.Slice(temCount, func(i, j int) bool {
		return temCount[i].Count>temCount[j].Count
	})

	result := strings.Builder{}
	for _,v := range temCount  {
		if v.Count == 0 {
			break
		}
		for j:=1;j<=v.Count;j++ {
			result.WriteByte(v.Letter+'0')
		}

	}

	return result.String()
}

func frequencySort1(s string) string {
	cnt := map[byte]int{}
	maxFreq := 0
	for i := range s {
		cnt[s[i]]++
		maxFreq = max(maxFreq, cnt[s[i]])
	}

	buckets := make([][]byte, maxFreq+1)
	for ch, c := range cnt {
		buckets[c] = append(buckets[c], ch)
	}

	ans := make([]byte, 0, len(s))
	for i := maxFreq; i > 0; i-- {
		for _, ch := range buckets[i] {
			ans = append(ans, bytes.Repeat([]byte{ch}, i)...)
		}
	}
	return string(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

