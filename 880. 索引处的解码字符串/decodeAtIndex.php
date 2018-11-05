<?php
heaer("Content-type:text/html;charset=utf-8");

/*  给定一个编码字符串 S。为了找出解码字符串并将其写入磁带，从编码字符串中每次读取一个字符，并采取以下步骤：
如果所读的字符是字母，则将该字母写在磁带上。
如果所读的字符是数字（例如 d），则整个当前磁带总共会被重复写 d-1 次。
现在，对于给定的编码字符串 S 和索引 K，查找并返回解码字符串中的第 K 个字母。
示例 1:

输入：S = "leet2code3", K = 10
输出："o"
解释：
解码后的字符串为 "leetleetcodeleetleetcodeleetleetcode"。
字符串中的第 10 个字母是 "o"。

示例 2：

输入：S = "ha22", K = 5
输出："h"
解释：
解码后的字符串为 "hahahaha"。第 5 个字母是 "h"。

示例 3：

输入：S = "a2345678999999999999999", K = 1
输出："a"
解释：
解码后的字符串为 "a" 重复 8301530446056247680 次。第 1 个字母是 "a"。
*/
//自己的方法是硬求变化后的字符串，但是用示例3的时候，内存爆了，后来看答案了，才写的下面这个
function decodeAtIndex($s, $k)
{
    $size = 0;
    $len  = strlen($s);

    for ($i = 0; $i < $len; $i++) {
        if (is_numeric($s[$i])) {
            $size *= $s[$i];
        } else {
            $size += 1;
        }
    }
    $s   = strrev($s);

    for ($i = 0; $i < $len; $i++) {
        $k %= $size;
        $c = $s[$i];
        if ($k == 0 && !is_numeric($c)) {
            return $c;
        }
        if (is_numeric($c)) {
            $size /= intval($c);
        } else {
            $size -= 1;
        }
    }
}


$s = 'leet2code3';
$k = 10;

if (decodeAtIndex($s, $k) == 'o') {
    echo "成功";
} else {
    echo "失败";
}


?>