<?php

class A {
    public static function get_self(){
        // 获取实例化的对象
        return new self();
    }
    public static function get_static(){
        return new static();
    }
}
class B extends A {}

var_dump(B::get_static());
var_dump(B::get_self());
var_dump(A::get_static());
var_dump(A::get_self());

// object(B)#1 (0) {
// }
// object(A)#1 (0) {
// }
// object(A)#1 (0) {
// }
// object(A)#1 (0) {
// }

// 以上输出结果为上述；但是具体这个demo在说什么我也不明白，不过我只是为了验证一下 get static 是干嘛的东西

// 下面我们一起探讨一下 - php中对象，类与内存的关系

// https://blog.csdn.net/q423498555/article/details/50835364/

// 这篇文章看完 - 我们仅仅知道，php变量-分为 引用 和实际内容存储；而引用是存储在 栈之中（先进先出），内容存储在堆之中（具体啥忘了）
// 一个类对象实例化了以后，会执行魔术方法2个；__construct \ __destruct 就是这两个；而第一个开始执行的时候（new class()就是开辟了一片内存
// 存储了这个变量 - destruct的时候就是将引用清零（回收内存） -- 当然这个执行方法在什么时候执行的呢？？？暂时我也不知道

// 目前关于静态变量仅仅知道他是全局唯一的，不属于任何一个类的实例化对象 - 这个static的变量 什么时候回收的呢？


// 说到swoole的连接池 - pdo对象；类静态变量的回收，你就要知道：
// swoole的内存管理机制是什么样子的？？？
// https://blog.csdn.net/guiyecheng/article/details/60866463

