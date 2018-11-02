<?php
header("Content-type:text/html;charset=utf-8");

/*  给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
*/

function coinChange($coins, $amount)
{
    $result = [];
    rsort($coins);
    $max = $coins[0];
    $res = [$max];

    while (true) {
        $sum = array_sum($res);
        if ($sum > $amount) {
            //换个小点的数
            $end = end($res);
            $key = array_search($end, $coins);
            if (isset($coins[$key + 1])) {
                //最后一个key换掉
                array_pop($res);
                $res[] = $coins[$key + 1];
            } else {
                break;
            }
        } elseif ($sum == $amount) {
            $result = $res;
            break;
        } else {
            //小于的话加上一个自己
            $res[] = end($res);
        }
    }
    return $result ? count($result) : -1;
}

$coins  = [1, 4, 2];
$amount = 11;

if (coinChange($coins, $amount) == 4) {

    echo "成功";
} else {
    echo "失败";
}

?>