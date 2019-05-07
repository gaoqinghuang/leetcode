<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{
    public function execute()
    {
        $this->standard();//n²
        $this->hand();//n² 空间占的少
//        $this->table();//n² 空间占的多
    }

    public function table()
    {
        //插入排序(从桌上拿起牌的)
        $a = [2,3,5,78,5,41,4,4,464,65,67,1,67];
        $b = [];//代存入的数组
        foreach ($a as $value){
            $n = null;
            foreach ($b as $key => $val){
                if($value < $val){
                    $n = $key;
                    break;
                }
            }
            if(is_null($n)){
                $b[] = $value;
            }else{
                array_splice($b,$n,0,$value);
            }
        }
        dump($a,$b);
    }

    public function hand()
    {
        //插入排序(整理手上牌的)
        $a = [2,3,5,78,5,41,4,4,464,65,67,1,67];
        foreach ($a as $key=>$value) {
            if ($key == 1) {
                continue;
            }
            foreach ($a as $ke => $val){
                if($ke>=$key){
                    break;
                }
                if($value<$val){
                    unset($a[$key]);
                    array_splice($a,$ke,0,$value);
                    break;
                }
            }
        }
        dump([2,3,5,78,5,41,4,4,464,65,67,1,67],$a);
    }

    public function standard()
    {
        //插入排序(官方的)
        $a = [2,3,5,78,5,41,4,4,464,65,67,1,67];
        foreach ($a as $key=>$value) {
            if($key == 0){
                continue;
            }
            $i = $key-1;
            while($i >=0 && $a[$i] > $value){
                $a[$i+1] = $a[$i];
                $i = $i -1;
            }
            $a[$i+1] = $value;
        }
        dump([2,3,5,78,5,41,4,4,464,65,67,1,67],$a);
    }
}
































function dump(...$var)
{
    header("Content-type:text/html;charset=utf-8");
    ob_start();
    var_dump(...$var);
    $output = preg_replace('/\]\=\>\n(\s+)/m', '] => ', ob_get_clean());
    $output = '<pre>' . $output . '</pre>';
    echo($output);
    die;
}

$help = new help();

$help->execute();