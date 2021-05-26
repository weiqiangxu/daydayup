<?php

// 如何用2个栈实现一个队列

class my_quene {
    public $stack_one = array();
    public $stack_two = array();

    function get(&$stack){
        // 返回最后一个元素
        return array_pop($stack);
    }
    function set(&$stack,$item){
        // 插入一个元素
        $stack[] = $item;
        return true;
    }

    function pop(){
        if(!empty($this->stack_two)){
            return $this->get($this->stack_two);
        }else{
            // 从栈1逐个元素往栈2压入
            $c = count($this->stack_one);
            for ($i=0; $i < $c; $i++) { 
                $this->set($this->stack_two,$this->get($this->stack_one));
            }
            var_dump($this->stack_one);
            return $this->get($this->stack_two);
        }
    }

    function push($item){
        $this->set($this->stack_one,$item);
        return true;
    }
}

$q = new my_quene();
$q->push("a");
$q->push("b");
$q->push("c");
var_dump($q->pop());
var_dump($q->pop());