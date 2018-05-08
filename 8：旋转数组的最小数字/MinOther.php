<?php
/*把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。 输入一个非递减排序的数组的一个旋转，
输出旋转数组的最小元素。 例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。*/
//直接用min()函数时间复杂度为O(n),下面这个方法大部分复杂度为O(logn)
function minOther($data)
{
    $length = count($data);
    if($length == 0)
    {
        return false;
    }
    $index1 = 0;
    $index2 = $length-1;
    $medIndex = $index1;
    while($data[$index1] >= $data[$index2])
    {
        if($index2-$index1 == 1)
        {
            $medIndex = $index2;
            break;
        }
        $medIndex = ($index2+$index1)>>1;
        if($data[$index1] == $data[$index2] and $data[$index1] == $data[$medIndex])
        {
            return min($data);
        }
        if($data[$medIndex] >= $data[$index1])
        {
            $index1 = $medIndex;
        }elseif($data[$medIndex] <= $data[$index2])
        {
            $index2 = $medIndex;
        }
    }
    return $data[$medIndex];
}

$data = [7,9,10,11,1,3,5];
if(minOther($data) !== 1)
{
    echo '出错';
}
$data = [1,0,1,1,1];
if(minOther($data) !== 0)
{
    echo '出错';
}