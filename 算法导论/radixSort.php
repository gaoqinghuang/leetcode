<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{
    /**
     * 基数排序
     */
    public function execute()
    {
//        $array = [16, 4, 10, 14, 7, 9, 3, 2, 8, 1];
        $array = [2, 5, 3, 0, 2, 3, 0, 3];
        dump($this->radixSort($array));
    }

    /** 基数排序,时间复杂度n
     * @param $A
     * @return array
     */
    private function radixSort($A)
    {
        $d = strlen(max($A));
        for ($i = 1; $i <= $d; $i++) {
            $A = $this->countingSort($A, $i);
        }
        return $A;
    }


    /**
     * 计数排序
     * @param $A
     * @param $w
     * @param int $k
     * @return array
     */
    private function countingSort($A, $w, $k = 9)
    {
        //C数组全部填充0
        $C = $B = [];
        for ($i = 0; $i <= $k; $i++) {
            $C[$i] = 0;
        }
        //各个位置的计数
        foreach ($A as $a) {
            $C[$this->getW($a, $w)] += 1;
        }

        //累加
        for ($i = 1; $i <= $k; $i++) {
            $C[$i] += $C[$i - 1];
        }

        for ($i = count($A) - 1; $i >= 0; $i--) {
            $B[$C[$this->getW($A[$i], $w)] - 1] = $A[$i];
            $C[$this->getW($A[$i], $w)]         -= 1;
        }
        ksort($B);
        return $B;
    }

    /**
     * 取位数
     * @param $num
     * @param $w
     * @return int
     */
    private function getW($num, $w)
    {
        return intval($num % (pow(10, $w)) / (pow(10, $w - 1)));
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