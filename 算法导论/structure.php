<?php
header("Content-type:text/html;charset=utf-8");
ini_set('display_errors', 'on');
ini_set('error_reporting', '-1');
ini_set('display_startup_errors', 'on');

class help
{
    public function execute()
    {
//        $array = [16, 4, 40, 14, 7, 9, 3, 2, 8, 1];
        $array = [];
        $queue = new linkList($array);
        $x     = new elementList(5);
        $queue->insert($x);
        $a = $queue->search(5);
        $queue->delete($a);
        $a = $queue->search(4);
        dump($a->next->key);
    }

}


class stack
{
    public $top;
    public $array;

    /**
     * 栈
     * stack constructor.
     * @param $array
     */
    public function __construct(array $array)
    {
        //key从1开始
        $array = array_merge([null], $array);
        unset($array[0]);
        $this->array = $array;
        $this->top   = count($array);
    }

    /**
     * 是否为空
     * @return bool
     */
    public function empty()
    {
        return $this->top == 0;
    }

    /**
     * 入栈
     * @param $x
     */
    public function push($x)
    {
        $this->top               += 1;
        $this->array[$this->top] = $x;
    }

    /**
     * 出栈
     * @return bool|mixed
     */
    public function pop()
    {
        if ($this->empty()) {
            return false;
        } else {
            $this->top -= 1;
            return $this->array[$this->top + 1];
        }
    }
}

class queue
{
    public $array;
    public $tail;
    public $head;
    public $length;

    /**
     * 队列
     * queue constructor.
     * @param array $array
     */
    public function __construct(array $array)
    {
        $this->length = count($array);
        //key从1开始
        $array = array_merge([null], $array);
        unset($array[0]);
        $this->array = $array;
        $this->tail  = $this->head = 1;
    }

    /**
     * 进队
     * @param $x
     */
    public function enQueue($x)
    {
        $this->array[$this->tail] = $x;
        if ($this->tail == $this->length) {
            $this->tail = 1;
        } else {
            $this->tail += 1;
        }
    }

    /**
     *出队
     */
    public function deQueue()
    {
        $x = $this->array[$this->head];
        if ($this->head == $this->length) {
            $this->head = 1;
        } else {
            $this->head = $this->head + 1;
        }
        return $x;
    }
}

class linkList
{
    public $nil = null;

    /**
     * 链表
     * linkList constructor.
     * @param $array
     */
    public function __construct($array)
    {
        $this->nil = new elementList(null);

        if (empty($array)) {
            $this->nil->next = $this->nil;
            $this->nil->prep = $this->nil;
        }
        $b = new elementList(null);
        foreach ($array as $key => $value) {
            if ($key == 0) {
                //头元素的特殊处理
                $a               = new elementList($value);
                $this->nil->next = $a;
                $a->prep         = $this->nil;
            } else {
                $a       = new elementList($value);
                $b->next = $a;
                $a->prep = $b;
            }
            $b = $a;

            //尾元素的特殊处理
            if ($key == count($array) - 1) {
                $b->next         = $this->nil;
                $this->nil->prep = $b;
            }
        }
    }

    /**
     * 按值寻找
     * @param $k
     * @return mixed
     */
    public function search($k)
    {
        $x = $this->nil->next;
        while ($x->key != null && $x->key != $k) {
            $x = $x->next;
        }
        return $x;
    }

    /**
     * 插入
     * @param elementList $x
     */
    public function insert(elementList $x)
    {
        $x->next               = $this->nil->next;
        $this->nil->next->prev = $x;
        $this->nil->next       = $x;
        $x->prep               = $this->nil;
    }

    /**
     * 删除
     * @param elementList $x
     */
    public function delete(elementList $x)
    {
        $x->prep->next = $x->next;
        $x->next->prep = $x->prep;
    }
}

class elementList
{
    public $next;
    public $prep;
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