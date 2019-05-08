<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{
    /**
     * n*n矩阵乘法
     * */
    public function execute()
    {
        $a = [
            [1, 1, 1, 1],
            [1, 1, 1, 1],
            [1, 1, 1, 12],
            [1, 1, 1, 1],
        ];

        $b = $c = $d = [
            [1, 1, 2, 1],
            [1, 1, 1, 1],
            [1, 1, 1, 1],
            [1, 1, 1, 1],
        ];
//        dump($this->squareMatrixMultiply($a, $b) == $this->squareMatrixMultiplyRecursive($a, $b) ? 'success' : 'fail');
        dump($this->squareMatrixMultiply($a, $b) == $this->strassen($a, $b) ? 'success' : 'fail');
    }

    /**
     * 暴力相乘 时间复杂度 n³
     * @param $a
     * @param $b
     * @return array
     */
    private function squareMatrixMultiply($A, $B): array
    {
        $c = [];
        $n = count($B);
        for ($i = 0; $i < $n; $i++) {
            for ($j = 0; $j < $n; $j++) {
                $res = 0;
                for ($k = 0; $k < $n; $k++) {
                    $res += $A[$i][$k] * $B[$k][$j];
                }
                $c[$i][$j] = $res;
            }
        }
        return $c;
    }


    /**
     * 递归分治 时间复杂度 n³
     * @param $A
     * @param $B
     * @return array
     */
    private function squareMatrixMultiplyRecursive($A, $B)
    {
        if (count($A) == 1) {
            return [[$A[0][0] * $B[0][0]]];
        }
        list($A1, $A2, $A3, $A4) = $this->split($A);
        list($B1, $B2, $B3, $B4) = $this->split($B);
        $a = $this->squareAdd($this->squareMatrixMultiplyRecursive($A1, $B1), $this->squareMatrixMultiplyRecursive($A2, $B3));
        $b = $this->squareAdd($this->squareMatrixMultiplyRecursive($A1, $B2), $this->squareMatrixMultiplyRecursive($A2, $B4));
        $c = $this->squareAdd($this->squareMatrixMultiplyRecursive($A3, $B1), $this->squareMatrixMultiplyRecursive($A4, $B3));
        $d = $this->squareAdd($this->squareMatrixMultiplyRecursive($A3, $B2), $this->squareMatrixMultiplyRecursive($A4, $B4));
        return $this->merge($a, $b, $c, $d);
    }


    /**
     * strassen 法，时间复杂度n^lg7,这算法有点厉害，lg8到lg7
     * @param $A
     * @param $B
     * @return array
     */
    private function strassen($A, $B)
    {
        if (count($A) == 1) {
            return [[$A[0][0] * $B[0][0]]];
        }
        list($A11, $A12, $A21, $A22) = $this->split($A);
        list($B11, $B12, $B21, $B22) = $this->split($B);
        $S1  = $this->squareReduce($B12, $B22);
        $S2  = $this->squareAdd($A11, $A12);
        $S3  = $this->squareAdd($A21, $A22);
        $S4  = $this->squareReduce($B21, $B11);
        $S5  = $this->squareAdd($A11, $A22);
        $S6  = $this->squareAdd($B11, $B22);
        $S7  = $this->squareReduce($A12, $A22);
        $S8  = $this->squareAdd($B12, $B22);
        $S9  = $this->squareReduce($A11, $A21);
        $S10 = $this->squareAdd($B11, $B12);

        $P1 = $this->squareMatrixMultiply($A11, $S1);
        $P2 = $this->squareMatrixMultiply($S2, $B22);
        $P3 = $this->squareMatrixMultiply($S3, $B11);
        $P4 = $this->squareMatrixMultiply($A22, $S4);
        $P5 = $this->squareMatrixMultiply($S5, $S6);
        $P6 = $this->squareMatrixMultiply($S7, $S8);
        $P7 = $this->squareMatrixMultiply($S9, $S10);

        $a = $this->squareAdd($this->squareReduce($this->squareAdd($P5, $P4), $P2), $P6);
        $b = $this->squareAdd($P1, $P2);
        $c = $this->squareAdd($P3, $P4);
        $d = $this->squareReduce($this->squareReduce($this->squareAdd($P5, $P1), $P3), $P7);
        return $this->merge($a, $b, $c, $d);
    }

    //矩阵加法 时间复杂度n²
    private function squareAdd($A, $B)
    {
        $c = [];
        foreach ($A as $i => $a) {
            foreach ($a as $j => $b) {
                $c[$i][$j] = $b + $B[$i][$j];
            }
        }
        return $c;
    }

    //矩阵减法 时间复杂度n²
    private function squareReduce($A, $B)
    {
        $c = [];
        foreach ($A as $i => $a) {
            foreach ($a as $j => $b) {
                $c[$i][$j] = $b - $B[$i][$j];
            }
        }
        return $c;
    }

    //拆分
    private function split($A)
    {
        $num = intval(count($A) / 2);
        $a   = $b = $c = $d = [];
        for ($i = 0; $i < $num; $i++) {
            for ($j = 0; $j < $num; $j++) {
                $a[$i][$j] = $A[$i][$j];
                $b[$i][$j] = $A[$i][$j + $num];
                $c[$i][$j] = $A[$i + $num][$j];
                $d[$i][$j] = $A[$i + $num][$j + $num];
            }
        }
        return [$a, $b, $c, $d];
    }

    //合并
    private function merge($A, $B, $C, $D)
    {
        $num = count($A);
        foreach ($B as $i => $b) {
            foreach ($b as $j => $value) {
                $A[$i][$num + $j] = $value;
            }
        }
        foreach ($C as $i => $b) {
            foreach ($b as $j => $value) {
                $A[$num + $i][$j] = $value;
            }
        }
        foreach ($D as $i => $b) {
            foreach ($b as $j => $value) {
                $A[$num + $i][$num + $j] = $value;
            }
        }
        return $A;
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