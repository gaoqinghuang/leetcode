<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{
    /**
     * 寻找中位数
     */
    public function execute()
    {
        $array = [16, 4, 40, 14, 7, 9, 3, 2, 8, 1];
//                $array = [0,2, 5, 3, 0, 2, 3, 0, 3];
        $array1 = $array;
        sort($array1);
        $mid = $array1[intval((count($array) + 1) / 2) - 1];
        dump($mid == $this->randomizedSelect($array, 0, count($array) - 1, intval((count($array) + 1) / 2)));
    }


    /**
     * 寻找第i小的元素,时间复杂度n^2,期望运行时间n
     * @param $A
     * @param $p
     * @param $r
     * @param $i
     * @return integer
     */
    private function randomizedSelect($A, $p, $r, $i)
    {
        if ($p == $r) {
            return $A[$p];
        }
        $q = $this->partition($A, $p, $r);
        $k = $q - $p + 1;
        if ($k == $i) {
            return $A[$q];
        } else if ($i < $k) {
            return $this->randomizedSelect($A, $p, $q - 1, $i);
        } else {
            return $this->randomizedSelect($A, $q + 1, $r, $i - $k);
        }
    }

    /**
     * 选取一个值作为中间值，并让左边小于中间值，右边的大于  时间复杂度n
     * @param $array
     * @param $p
     * @param $r
     * @return int
     */
    private function partition(&$array, $p, $r)
    {
        $x = $array[$r];
        $i = $p - 1;
        for ($j = $p; $j <= ($r - 1); $j++) {
            if ($array[$j] <= $x) {
                $i++;
                $this->exchange($array[$i], $array[$j]);
            }
        }
        $this->exchange($array[$i + 1], $array[$r]);

        return $i + 1;
    }

    /**
     * 交换a,b两个的值
     * @param $a
     * @param $b
     */
    private function exchange(&$a, &$b)
    {
        list($a, $b) = [$b, $a];
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