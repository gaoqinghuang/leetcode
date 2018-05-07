<?php
//输入一个链表，从尾到头打印链表每个节点的值。
//思考：PHP中没有链表的概念，目前理解成一个数组，而且是只能重头访问的数组,
function printListReversingly($head)
{
    $list = array();
    $i = 0;
    while($head[$i])
    {
        array_unshift($list,$head[$i]);
        $i++;
    }
    print_r($list);
}

$head = [1,2,3,4,5];
printListReversingly($head);//因为不是链表，会产生一个数组下标超出的Notice错误