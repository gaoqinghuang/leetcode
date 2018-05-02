<?php
/*在长度为n的数组中，所有的元素都是0到n-1的范围内。
数组中的某些数字是重复的，但不知道有几个重复的数字，
也不知道重复了几次，请找出任意重复的数字。 例如，
输入长度为7的数组{2,3,1,0,2,5,3}，那么对应的输出为2或3*/

//直接用php中的函数实现
/*function duplicateNumber($data)
{
    $unique_arr = array_unique($data);
    return array_diff_assoc($data,$unique_arr);
}*/
//foreach循环一次,时间复杂度O(n),此方法更好
function duplicateNumber($data)
{
    $newData = [];
    $returnDate = [];
    foreach($data as $key=>$value)
    {
        if(isset($newData[$value]))
        {
            $returnDate[] = $value;
        }
        else
        {
            $newData[$value] = $value;
        }
    }
    return  $returnDate;
}
$data=array(2,3,1,0,2,5,3);
print_r(duplicateNumber($data));