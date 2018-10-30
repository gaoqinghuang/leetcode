<?php
/*给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。*/
function reverseWords($s)
{
    $arr = explode(' ',$s);
    foreach ($arr as &$value){
        $value = strrev($value);
    }
    return implode(' ',$arr);
}

$s = "Let's take LeetCode contest";
$result = reverseWords($s);
if("s'teL ekat edoCteeL tsetnoc" == $result){
    var_dump('success');
}else{
    var_dump('fail');
}
exit;