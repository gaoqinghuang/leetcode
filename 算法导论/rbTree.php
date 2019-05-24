<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');

define('BLOCK', 'block');
define('RED', 'red');

class help
{
    public function execute()
    {
        //头晕目眩，后期还是得再回来看下，就是情况复杂了点，基本就是整理下有多少种情况，然后采取某种固定的性质策略，结点往上升，整理
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
     *
     * 没看得很详细,基础解决的办法是:
     * 插入的时候是在叶子节点上，然后每次循环都保持一种状态，从叶子结点往上修复，只到最终修复
     *
     * @param $T
     * @param elementList $z
     */
    private function rbInsert($T, elementList $z)
    {
        $x = $T->root;
        $y = $T->nil;
        while ($x->key != null) {
            $y = $x;
            if ($x->key < $z->key) {
                $x = $x->right;
            } else {
                $x = $x->left;
            }
        }
        $z->p     = $y;
        $z->right = $T->nil;
        $z->left  = $T->nil;
        if ($y->key == null) {
            $T->root = $z;
        } elseif ($z->key < $y->key) {
            $y->left = $z;
        } else {
            $y->right = $z;
        }
        $z->color = RED;
        $this->rbInsertFixup($T, $z);
    }

    /**
     * 插入后重新着色
     * @param $T
     * @param $z
     */
    private function rbInsertFixup($T, $z)
    {
        while ($z->p->color == RED) {
            if ($z->p == $z->p->p->left) {
                $y = $z->p->p->right;
                if ($y->color == RED) {
                    $z->p->color    = BLOCK;
                    $y->color       = BLOCK;
                    $z->p->p->color = RED;
                    $z              = $z->p->p;
                } elseif ($z == $z->p->right) {
                    $z = $z->p;
                    $this->leftRotate($T, $z);
                } else {
                    $z->p->color    = BLOCK;
                    $z->p->p->color = RED;
                    $this->rightRotate($T, $z->p->p);
                }
            } else {
                $y = $z->p->p->left;
                if ($y->color == RED) {
                    $z->p->color    = BLOCK;
                    $y->color       = BLOCK;
                    $z->p->p->color = RED;
                    $z              = $z->p->p;
                } elseif ($z == $z->p->left) {
                    $z = $z->p;
                    $this->rightRotate($T, $z);
                } else {
                    $z->p->color    = BLOCK;
                    $z->p->p->color = RED;
                    $this->leftRotate($T, $z->p->p);
                }
            }
        }
        $T->root->color = BLOCK;
    }

    /**
     * 替换
     * @param $T
     * @param $u //被替换者
     * @param $v //替换者
     */
    private function rbTransplant($T, $u, $v)
    {
        if ($u->p->key == null) {
            $T->root = $v;
        } elseif ($u == $u->p->left) {
            $u->p->left = $v;
        } else {
            $u->p->right = $v;
        }
        $v->p = $u->p;
    }

    /**
     * 删除
     * @param $T
     * @param $z
     */
    private function rbDelete($T, $z)
    {
        $y              = $z;
        $yOriginalColor = $y->color;
        if ($z->left->key == null) {
            $this->rbTransplant($T, $z, $z->right);
        } elseif ($z->right->key == null) {
            $this->rbTransplant($T, $z, $z->left);
        } else {
            $y              = $this->treeMinimum($z->right);
            $yOriginalColor = $y->color;
            $x              = $y->right;
            if ($y->p != $z) {
                $this->rbTransplant($T, $y, $y->right);
                $y->right    = $z->right;
                $y->right->p = $y;
            } else {
                $x->p = $y;
            }
            $this->rbTransplant($T, $z, $y);
            $y->left    = $z->left;
            $y->left->p = $y;
            $y->color   = $z->color;
        }
        if ($yOriginalColor == BLOCK) {
            $this->rbDeleteFixup($T, $x);
        }
    }

    /**
     * 删除后着色
     * @param $T
     * @param $x
     */
    private function rbDeleteFixup($T, $x)
    {
        while ($x != $T->root && $x->color == BLOCK) {
            if ($x == $x->p->left) {
                $w = $x->p->right;
                if ($w->color == RED) {
                    $w->color    = BLOCK;
                    $x->p->color = RED;
                    $this->leftRotate($T, $x->p);
                    $w = $x->p->right;
                }
                if ($w->left->color == BLOCK && $w->right->color == BLOCK) {
                    $w->color = RED;
                    $x        = $x->p;
                } else {
                    if ($w->right->color == BLOCK) {
                        $w->left->color = BLOCK;
                        $w->color       = RED;
                        $this->rightRotate($T, $w);
                        $w = $x->p->right;
                    }
                    $w->color        = $x->p->color;
                    $x->p->color     = BLOCK;
                    $w->right->color = BLOCK;
                    $this->leftRotate($T, $x->p);
                    $x = $T->root;

                }
            } else {
                if ($x == $x->p->right) {
                    $w = $x->p->left;
                    if ($w->color == RED) {
                        $w->color    = BLOCK;
                        $x->p->color = RED;
                        $this->rightRotate($T, $x->p);
                        $w = $x->p->left;
                    }
                    if ($w->right->color == BLOCK && $w->left->color == BLOCK) {
                        $w->color = RED;
                        $x        = $x->p;
                    } else {
                        if ($w->left->color == BLOCK) {
                            $w->right->color = BLOCK;
                            $w->color        = RED;
                            $this->leftRotate($T, $w);
                            $w = $x->p->left;
                        }
                        $w->color       = $x->p->color;
                        $x->p->color    = BLOCK;
                        $w->left->color = BLOCK;
                        $this->rightRotate($T, $x->p);
                        $x = $T->root;
                    }

                }
            }
        }
        $x->color = BLOCK;
    }

    /**
     * 左旋
     * @param $T
     * @param $x
     */
    private function leftRotate($T, $x)
    {
        $y        = $x->right;
        $x->right = $y->left;
        if ($y->left->key != null) {
            $y->left->p = $x;
        }

        $y->p = $x->p;
        if ($x->p->key == null) {
            $T->root = $y;
        } elseif ($x == $x->p->left) {
            $x->p->left;
        } else {
            $x->p->right = $y;
        }
        $y->left = $x;
        $x->p    = $y;
    }

    /**
     * 右旋
     * @param $T
     * @param $x
     */
    private function rightRotate($T, $x)
    {
        $y       = $x->left;
        $x->left = $y->right;
        if ($y->right->key != null) {
            $y->right->p = $x;
        }

        $y->p = $x->p;
        if ($x->p->key == null) {
            $T->root = $y;
        } elseif ($x == $x->p->right) {
            $x->p->right;
        } else {
            $x->p->left = $y;
        }
        $y->right = $x;
        $x->p     = $y;
    }
}

class elementList
{
    public $left;
    public $right;
    public $p;
    public $key;
    public $color = BLOCK;

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
     * 红黑树
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