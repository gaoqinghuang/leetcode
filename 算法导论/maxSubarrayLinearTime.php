<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{
    /**
     * 寻找最大子数组的和，分治算法  时间复杂度nlgn
     * */
    public function execute()
    {
        $array = [13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7];
        $max   = $this->findMaximumSubarray($array, 0, count($array) - 1);
        dump($max);
    }

    /**
     * 寻找最大子数组的和
     * @param $array
     * @param $low
     * @param $high
     * @return int
     */
    private function findMaximumSubarray($array, $low, $high): int
    {
        if ($low == $high) {
            return $array[$low];
        }
        $mid = floor(($low + $high) / 2);
        //要嘛在左
        $leftMax = $this->findMaximumSubarray($array, $low, $mid);
        //要嘛在右
        $rightMax = $this->findMaximumSubarray($array, $mid + 1, $high);
        //要嘛在中间
        $midMax = $this->findMaxCrossiongSubarray($array, $low, $mid, $high);
        return max($leftMax, $rightMax, $midMax);
    }

    /**
     * 寻找跨越中间值最大子数组的和 时间复杂度n
     * @param $array
     * @param $low
     * @param $mid
     * @param $high
     * @return int
     */
    private function findMaxCrossiongSubarray($array, $low, $mid, $high): int
    {
        //左边
        $leftMax = $rightMax = $now = 0;
        for ($i = $mid; $i >= $low; $i--) {
            $now     += $array[$i];
            $leftMax = max($now, $leftMax);
        };

        //右边
        $now = 0;
        for ($i = $mid + 1; $i <= $high; $i++) {
            $now      += $array[$i];
            $rightMax = max($now, $rightMax);
        }
        return $leftMax + $rightMax;
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