<?php

/*给定一个仅包含数字 0-9 的字符串和一个目标值，在数字之间添加二元运算符（不是一元）+、- 或 * ，返回所有能够得到目标值的表达式。
示例 1:

输入: num = "123", target = 6
输出: ["1+2+3", "1*2*3"]*/

//时间复杂度过高O(n*4^n)
function addOperators($num, $target)
{
    $count = strlen($num);
    if ($count == 0) {
        throw new \Exception('数字长度不能为0');
    } elseif ($count == 1) {
        return $target == $num ? [$num] : [];
    } else {
        $sto[] = $num[0];
        for ($i = 1; $i < $count; $i++) {
            $storage = [];
            foreach ($sto as $value) {
                $storage[] = "$value+$num[$i]";
                $storage[] = "$value-$num[$i]";
                $storage[] = "$value*$num[$i]";
                $storage[] = "$value$num[$i]";
            }

            $sto = $storage;
        }
    }
    $result = [];
    foreach ($storage as $value) {
        if (eval("return $value;") == $target) {
            $result[] = $value;
        }
    }
    return $result;
}

$reuslt = addOperators('3456237490',9191);
dump($reuslt);