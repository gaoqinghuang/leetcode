<?php
/*在一个二维数组中，每一行都按照从左到右递增的顺序排序，
每一列都按照从上到下递增的顺序排序。请完成一个函数，
输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。*/
function find($data,$num)
{
    $found = false;
    $rows = count($data);
    $cols = count($data[0]);
    if($rows > 0 and $cols > 0)
    {
        $row = 0;
        $col = $cols-1;
        while($row < $rows and $col>=0)
        {
            if($data[$row][$col] == $num)
            {
                $found = true;
                break;
            }
            elseif($data[$row][$col] > $num)
            {
                $col--;
            }
            else
            {
                $row++;
            }
        }
    }
    return $found;
}

$data = [
    [1,2,8,9],
    [2,4,9,12],
    [4,7,10,13],
    [6,8,11,15],
];
echo find($data,16)? '是':'否';