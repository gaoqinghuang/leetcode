<?php
header("Content-type:text/html;charset=utf-8");

/*  给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。
如果小数部分为循环小数，则将循环的部分括在括号内。
示例 1:

输入: numerator = 1, denominator = 2
输出: "0.5"

示例 2:

输入: numerator = 2, denominator = 1
输出: "2"

示例 3:

输入: numerator = 2, denominator = 3
输出: "0.(6)"
*/

function fractionToDecimal($a, $b)
{
    //先求出整数部分
    $integer = intval($a / $b);
    //整数后的余数部分
    $rem = $a % $b;
    if ($rem == 0) return $integer;
    $ans = $integer . '.';

    $t   = [];
    $cou = 0;
    while ($rem) {
        //再来分析小数部分，解析循环体
        if (isset($t[$rem])) {
            //循环体
            $count       = end($t) - $t[$rem] + 1;
            $stringCount = strlen($ans);
            $ans         = substr($ans, 0, $stringCount - $count) . '(' . substr($ans, -$count) . ')';
            break;
        }

        $cou++;
        $t[$rem]  = $cou;
        $xInteger = intval($rem * 10 / $b);
        $ans      = $ans . $xInteger;
        $rem      = $rem * 10 % $b;
    }
    return $ans;
}

$a = 2;
$b = 3;

if (fractionToDecimal($a, $b) == "0.(6)") {
    echo "成功";
} else {
    echo "失败";
}

?>