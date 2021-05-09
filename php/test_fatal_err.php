<?php

try {
    //error 这是错误  不是异常  7.0之前 捕获不到
    test();
    //code...
} catch (\Throwable $th) {
    //throw $th;
    var_dump($th);
}



