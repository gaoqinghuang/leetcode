<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{

    /**
     * 堆的实现,要去除0的位置，&为占位符，不然下面的left,right,parent就不准了
     */
    public function execute()
    {
        $array = ['&', 16, 4, 10, 14, 7, 9, 3, 2, 8, 1];
        //构建堆
        $heap            = new stdClass();
        $heap->array     = $array;
        $heap->length    = count($array) - 1;//长度
        $heap->heap_size = $heap->length;//有多少堆元素存储在数组中

        $this->heapSort($heap);
        dump($heap);
    }

    /**
     * 堆排序 时间复杂度 lgn*n
     * @param $array
     */
    public function heapSort(&$heap)
    {
        $this->buildMaxHeap($heap);
        for ($i = $heap->length; $i >= 2; $i--) {
            $this->exchange($heap->array[1], $heap->array[$i]);
            $heap->heap_size -= 1;
            $this->maxHeapify($heap, 1);
        }
    }

    /**
     * 最大堆，时间复杂度 n
     * @param $array
     */
    private function buildMaxHeap(&$heap)
    {
        $heap->heap_size = $heap->length;
        for ($i = $heap->length >> 1; $i >= 1; $i--) {
            $this->maxHeapify($heap, $i);
        }
    }


    /**
     * 已知$i的左右都是最大堆，让$i为最大堆  时间复杂度 lgn
     * @param $array
     * @param $i
     */
    private function maxHeapify(&$heap, $i)
    {
        $l = $this->left($i);
        $r = $this->right($i);

        $largest = $i;
        if ($l <= $heap->heap_size && $heap->array[$l] > $heap->array[$i]) {
            $largest = $l;
        }
        if ($r <= $heap->heap_size && $heap->array[$r] > $heap->array[$largest]) {
            $largest = $r;
        }
        if ($largest != $i) {
            $this->exchange($heap->array[$i], $heap->array[$largest]);
            $this->maxHeapify($heap, $largest);
        }
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

    /**
     * 获取父节点
     * @param $i
     * @return int
     */
    private function parent($i)
    {
        return $i >> 1;
    }

    /**
     * 获取左节点
     * @param $i
     * @return int
     */
    private function left($i)
    {
        return $i << 1;
    }

    /**
     * 获取右节点
     * @param $i
     * @return int
     */
    private function right($i)
    {
        return ($i << 1) + 1;
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