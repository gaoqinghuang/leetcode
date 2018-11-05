<?php
header("Content-type:text/html;charset=utf-8");

/*  给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
*/

function coinChange($coins, $amount)
{
    $max   = $amount + 1;
    $dp    = array_fill(0, $max, $max);
    $dp[0] = 0;
    //一定包含一个这个数，然后一个一个往上加，get
    for ($i = 1; $i <= $amount; $i++) {
        for ($j = 0; $j < count($coins); $j++) {
            if ($coins[$j] <= $i) {
                $dp[$i] = min($dp[$i], $dp[$i - $coins[$j]] + 1);
            }
        }
    }

    return $dp[$amount] > $amount ? -1 : $dp[$amount];
}

$coins  = [5, 2, 4];
$amount = 11;

if (coinChange($coins, $amount) == 3) {
    echo "成功";
} else {
    echo "失败";
}

?>