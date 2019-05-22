<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');

class help
{
    public function execute()
    {
        //还没看到构建那章，先代码写死构建
        $nil = new elementList(null);
        $a   = new elementList(6);
        $b   = new elementList(5);
        $c   = new elementList(7);
        $d   = new elementList(2);
        $e   = new elementList(5);
        $f   = new elementList(8);

        $a->p     = $nil;
        $a->left  = $b;
        $a->right = $c;

        $b->p     = $a;
        $b->left  = $d;
        $b->right = $e;

        $c->p     = $a;
        $c->left  = $nil;
        $c->right = $f;

        $d->p     = $b;
        $d->left  = $nil;
        $d->right = $nil;

        $e->p     = $b;
        $e->left  = $nil;
        $e->right = $nil;

        $f->p     = $c;
        $f->left  = $nil;
        $f->right = $nil;

        dump($this->inorderTreeWalk($a));
    }


    /**
     * 搜索二叉树中序遍历排序
     * @param $x
     * @return array
     */
    private function inorderTreeWalk($x)
    {
        static $relust = [];
        if ($x->key != null) {
            $this->inorderTreeWalk($x->left);
            $relust[] = $x->key;
            $this->inorderTreeWalk($x->right);
        }
        return $relust;
    }
}

class elementList
{
    public $left;//左数
    public $right;//右树
    public $p;//父树
    public $key;//值

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