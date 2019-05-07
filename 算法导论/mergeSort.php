<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{
    public function execute()
    {
        //并归排序，时间复杂度nlgn
        $array = [2, 3, 5, 78, 5, 41, 4, 4, 464, 65, 67, 1, 67];
        $array = $this->mergeSort($array, 0, count($array) - 1);
        dump($array);
    }

    /**
     * @param $array
     * @param $p
     * @param $r
     * @return array
     */
    private function mergeSort($array, $p, $r)
    {
        if ($p < $r) {
            $q     = floor(($p + $r) / 2);
            $array = $this->mergeSort($array, $p, $q);
            $array = $this->mergeSort($array, $q + 1, $r);
            $array = $this->merge($array, $p, $q, $r);
        }
        return $array;
    }

    /**
     * @param array $array
     * @param int $p
     * @param int $q
     * @param int $r
     * @return array
     */
    public function merge(array $array, int $p, int $q, int $r)
    {
        $left  = array_slice($array, $p, $q - $p + 1);
        $right = array_slice($array, $q + 1, $r - $q);
        //左右的初始位置
        $leftI = $rightI = 0;
        for ($i = $p; $i <= $r; $i++) {
            if (!isset($left[$leftI])) {
                $array[$i] = $right[$rightI];
                $rightI++;
                continue;
            }
            if (!isset($right[$rightI])) {
                $array[$i] = $left[$leftI];
                $leftI++;
                continue;
            }
            if ($left[$leftI] > $right[$rightI]) {
                $array[$i] = $right[$rightI];
                $rightI++;
            } else {
                $array[$i] = $left[$leftI];
                $leftI++;
            }
        }
        return $array;
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