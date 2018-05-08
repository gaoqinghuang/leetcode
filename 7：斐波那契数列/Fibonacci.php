<?php
//输入一个整数n，请你输出斐波那契数列的第n项。

function fibonacci($n)
{
    $result = [0,1];
    if($n < 2 ) return $result[$n];
    $num = 0;
    $one = 0;
    $two = 1;
    for($i=2;$i<=$n;$i++)
    {
        $num = $one+$two;
        $one = $two;
        $two = $num;
    }
    return $num;
}

echo fibonacci(100);