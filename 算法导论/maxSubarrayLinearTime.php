<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{
    /**
     * 寻找最大子数组的和，动态规划  时间复杂度n
     * */
    public function execute()
    {
        $array     = [];
        $array[]   = [1, 2, 3, 4, 5, 6];//21
        $array[]   = [0, 0, 18, 0, 0, -6];//18
        $array[]   = [13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7];//43
        $array[]   = [-2, 11, -4, 13, -5, -2];//20
        $array[]   = [-2, 11, -4, 13, -5, -2];//20
        $relust    = [21, 18, 43, 20, 20];
        $isSuccess = true;
        foreach ($array as $key => $arr) {
            if ($this->maxSubarrayLinearTime($arr) != $relust[$key]) {
                $isSuccess = false;
                $errorMes  = "第{$key}个错误";
            }
        }

        if (!$isSuccess) {
            echo $errorMes;
        } else {
            echo '成功啦';
        }
    }

    /**
     * @param $array
     * @return int
     */
    private function maxSubarrayLinearTime($array)
    {
        $max = $now = 0;
        foreach ($array as $value) {
            $now = $now + $value;
            if ($now > $max) {
                $max = $now;
            }
            if ($now <= 0) {
                $now = 0;
            }

        }
        return $max;
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