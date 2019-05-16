<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{
    /**
     * 桶排序
     */
    public function execute()
    {
        $array = [16, 4, 40, 14, 7, 9, 3, 2, 8, 1];
//        $array = [2, 5, 3, 0, 2, 3, 0, 3];
        dump($this->bucketSort($array));
    }

    private function bucketSort($A)
    {
        $n = count($A);
        $B = [];
        for ($i = 0; $i <= ($n - 1); $i++) {
            $B = [];
        }
        foreach ($A as $a) {
            $B[intval($a / $n)][] = $a;
        }
        foreach ($B as &$b) {
            $b = $this->insertionSort($b);
        }
        ksort($B);
        $C = [];
        foreach ($B as $value) {
            $C = array_merge($C, $value);
        }
        return $C;
    }

    /**
     * 插入排序
     * @param $a
     * @return mixed
     */
    public function insertionSort($a)
    {
        foreach ($a as $key => $value) {
            if ($key == 0) {
                continue;
            }
            $i = $key - 1;
            while ($i >= 0 && $a[$i] > $value) {
                $a[$i + 1] = $a[$i];
                $i         = $i - 1;
            }
            $a[$i + 1] = $value;
        }
        return $a;
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