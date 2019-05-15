<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{

    /**
     * 快速排序
     */
    public function execute()
    {
        $array = [16, 4, 10, 14, 7, 9, 3, 2, 8, 1];
        $this->quickSort($array, 0, count($array) - 1);
        dump($array);
    }

    /**
     * 快排 时间复杂度n*lgn
     * @param $array
     * @param $p
     * @param $r
     */
    private function quickSort(&$array, $p, $r)
    {
        if ($p < $r) {
            $q = $this->partition($array, $p, $r);
            $this->quickSort($array, $p, $q - 1);
            $this->quickSort($array, $q + 1, $r);
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