<?php

require_once 'ErrorCatch.php';
// require_once 'bb.php';
// E_ERROR
// throw new Exception("Value must be 1 or below");
// $a = array();

// echo 123; 
// throw new Exception("Value must be 1 or below");
// var_dump(999);
// test();
echo $b;
// var_dump(123);

// throw new Exception("Value must be 1 or below");

// try {
//     echo b();
// } catch (\Throwable $th) {
//     var_dump('try-catch',$th->getMessage());
// }


// $userInfo = array('name' => 'yaxin');
// echo $userInfo['age'];

// var_dump($a["tet"]);
// var_dump(explode(',',$a)); // Fatal error: Uncaught TypeError: explode(
// echo file_get_contents('./a.t'); -- Warning: file_get_contents(./a.t): Failed to open stream:
// echo $a['test'];//PHP Warning:  Undefined variable $ -- 会继续往下走的哦

// var_dump(123);

// try {
    // echo $b;
// } catch (\Throwable $th) {
    // var_dump(666,$th->getMessage());//无法捕获 PHP Warning:  Undefined variable $
// }

// $b = array();

// echo $b['test'];// Warning: Undefined array key "test"

// echo tt();//Fatal error: Uncaught Error:  不会继续往下走的
// var_dump(999);
// $a = new test();

// try {
//     throw new Exception("Value must be 1 or below");
// } catch (\Throwable $th) {
//     var_dump(666,$th->getMessage());//无法捕获 PHP Warning:  Undefined variable $
// }



// https://www.php.net/manual/zh/book.errorfunc.php
// 错误和异常不是同一种东西，异常是php5之后才有的。

// php的错误可以归为三类：notice(提示)，wanring(警告)，error(错误)；

// notice和warning都是会继续往下执行的，而error是会直接中断程序执行


// error级别错误有： 

// 1 函数传递参数错误 ：fatal error: Uncaught TypeError: explode(): Argument #2 ($string) must be of type string
// 2 调用未定义的函数：Fatal error: Uncaught Error: Call to undefined function test()
// 3 调用未定义的类：PHP Fatal error:  Uncaught Error: Class "test" not found 

// warning级别错误有：

// 1 调用未定义的变量：Warning: Undefined variable $a
// 2 数组键值不存在：PHP Warning:  Undefined array key "test" 

// 关于异常：

// 异常是需要开发者手动使用throw抛出来的。就是说，如果不使用throw，php不会发生异常

// 如何处理：

// set_error_handler 
// set_exception_handler
// register_shutdown_function

// 其中，对与try-catch，php waring级别的错误是不会被捕获到的

// 对于fatal-error比如调用未定义的函数test()，在php7.0之前，只能通过register_shutdown_function + error_get_last捕获 
// 在PHP7.0之后，会被 set_exception_handler 捕获；同理，也会被try-catch捕获
// (改动声明：https://www.php.net/manual/zh/language.errors.php7.php)
// notice和warning（错误）是不会被set_exception_handler捕获的也不会被try-catch捕获

// 所以，大可不必像PHP5.6那样通过register_shutdown_function + error_get_last捕获E_ERROR级别的错误



// The following error types cannot be handled with a user defined function: 
// 以下错误类型不能用用户定义的函数处理
// E_ERROR(致命错误-fatal error), 
// E_PARSE(错误是编译时发生的,比如语句结束没加上分号),
// E_CORE_ERROR（PHP引擎核心产生的，发生在PHP初始化启动过程中）, 
// E_CORE_WARNING（PHP引擎核心产生的，发生在PHP初始化启动过程中）, 
// E_COMPILE_ERROR（编译时产生的，发生在脚本编译过程中）, 
// E_COMPILE_WARNING（编译时产生的，发生在脚本编译过程中）
// 这几个函数都是直接导致脚本终止运行的



