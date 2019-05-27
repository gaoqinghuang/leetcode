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
     * 寻找T中与区间i重叠的结点
     * @param $T
     * @param $i
     * @return mixed
     */
    private function intervalSearch($T, $i)
    {
        $x = $T->root;
        while ($x != $T->nil && !($i->low <= $x->int->high && $x->int->low <= $i->high)) {
            if ($x->left != $T->nil && $x->left->max >= $i->low) {
                $x = $x->left;
            } else {
                $x = $x->right;
            }
        }
        return $x;
    }

}

class elementList
{
    public $left;
    public $right;
    public $p;
    public $int;
    public $max;

    /**
     * 链表里的各个元素
     * elementList constructor.
     * @param $key
     */
    public function __construct($key)
    {
        $int       = new Section();
        $int->low  = $key;
        $this->int = $int;
    }
}


class tree
{
    public $root;
    public $nil;

    /**
     * 区间树
     * tree constructor
     * @param elementList $x
     */
    public function __construct($x)
    {
        $this->root = $x;
        $this->nil  = new elementList(null);
        $x->p       = $this->nil;
    }
}

class Section
{
    public $low;
    public $high;
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