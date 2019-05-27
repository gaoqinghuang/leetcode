<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{

    /**
     *钢条切割,给定长度为n的钢条，对应的长度价格为p,问怎么切割钢条可以获取最大利润
     */
    public function execute()
    {
        $p = [
            1  => 1,
            2  => 5,
            3  => 8,
            4  => 9,
            5  => 10,
            6  => 17,
            7  => 17,
            8  => 20,
            9  => 24,
            10 => 30,
        ];
        $n = 9;
        list($r, $s) = $this->extendedBottomUpCupRod($p, $n);
        $res = [];
        while ($n > 0) {
            $res[] = $s[$n];
            $n     = $n - $s[$n];
        }
        dump($r, $res);
    }

    /**
     * 自底向上法，O(n^2),记录最优解
     * @param $p
     * @param $n
     * @return mixed
     */
    private function extendedBottomUpCupRod($p, $n)
    {
        $r = [0];
        $s = [0];
        for ($j = 1; $j <= $n; $j++) {
            $q = -1;
            for ($i = 1; $i <= $j; $i++) {
                if ($q > $p[$i] + $r[$j - $i]) {
                } else {
                    $q     = $p[$i] + $r[$j - $i];
                    $s[$j] = $i;
                }
            }
            $r[$j] = $q;
        }
        return [$r[$n], $s];
    }

    /**
     * 自底向上法，O(n^2)
     * @param $p
     * @param $n
     * @return mixed
     */
    private function bottomCupRod($p, $n)
    {
        $r = [0];
        for ($j = 1; $j <= $n; $j++) {
            $q = -1;
            for ($i = 1; $i <= $j; $i++) {
                $q = max($q, $p[$i] + $r[$j - $i]);
            }
            $r[$j] = $q;
        }
        return $r[$n];
    }

    /**
     * 这里的p的价格从1对应到n,带备忘  O(n^2)
     * @param $p
     * @param $n
     * @return int|mixed
     */
    private function memoizedCutRod($p, $n)
    {
        static $linshi = [];
        if ($n == 0) {
            return 0;
        }
        $q = -1;
        for ($i = 1; $i <= $n; $i++) {
            if (!isset($linshi[$n - $i])) {
                $linshi[$n - $i] = $this->cutRod($p, $n - $i);
            }
            $q = max($q, $p[$i] + $linshi[$n - $i]);
        }
        return $q;
    }

    /**
     * 这里的p的价格从1对应到n  O(2^10)
     * @param $p
     * @param $n
     * @return int|mixed
     */
    private function cutRod($p, $n)
    {
        if ($n == 0) {
            return 0;
        }
        $q = -1;
        for ($i = 1; $i <= $n; $i++) {
            $q = max($q, $p[$i] + $this->cutRod($p, $n - $i));
        }
        return $q;
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