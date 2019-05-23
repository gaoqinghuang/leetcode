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


        $T = new tree($a);

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

//        dump($this->inorderTreeWalk($a));
//        dump($this->treeSearch($a, 5));
//        dump($this->treeMinimum($a)->key);
//        dump($this->treeSuccessor($f)->key);
//        $z = new elementList(12);
//        $this->treeInsert($T, $z);
//        dump($this->inorderTreeWalk($a));
//        $z = new elementList(null);
//        $this->transplant($T, $b, $z);
//        dump($this->inorderTreeWalk($T->root));

        $this->treeDelete($T, $a);
        dump($this->inorderTreeWalk($T->root));
    }


    /**
     * 中序遍历排序
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

    /**
     * 搜索
     * @param $x
     * @param $k
     * @return mixed
     */
    private function treeSearch($x, $k)
    {
        if ($x->key == null || $k == $x->key) {
            return $x;
        }
        if ($k < $x->key) {
            return $this->treeSearch($x->left, $k);
        }
        return $this->treeSearch($x->right, $k);
    }

    /**
     * 查找最小关键字
     * @param $x
     * @return mixed
     */
    private function treeMinimum($x)
    {
        while ($x->left->key != null) {
            $x = $x->left;
        }
        return $x;
    }

    /**
     * 查找最大关键字
     * @param $x
     * @return mixed
     */
    private function treeMaximum($x)
    {
        while ($x->right->key != null) {
            $x = $x->right;
        }
        return $x;
    }

    /**
     * 寻找后继
     * @param $x
     * @return mixed
     */
    private function treeSuccessor($x)
    {
        if ($x->right->key != null) {
            return $this->treeMinimum($x->right);
        }
        $y = $x->p;
        while ($y->key != null && $x == $y->right) {
            $x = $y;
            $y = $y->p;
        }
        return $y;
    }

    /**
     * 插入
     * @param $T
     * @param elementList $z
     */
    private function treeInsert($T, elementList $z)
    {
        $x = $T->root;
        $y = new elementList(null);
        while ($x->key != null) {
            $y = $x;
            if ($x->key < $z->key) {
                $x = $x->right;
            } else {
                $x = $x->left;
            }
        }
        $z->p     = $y;
        $z->right = new elementList(null);
        $z->left  = new elementList(null);
        if ($y->key == null) {
            $T->root = $z;
        } elseif ($z->key < $y->key) {
            $y->left = $z;
        } else {
            $y->right = $z;
        }
    }

    /**
     * 替换
     * @param $T
     * @param $u //被替换者
     * @param $v //替换者
     */
    private function transplant($T, $u, $v)
    {
        if ($u->p->key == null) {
            $T->root = $v;
        } elseif ($u == $u->p->left) {
            $u->p->left = $v;
        } else {
            $u->p->right = $v;
        }
        if ($v->key != null) {
            $v->p = $u->p;
        }
    }

    /**
     * 删除
     * @param $T
     * @param $z
     */
    private function treeDelete($T, $z)
    {
        if ($z->left->key == null) {
            $this->transplant($T, $z, $z->right);
        } elseif ($z->right->key == null) {
            $this->transplant($T, $z, $z->left);
        } else {
            $y = $this->treeMinimum($z->right);
            if ($y->p != $z) {
                $this->transplant($T, $y, $y->right);
                $y->right    = $z->right;
                $y->right->p = $y;
            }
            $this->transplant($T, $z, $y);
            $y->left    = $z->left;
            $y->left->p = $y;
        }
    }
}

class elementList
{
    public $left;
    public $right;
    public $p;
    public $key;

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

    /**
     * 搜索二叉树
     * tree constructor.
     * @param elementList $x
     */
    public function __construct($x)
    {
        $this->root = $x;
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