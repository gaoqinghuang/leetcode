<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');


class help
{
    public function execute()
    {

    }

    /**
     * 计算一个元素的秩  O(lgn)
     * @param $T
     * @param $x
     * @return mixed
     */
    private function osRank($T, $x)
    {
        $r = $x->left->size + 1;
        $y = $x;
        while ($y != $T->root) {//3
            if ($y == $y->p->right) {
                $r = $r + $y->p->left->size + 1;
            }
            $y = $y->p;
        }
        return $r;
    }

    /**
     * 查找第i小的的关键结点 时间复杂度为 lgn
     * @param $x
     * @param $i
     * @return mixed
     */
    private function osSelect($x, $i)
    {
        $r = $x->left->size + 1;
        if ($r == $i) {
            return $x;
        } elseif ($i < $r) {
            return $this->osSelect($x->left, $i);
        } else {
            return $this->osSelect($x->right, $i);
        }
    }
}

class elementList
{
    public $left;
    public $right;
    public $p;
    public $key;
    public $color = BLOCK;
    public $size = 0;

    /**
     * 链表里的各个元素
     * elementList constructor.
     * @param $key
     */
    public function __construct($key)
    {
        $this->key = $key;
    }
}


class tree
{
    public $root;
    public $nil;

    /**
     * 顺序统计树
     * tree constructor.
     * @param elementList $x
     */
    public function __construct($x)
    {
        $this->root = $x;
        $this->nil  = new elementList(null);
        $x->p       = $this->nil;
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