<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{
    /**
     * 计数排序
     */
    public function execute()
    {
//        $array = [16, 4, 10, 14, 7, 9, 3, 2, 8, 1];
        $array = [2, 5, 3, 0, 2, 3, 0, 3];
        dump($this->countingSort($array, max($array)));
    }

    /**
     * 计数排序,时间复杂度n,站的空间太多了
     * @param $A
     * @param $k
     * @return array
     */
    private function countingSort($A, $k)
    {
        //C数组全部填充0
        $C = $B = [];
        for ($i = 0; $i <= $k; $i++) {
            $C[$i] = 0;
        }
        //各个位置的计数
        foreach ($A as $a) {
            $C[$a] += 1;
        }
        //累加
        for ($i = 1; $i <= $k; $i++) {
            $C[$i] += $C[$i - 1];
        }
        for ($i = count($A) - 1; $i >= 0; $i--) {
            $B[$C[$A[$i]] - 1] = $A[$i];
            $C[$A[$i]]         = $C[$A[$i]] - 1;
        }
        //这里的ksort 在很多语言中是不需要的，这里是为了输出是清晰
        ksort($B);
        return $B;
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