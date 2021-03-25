<?php

// 关于mixphp的model的db连接池

// 执行顺序是什么样子的，请问他的db模型体现了什么设计模式

// db连接池的话什么时候连接什么时候释放？长链接还是短连接？现在的事务（basemode类静态变量实现）在a连接没有rollback也没有commit的情况下
// 进来b有没有问题？没问题的因为不管update还是其他的都是直接 getConnection 获取的对象

// basemodel 的静态变量有没有释放

// Mix\Database\Pool ConnectionPool
// AbstractConnectionPool
// creare 
//  pdo func
// [sql] 2021-03-20 14:31:05 <2772> [message] 0.22ms SELECT * FROM zp_job_info LIMIT 0, 1
// discatd

// Mix\Database\Coroutine   -- use Mix\Pool\ConnectionTrait  （这里的ondestruct会进行 discard）

// Mix\Pool\AbstractConnectionPool->onInitialize
// Mix\Database\Pool\ConnectionPool->getConnection
// Mix\Pool\AbstractConnectionPool->getConnection
// Mix\Pool\AbstractConnectionPool->getIdleNumber
// Mix\Pool\AbstractConnectionPool->getTotalNumber
// Mix\Pool\AbstractConnectionPool->getIdleNumber
// Mix\Pool\AbstractConnectionPool->getActiveNumber
// Mix\Pool\AbstractConnectionPool->createConnection
// 怎么过来的
// Mix\Database\Coroutine\PDOConnection->onInitialize
// Mix\Database\Base\AbstractPDOConnection->onInitialize
// Mix\Database\Base\AbstractPDOConnection->prepare
// Mix\Database\Base\AbstractPDOConnection->buildQueryFragment
// Mix\Database\Base\AbstractPDOConnection->buildQueryFragment
// Mix\Database\Base\AbstractPDOConnection->buildQueryFragment
// Mix\Database\Base\AbstractPDOConnection->bindParams
// Mix\Database\Base\AbstractPDOConnection->build
// Mix\Database\Base\AbstractPDOConnection->autoConnect
// Mix\Database\Base\AbstractPDOConnection->connect
// Mix\Database\Base\AbstractPDOConnection->createConnection
// Mix\Database\Base\AbstractPDOConnection->microtime
// Mix\Database\Base\AbstractPDOConnection->microtime
// Mix\Database\Base\AbstractPDOConnection->clearBuild
// Mix\Database\Base\AbstractPDOConnection->getLastLog
// [sql] 2021-03-20 15:28:52 <11826> [message] 0.22ms SELECT * FROM zp_job_info LIMIT 0, 1
// Mix\Database\Coroutine\PDOConnection->onDestruct
// Mix\Database\Persistent\PDOConnection->onDestruct
// Mix\Database\Base\AbstractPDOConnection->disconnect
// Mix\Pool\AbstractConnectionPool->discard


// ---

// "autoload": {
//     "psr-4": {
//       "Common\\": "applications/common/src/",
//       "Console\\": "applications/console/src/",
//       "Http\\": "applications/http/src/",
//       "Core\\": "applications/core/src/",
//       "WebSocket\\": "applications/websocket/src/",
//       "Tcp\\": "applications/tcp/src/",
//       "Udp\\": "applications/udp/src/"
//     },
//     "classmap": [
//       "applications/common/src/ApplicationInterface.php"
//     ],
//     "files": [
//       "applications/common/src/functions.php"
//     ]
//   },

//   composer的这几个参数是什么意思需要解释一下